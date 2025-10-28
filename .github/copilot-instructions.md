# AI Coding Agent Instructions for go-zulip


## Project Overview

**go-zulip** is a Go API client for the Zulip group chat platform. The codebase consists of 100+ Go files organized into a modular service-based architecture. The project structure and tooling have been updated for improved maintainability and code quality.


### Key Facts
- **Language**: Go 1.25.2
- **Scope**: Zulip REST API v1.0.0 client
- **Testing**: Integration tests with live Zulip server, automated via GitHub Actions
- **Linters/Formatters**: `golangci-lint`, `golangci-lint fmt`, enforced via pre-commit hooks

---


## Directory Structure

```
go.mod, go.sum, LICENSE, README.md
internal/
  apiutils/         # API utility helpers
  clients/          # HTTP client implementations and helpers
  enum/             # Enum types
  strict_decoder/   # Strict JSON decoding utilities
  test_utils/       # Test helpers
  union/            # Union type marshaling logic
scripts/            # Dev scripts
zulip/
  *.go              # Data models, API logic, and tests
  api/              # API domain service implementations
  client/           # Main client and statistics
  events/           # Event types and event handling
.github/
  workflows/        # CI configuration (GitHub Actions)
.pre-commit-config.yaml # Pre-commit hooks config
```

---

## Architecture Patterns

### 1. **Service-Based Client Architecture**
The codebase uses a **service-oriented architecture** where the main client (`client.Client`) aggregates multiple API domain services:
- **Client**: `client.Client` interface in `zulip/client/client.go` aggregates multiple API domain interfaces
- **Services**: Each API domain (authentication, channels, messages, etc.) has its own service implementation
- **API Client**: Core HTTP client in `zulip/internal/api_client/APIClient` handles HTTP communication
- **Why**: Clean separation of concerns, easy testing, and modular design

### 2. **Builder Pattern for Requests**
Every API endpoint uses the **Builder Pattern** for fluent configuration:
```go
// Example from api/channels/api_channels.go
resp, httpRes, err := client.APIChannels.AddDefaultChannel(ctx).
    ChannelID(123).
    Execute()
```
Structure:
- Service method returns a `Request` struct (e.g., `AddDefaultChannelRequest`)
- Setter methods (e.g., `.ChannelID()`) return the same request type for chaining
- `.Execute()` finalizes and sends the request

### 3. **API Domain Organization**
API functionality is organized by domain with clear separation:
- `zulip/api/*/api_*.go` - API service implementations with request builders
- `zulip/api/*/api_*_responses.go` - Response types for each domain
- `zulip/*.go` - Data models (150+ enum and struct types)

---


## CI, Testing, and Linters

### Continuous Integration (CI)
- **GitHub Actions**: All pushes and PRs trigger workflows in `.github/workflows/`.
  - Linting: Runs `golangci-lint run` and `golangci-lint fmt` checks.
  - Testing: Runs `go test ./zulip/...` (requires Zulip server credentials for integration tests).
  - Pre-commit: Validates all pre-commit hooks pass.

### Pre-commit Hooks
- Configured in `.pre-commit-config.yaml` (run `pre-commit install` to enable).
- Enforces:
  - `golangci-lint` (static analysis, style, and bug checks)
  - `golangci-lint fmt` (formatting)
  - Trailing whitespace, end-of-file, and YAML checks

### Running Locally
- **Lint:** `golangci-lint run`
- **Format:** `golangci-lint fmt ./...`
- **Test:** `go test ./zulip/...`
- **Pre-commit:** `pre-commit run --all-files`

---

## Configuration & Authentication

### Client Configuration
Create a client using `zulip.RC` config with email, API key, and server URL:
```go
config := &z.RC{
  Email:  "user@example.com",
  APIKey: "your-api-key",
  Site:   "https://zulip.example.com",
}
client, err := client.NewClient(config)
```

Optional configuration:
- Client certificates: `ClientCert`, `ClientCertKey`, `CertBundle`
- Skip TLS verification: `Insecure: true`
- Custom HTTP client: `client.WithHTTPClient(httpClient)`
- Custom logger: Configure via `client.WithLogger()` option (uses `slog` by default)

For testing, there are helper functions that automatically fetch test credentials from the development server.


---

## Critical Patterns & Conventions

### 1. **Error Handling**
Three error types to know:
- `APIError` - General API failures (wraps body and unmarshaled model)
- `CodedError` - Server error with `code` field (e.g., "BAD_REQUEST", "UNAUTHORIZED")
- Special errors: `RateLimitedError`, `BadEventQueueIdError`, `NonExistingChannelIdError`

Example:
```go
resp, _, err := client.GetUser(ctx, 123).Execute()
if err != nil {
  var codedErr z.CodedError
  if errors.As(err, &codedErr) {
    log.Printf("API error %s: %s", codedErr.Code, codedErr.Msg)
  }
}
```

### 2. **Union Types** (in `zulip/internal/utils/union.go`)
Handles OpenAPI discriminated unions using reflection-based marshaling/unmarshaling:
- Struct with pointer fields for each union variant
- `MarshalUnionType()`/`UnmarshalUnionType()` ensures exactly ONE field is set
- Used for ambiguous response types (e.g., narrow types, event types)

**Note:** Event types live in the `events` package (e.g., `events/message_event.go`) instead of the `zulip` package, because they blow up the package docks significantly and the real-time events API behaves very differently.

### 3. **Narrow Types** (in `narrow.go`)
Represents query filters for messages - complex union type with 15+ variants:
```go
client.GetMessages(ctx).
  Narrow(z.Where(z.ChannelNameIs("announcements")).
    And(z.TopicIs("releases")))
```

### 4. **Response Objects**
All API responses inherit from `Response`:
```go
type Response struct {
  Result  string `json:"result"`           // "success" or "error"
  Msg     string `json:"msg,omitempty"`   // error message if result="error"
}
```
Specific responses add domain fields (e.g., `GetUsersResponse.Members []User`).

**Event types are now located in the `events` package.** For example, see `events/message_event.go` for message event types.

### 5. **Timestamp Handling**
Timestamps from Zulip API are returned as integers or floats, but the client converts them to from and to `time.Time` or `time.Duration` transparently.


---

## Testing Patterns

### Test File Structure
Tests follow standard Go patterns with integration focus:
- Tests in `*_test.go` files in separate `zulip_test` package
- Requires valid Zulip server running in development mode
- Tests use helper functions defined within test files

### Common Test Patterns
```go
// 1. Setup - create resources
channelID := createChannelWithAllClients(t)

// 2. Execute
resp, _, err := client.AddReaction(ctx, msgID).
  EmojiName("smile").
  Execute()

// 3. Verify
require.NoError(t, err)
require.NotNil(t, resp)
```

### Multi-Client Testing
Tests often run against multiple clients (normal, admin, bot) to verify permission handling:
```go
t.Run("TestName", RunForAllClients(t, func(t *testing.T, apiClient z.Client) {
  // test runs once per client
}))
```

---


### Build, Lint, and Test
```bash
# Build
go build ./...

# Lint (static analysis)
golangci-lint run

# Format
golangci-lint fmt ./...

# Run all tests (requires live Zulip server)
go test ./... -v

# Run specific test
go test ./... -run TestName

# Run all pre-commit hooks
pre-commit run -a
```

### Manual Code Organization
This is a hand-written codebase with clear separation of concerns:
- `zulip/api/*/api_*.go` - API domain service implementations with request builders
- `zulip/api/*/api_*_responses.go` - Response types organized by domain
- `zulip/*.go` - Data models and enums
- `zulip/client/client.go` - Main client aggregating all API services
- `zulip/internal/api_client/api_client.go` - Core HTTP client implementation
- `zulip/errors.go` - Error type definitions
- `zulip/internal/utils/union.go` - Union type marshaling logic
- `zulip/zuliprc/zuliprc.go` - Configuration loading from INI files
- `*_test.go` - Integration tests

---


## Key Files & Directories

| File/Dir | Purpose |
|----------|---------|
| `zulip/client/client.go` | Client interface definition and service aggregation |
| `internal/clients/` | HTTP client implementations and helpers |
| `internal/union/union.go` | Union type marshaling logic |
| `zulip/rc.go` | Configuration loading from INI files |
| `zulip/errors.go` | Error type definitions |
| `zulip/api/*/api_*.go` | Domain-specific API service implementations |
| `zulip/events/` | Event types and event handling |
| `zulip/*.go` | Data model structs and enums |
| `*_test.go` | Integration tests |


---

## When Adding New Features
1. **Add to appropriate API interface** - Domain files already have interface definitions
2. **Update request type** - Add request struct and setter methods (follow existing patterns in `api_*.go`)
3. **Write integration tests** - Use helpers and `RunForAllClients()` for multi-client testing
4. **Handle edge cases** - Account for permission-based API differences (admin vs user vs bot)
5. **Update documentation** - Ensure README and code comments reflect new functionality
6. **Lint and test** - Run `golangci-lint` and tests to ensure code quality
