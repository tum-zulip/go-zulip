# ZulipRC configuration

The Go client reads connection details from the same `zuliprc` format as the
Python tooling. Each entry lives under the `[api]` section. You can either load
an existing file with `api.NewZulipRCFromFile` or construct the struct manually
before passing it to `api.NewZulipClient`.

## Supported keys

| Key | `ZulipRC` field | Description |
|-----|-----------------|-------------|
| `email` | `Email` | Account email address to authenticate with. |
| `key` | `APIKey` | API key for the email address. |
| `site` | `Site` | Base Zulip server URL, including scheme (e.g. `https://chat.zulip.org`). |
| `client_cert` | `ClientCert` | Optional path to a client TLS certificate. |
| `client_cert_key` | `ClientCertKey` | Optional path to the private key matching `client_cert`. |
| `cert_bundle` | `CertBundle` | Optional PEM bundle of custom certificate authorities. |
| `insecure` | `Insecure` | Set to `true` to skip certificate validation (testing only). |

Empty fields fall back to environment variables with matching semantics:

| Environment variable | Notes |
|----------------------|-------|
| `ZULIP_EMAIL` | Overrides `email`. |
| `ZULIP_API_KEY` | Overrides `key`. |
| `ZULIP_SITE` | Overrides `site`. |
| `ZULIP_CERT` | Overrides `client_cert`. |
| `ZULIP_CERT_KEY` | Overrides `client_cert_key`. |
| `ZULIP_CERT_BUNDLE` | Overrides `cert_bundle`. |
| `ZULIP_ALLOW_INSECURE` | Overrides `insecure`. Must be `true` or `false`. |

Paths within the file or environment variables may contain `~` or other
environment variables; both are expanded automatically before being used.

## TLS behaviour

`buildHTTPClient` configures TLS according to the resolved values:

- When `cert_bundle` is provided, it is appended to the root certificate pool.
- When `client_cert` (and optionally `client_cert_key`) are provided, the client
  presents that certificate during the TLS handshake.
- When `insecure` is `true`, certificate verification is disabled. This should
  only be used for local testing.

If both `client_cert` and `client_cert_key` are set, the paths must point to
existing files. Any failures result in a `ConfigNotFoundError` during client
construction.
