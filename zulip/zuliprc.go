package zulip

import (
	"errors"
	"fmt"
	"os"
	"strings"

	"gopkg.in/ini.v1"
)

// contextKeys are used to identify the type of value in the context.
// Since these are string, it is possible to get a short description of the
// context key for logging and debugging using key.String().

type contextKey string

// ContextBasicAuth takes BasicAuth as authentication for the request.
const ContextBasicAuth = contextKey("basic")

func (c contextKey) String() string {
	return "auth " + string(c)
}

// ErrMissingURLError is returned when no Zulip server URL is provided.
var ErrMissingURLError = errors.New("missing Zulip server URL; specify via --site or ~/.zuliprc")

type ConfigMissingError struct {
	message string
}

func (e ConfigMissingError) Error() string { return e.message }

// BasicAuth provides basic http authentication to a request passed via context using ContextBasicAuth.
type BasicAuth struct {
	UserName string `json:"userName,omitempty"`
	Password string `json:"password,omitempty"`
}

// APIKey provides API key based authentication to a request passed via context using ContextAPIKey.
type APIKey struct {
	Key    string
	Prefix string
}

// Configuration stores the configuration of the API client.
type RC struct {
	Email         string `json:"email,omitempty"`
	APIKey        string `json:"apiKey,omitempty"`
	Site          string `json:"site,omitempty"`
	ClientCert    string `json:"clientCert,omitempty"`
	ClientCertKey string `json:"clientCertKey,omitempty"`
	CertBundle    string `json:"certBundle,omitempty"`
	Insecure      *bool  `json:"insecure,omitempty"`
	ConfigFile    string `json:"configFile,omitempty"`
}

func NewZulipRCFromFile(zuliprc string) (*RC, error) {
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

	rc := &RC{
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
			return nil, fmt.Errorf(
				"insecure is set to '%s', it must be 'true' or 'false' if it is used in %s",
				insecureSetting,
				zuliprc,
			)
		}
		rc.Insecure = &parsed
	}

	populateFromEnv(rc)

	err = rc.Validate()
	if err != nil {
		return nil, err
	}

	return rc, nil
}

func NewZulipRC() (*RC, error) {
	rc := &RC{}

	err := populateFromEnv(rc)
	if err != nil {
		return nil, fmt.Errorf("unable to populate config from environment: %w", err)
	}
	err = rc.Validate()
	if err != nil {
		return nil, err
	}
	return rc, nil
}

func populateFromEnv(rc *RC) error {
	if rc.Insecure == nil {
		if insecureSetting := os.Getenv("ZULIP_ALLOW_INSECURE"); insecureSetting != "" {
			parsed, err := parseStrictBool(insecureSetting)
			if err != nil {
				return fmt.Errorf(
					"the ZULIP_ALLOW_INSECURE environment variable is set to %q, it must be 'true' or 'false'",
					insecureSetting,
				)
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

func (rc *RC) Validate() error {
	if rc == nil {
		return errors.New("invalid configuration: nil")
	}

	if rc.APIKey == "" {
		return ConfigMissingError{message: "api_key not specified"}
	}

	if rc.Email == "" {
		return ConfigMissingError{message: "email not specified"}
	}

	if rc.Site == "" {
		return ErrMissingURLError
	}

	if rc.ClientCert == "" && rc.ClientCertKey != "" {
		return ConfigMissingError{
			message: fmt.Sprintf(
				"client cert key '%s' specified, but no client cert public part provided",
				rc.ClientCertKey,
			),
		}
	}

	if rc.ClientCert != "" {
		if _, err := os.Stat(rc.ClientCert); err != nil {
			return ConfigMissingError{message: fmt.Sprintf("client cert '%s' does not exist", rc.ClientCert)}
		}
		if rc.ClientCertKey != "" {
			if _, err := os.Stat(rc.ClientCertKey); err != nil {
				return ConfigMissingError{message: fmt.Sprintf("client cert key '%s' does not exist", rc.ClientCertKey)}
			}
		}
	}

	if rc.CertBundle != "" {
		if _, err := os.Stat(rc.CertBundle); err != nil {
			return ConfigMissingError{message: fmt.Sprintf("tls bundle '%s' does not exist", rc.CertBundle)}
		}
	}
	return nil
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
