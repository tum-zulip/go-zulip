package clients

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSanitizeHTTPDump_AuthorizationBasic(t *testing.T) {
	input := `GET /api/v1/users/me HTTP/1.1
Host: zulip.example.com
Authorization: Basic dXNlckBleGFtcGxlLmNvbTpzZWNyZXRrZXkxMjM=
User-Agent: go-zulip/1.0.0

`

	result := sanitizeHTTPDump(input)

	assert.Contains(t, result, "Authorization: Basic [REDACTED]")
	assert.NotContains(t, result, "dXNlckBleGFtcGxlLmNvbTpzZWNyZXRrZXkxMjM=")
	assert.Contains(t, result, "Host: zulip.example.com")
	assert.Contains(t, result, "User-Agent: go-zulip/1.0.0")
}

func TestSanitizeHTTPDump_AuthorizationBearer(t *testing.T) {
	input := `GET /api/v1/users/me HTTP/1.1
Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9

`

	result := sanitizeHTTPDump(input)

	assert.Contains(t, result, "Authorization: Bearer [REDACTED]")
	assert.NotContains(t, result, "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9")
}

func TestSanitizeHTTPDump_Cookie(t *testing.T) {
	input := `GET /api/v1/users/me HTTP/1.1
Cookie: session=abc123; token=xyz789

`

	result := sanitizeHTTPDump(input)

	assert.Contains(t, result, "Cookie: [REDACTED-")
	assert.NotContains(t, result, "session=abc123")
}

func TestSanitizeHTTPDump_SetCookie(t *testing.T) {
	input := `HTTP/1.1 200 OK
Set-Cookie: session=abc123; Path=/; HttpOnly

`

	result := sanitizeHTTPDump(input)

	assert.Contains(t, result, "Set-Cookie:")
	assert.NotContains(t, result, "session=abc123")
}

func TestSanitizeHTTPDump_MultipleHeaders(t *testing.T) {
	input := `GET /api/v1/users/me HTTP/1.1
Authorization: Basic dXNlckBleGFtcGxlLmNvbTpzZWNyZXRrZXkxMjM=
Cookie: session=abc123
X-API-Key: secret-api-key-456

`

	result := sanitizeHTTPDump(input)

	assert.Contains(t, result, "Authorization: Basic [REDACTED]")
	assert.Contains(t, result, "Cookie: [REDACTED-")
	assert.Contains(t, result, "X-API-Key: [REDACTED-")
	assert.NotContains(t, result, "dXNlckBleGFtcGxlLmNvbTpzZWNyZXRrZXkxMjM=")
	assert.NotContains(t, result, "session=abc123")
	assert.NotContains(t, result, "secret-api-key-456")
}

func TestSanitizeHTTPDump_CaseInsensitive(t *testing.T) {
	input := `GET /api/v1/users/me HTTP/1.1
authorization: Basic dXNlckBleGFtcGxlLmNvbTpzZWNyZXRrZXkxMjM=
COOKIE: session=abc123

`

	result := sanitizeHTTPDump(input)

	assert.Contains(t, result, "authorization: Basic [REDACTED]")
	assert.Contains(t, result, "COOKIE: [REDACTED-")
}

func TestSanitizeHTTPDump_NoSensitiveHeaders(t *testing.T) {
	input := `GET /api/v1/users/me HTTP/1.1
Host: zulip.example.com
Content-Type: application/json

`

	result := sanitizeHTTPDump(input)

	assert.Equal(t, input, result)
}

func TestSanitizeHTTPDump_WithBody(t *testing.T) {
	input := `POST /api/v1/messages HTTP/1.1
Authorization: Basic dXNlckBleGFtcGxlLmNvbTpzZWNyZXRrZXkxMjM=
Content-Type: application/json

{"type":"stream","to":"general","content":"Hello"}
`

	result := sanitizeHTTPDump(input)

	assert.Contains(t, result, "Authorization: Basic [REDACTED]")
	assert.NotContains(t, result, "dXNlckBleGFtcGxlLmNvbTpzZWNyZXRrZXkxMjM=")
	assert.Contains(t, result, `{"type":"stream","to":"general","content":"Hello"}`)
}

func TestSanitizeHTTPDump_EmptyString(t *testing.T) {
	result := sanitizeHTTPDump("")
	assert.Equal(t, "", result)
}

func TestCreateRedactedValue(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "Basic auth",
			input:    "Basic dXNlckBleGFtcGxlLmNvbTpzZWNyZXRrZXkxMjM=",
			expected: "Basic [REDACTED]",
		},
		{
			name:     "Bearer token",
			input:    "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9",
			expected: "Bearer [REDACTED]",
		},
		{
			name:     "Plain value",
			input:    "secret-value-123",
			expected: "[REDACTED-16-chars]",
		},
		{
			name:     "Empty value",
			input:    "",
			expected: "[REDACTED]",
		},
		{
			name:     "Whitespace only",
			input:    "   ",
			expected: "[REDACTED]",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := createRedactedValue(tt.input)
			assert.Equal(t, tt.expected, result)
		})
	}
}
