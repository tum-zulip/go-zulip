package clients

import (
	"crypto/tls"
	"crypto/x509"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"os"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/tum-zulip/go-zulip/zulip"
)

//nolint:funlen
func buildHTTPClient(
	baseClient *http.Client,
	params *zulip.RC,
	clientName string,
	logger *slog.Logger,
	warnOnInsecureTLS bool,
) (*http.Client, string, error) {
	if params == nil {
		return nil, "", errors.New("invalid configuration: nil")
	}

	if logger == nil {
		logger = slog.Default()
	}

	if baseClient == nil {
		baseClient = http.DefaultClient
	}

	userAgent := buildUserAgent(clientName)

	transport := http.DefaultTransport
	if baseClient.Transport != nil {
		transport = http.DefaultTransport
	}

	if t, ok := transport.(*http.Transport); ok {
		cp := t.Clone()
		configureTransportTLS(cp, params, logger, warnOnInsecureTLS)
		transport = cp
	} else {
		if params.CertBundle != "" || params.ClientCert != "" {
			logger.Warn(
				"custom HTTP transport does not support TLS customization; certificate options are ignored",
			)
		}
		if params.Insecure != nil && *params.Insecure && warnOnInsecureTLS {
			logger.Warn(
				"custom HTTP transport does not support TLS customization; insecure option may be ignored",
			)
		}
	}

	clientCopy := *baseClient
	clientCopy.Transport = &basicAuthRoundTripper{
		email:     params.Email,
		apiKey:    params.APIKey,
		userAgent: userAgent,
		base:      transport,
	}

	return &clientCopy, userAgent, nil
}

func configureTransportTLS(transport *http.Transport, params *zulip.RC, logger *slog.Logger, warnOnInsecureTLS bool) error {
	if transport.TLSClientConfig == nil {
		transport.TLSClientConfig = &tls.Config{
			MinVersion: tls.VersionTLS12,
		}
	} else {
		cfgCopy := transport.TLSClientConfig.Clone()
		if cfgCopy.MinVersion == 0 {
			cfgCopy.MinVersion = tls.VersionTLS12
		}
		transport.TLSClientConfig = cfgCopy
	}

	if params.Insecure != nil && *params.Insecure {
		if warnOnInsecureTLS {
			logger.Warn(
				"Insecure mode enabled. The server's SSL/TLS certificate will not be validated, making the HTTPS connection potentially insecure",
			)
		}
		transport.TLSClientConfig.InsecureSkipVerify = true
	}

	if params.CertBundle != "" {
		if err := addCertBundleToTransport(transport, params.CertBundle); err != nil {
			return fmt.Errorf("failed to add certificate bundle: %w", err)
		}
	}

	if params.ClientCert != "" {
		if err := addClientCertToTransport(transport, params.ClientCert, params.ClientCertKey); err != nil {
			return fmt.Errorf("failed to add client certificate: %w", err)
		}
	}
	return nil
}

func addCertBundleToTransport(transport *http.Transport, bundlePath string) error {
	bundlePath, err := normalizePath(bundlePath)
	if err != nil {
		return err
	}

	data, err := os.ReadFile(bundlePath)
	if err != nil {
		return err
	}

	var pool *x509.CertPool
	if transport.TLSClientConfig.RootCAs != nil {
		pool = transport.TLSClientConfig.RootCAs
	} else {
		pool, err = x509.SystemCertPool()
		if err != nil {
			pool = x509.NewCertPool()
		}
	}

	if ok := pool.AppendCertsFromPEM(data); !ok {
		return fmt.Errorf("failed to parse certificate bundle %s", bundlePath)
	}
	transport.TLSClientConfig.RootCAs = pool
	return nil
}

func addClientCertToTransport(transport *http.Transport, clientCert, keyPath string) error {
	certPath, err := normalizePath(clientCert)
	if err != nil {
		return err
	}

	if keyPath == "" {
		keyPath = clientCert
	}

	keyPath, err = normalizePath(keyPath)
	if err != nil {
		return err
	}

	cert, err := tls.LoadX509KeyPair(certPath, keyPath)
	if err != nil {
		return fmt.Errorf("failed to load client certificate: %w", err)
	}
	transport.TLSClientConfig.Certificates = []tls.Certificate{cert}
	return nil
}

func buildUserAgent(clientName string) string {
	vendor, version := detectPlatform()
	return fmt.Sprintf("%s (%s; %s)", clientName, vendor, version)
}
func normalizePath(pathStr string) (string, error) {
	expanded := os.ExpandEnv(pathStr)
	if strings.HasPrefix(expanded, "~") {
		home, err := os.UserHomeDir()
		if err != nil {
			return "", err
		}
		expanded = filepath.Join(home, strings.TrimPrefix(expanded, "~"))
	}
	return filepath.Abs(expanded)
}

type basicAuthRoundTripper struct {
	email     string
	apiKey    string
	userAgent string
	base      http.RoundTripper
}

func (t *basicAuthRoundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	clone := req.Clone(req.Context())
	clone.SetBasicAuth(t.email, t.apiKey)
	if t.userAgent != "" {
		clone.Header.Set("User-Agent", t.userAgent)
	}
	return t.base.RoundTrip(clone)
}

func detectPlatform() (string, string) {
	switch runtime.GOOS {
	case "linux":
		if name, version, err := readOSRelease(); err == nil {
			return name, version
		}
		return "Linux", runtime.GOARCH
	case "windows":
		return "Windows", runtime.GOARCH
	case "darwin":
		return "Darwin", runtime.GOARCH
	default:
		return runtime.GOOS, runtime.GOARCH
	}
}

func readOSRelease() (string, string, error) {
	file, err := os.Open("/etc/os-release")
	if err != nil {
		return "", "", err
	}
	defer file.Close()

	content, err := io.ReadAll(file)
	if err != nil {
		return "", "", err
	}

	var name, version string
	lines := strings.Split(string(content), "\n")
	for _, line := range lines {
		if strings.HasPrefix(line, "NAME=") && name == "" {
			name = trimOSReleaseValue(line[5:])
		}
		if strings.HasPrefix(line, "VERSION_ID=") && version == "" {
			version = trimOSReleaseValue(line[11:])
		}
	}

	if name == "" {
		name = "Linux"
	}
	if version == "" {
		version = runtime.GOARCH
	}

	return name, version, nil
}

func trimOSReleaseValue(value string) string {
	value = strings.TrimSpace(value)
	value = strings.TrimPrefix(value, "\"")
	value = strings.TrimSuffix(value, "\"")
	return value
}
