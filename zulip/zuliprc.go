package zulip

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"gopkg.in/ini.v1"
)

// contextKeys are used to identify the type of value in the context.
// Since these are string, it is possible to get a short description of the
// context key for logging and debugging using key.String().

type contextKey string

func (c contextKey) String() string {
	return "auth " + string(c)
}

var (
	// ContextBasicAuth takes BasicAuth as authentication for the request.
	ContextBasicAuth = contextKey("basic")

	ErrMissingURLError = errors.New("missing Zulip server URL; specify via --site or ~/.zuliprc")
)

type ConfigMissingError struct {
	message string
}

func (e ConfigMissingError) Error() string { return e.message }

// BasicAuth provides basic http authentication to a request passed via context using ContextBasicAuth
type BasicAuth struct {
	UserName string `json:"userName,omitempty"`
	Password string `json:"password,omitempty"`
}

// APIKey provides API key based authentication to a request passed via context using ContextAPIKey
type APIKey struct {
	Key    string
	Prefix string
}

// Configuration stores the configuration of the API client
type ZulipRC struct {
	Email         string `json:"email,omitempty"`
	APIKey        string `json:"apiKey,omitempty"`
	Site          string `json:"site,omitempty"`
	ClientCert    string `json:"clientCert,omitempty"`
	ClientCertKey string `json:"clientCertKey,omitempty"`
	CertBundle    string `json:"certBundle,omitempty"`
	Insecure      *bool  `json:"insecure,omitempty"`
	ConfigFile    string `json:"configFile,omitempty"`
}

func NewZulipRCFromFile(zuliprc string) (*ZulipRC, error) {
	info, err := os.Stat(zuliprc)
	if err != nil {
		if os.IsNotExist(err) {
			return nil, ConfigMissingError{message: fmt.Sprintf("config file %s not found", zuliprc)}
		}
		return nil, err
	}
	if info.IsDir() {
		return nil, fmt.Errorf("config path %s is a directory", zuliprc)
	}

	cfg, err := ini.LoadSources(ini.LoadOptions{IgnoreInlineComment: true}, zuliprc)
	if err != nil {
		return nil, err
	}

	section, err := cfg.GetSection("api")
	if err != nil {
		return nil, fmt.Errorf("missing [api] section in config file %s", zuliprc)
	}

	rc := &ZulipRC{
		APIKey:        section.Key("key").String(),
		Email:         section.Key("email").String(),
		Site:          section.Key("site").String(),
		ClientCert:    section.Key("client_cert").String(),
		ClientCertKey: section.Key("client_cert_key").String(),
		CertBundle:    section.Key("cert_bundle").String(),
		ConfigFile:    zuliprc,
	}

	if insecureSetting := section.Key("insecure").String(); insecureSetting != "" {
		parsed, err := parseStrictBool(insecureSetting)
		if err != nil {
			return nil, fmt.Errorf("insecure is set to '%s', it must be 'true' or 'false' if it is used in %s", insecureSetting, zuliprc)
		}
		rc.Insecure = &parsed
	}

	populateFromEnv(rc)

	return rc, nil
}

func NewZulipRC() (*ZulipRC, error) {
	rc := &ZulipRC{}

	err := populateFromEnv(rc)
	if err != nil {
		return nil, fmt.Errorf("unable to populate config from environment: %w", err)
	}
	return rc, nil
}

func populateFromEnv(rc *ZulipRC) error {
	if rc.Insecure == nil {
		if insecureSetting := os.Getenv("ZULIP_ALLOW_INSECURE"); insecureSetting != "" {
			parsed, err := parseStrictBool(insecureSetting)
			if err != nil {
				return fmt.Errorf("the ZULIP_ALLOW_INSECURE environment variable is set to %q, it must be 'true' or 'false'", insecureSetting)
			}
			rc.Insecure = &parsed
		}
	}

	if rc.APIKey == "" {
		rc.APIKey = os.Getenv("ZULIP_API_KEY")
	}
	if rc.Email == "" {
		rc.Email = os.Getenv("ZULIP_EMAIL")
	}
	if rc.Site == "" {
		rc.Site = os.Getenv("ZULIP_SITE")
	}
	if rc.ClientCert == "" {
		rc.ClientCert = os.Getenv("ZULIP_CERT")
	}
	if rc.ClientCertKey == "" {
		rc.ClientCertKey = os.Getenv("ZULIP_CERT_KEY")
	}
	if rc.CertBundle == "" {
		rc.CertBundle = os.Getenv("ZULIP_CERT_BUNDLE")
	}
	return nil
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

func parseStrictBool(value string) (bool, error) {
	normalized := strings.ToLower(strings.TrimSpace(value))
	switch normalized {
	case "true":
		return true, nil
	case "false":
		return false, nil
	default:
		return false, fmt.Errorf("invalid boolean value: %s", value)
	}
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

func getDefaultConfigFilename() string {
	home, err := os.UserHomeDir()
	if err != nil {
		return ""
	}
	return filepath.Join(home, ".zuliprc")
}
