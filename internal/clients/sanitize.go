package clients

import (
	"fmt"
	"regexp"
	"strings"
)

// sensitiveHeaders lists HTTP headers that should be sanitized in debug logs
// to prevent credential leakage.
var sensitiveHeaders = []string{
	"Authorization",
	"Cookie",
	"Set-Cookie",
	"X-API-Key",
	"Proxy-Authorization",
	"WWW-Authenticate",
	"Proxy-Authenticate",
}

// sanitizeHTTPDump removes sensitive header values from HTTP request/response dumps
// while preserving the structure for debugging purposes.
//
// This function is critical for security: it prevents API credentials, session tokens,
// and other sensitive data from being exposed in debug logs. The Authorization header,
// which contains base64-encoded email:apiKey credentials, is the primary concern.
//
// The function preserves non-sensitive headers and HTTP bodies, maintaining the utility
// of debug logs while protecting credentials.
func sanitizeHTTPDump(dump string) string {
	result := dump

	for _, header := range sensitiveHeaders {
		// Match header line (case-insensitive): "Header-Name: value" or "Header-Name: value\r"
		// The pattern uses:
		// - (?mi): multiline + case-insensitive flags
		// - ^: start of line
		// - (\r?): optional carriage return at end
		pattern := regexp.MustCompile(`(?mi)^(` + regexp.QuoteMeta(header) + `):\s*(.+?)(\r?)$`)

		result = pattern.ReplaceAllStringFunc(result, func(match string) string {
			submatches := pattern.FindStringSubmatch(match)
			if len(submatches) < 4 {
				return match
			}

			headerName := submatches[1]
			originalValue := submatches[2]
			lineEnding := submatches[3]

			redactedValue := createRedactedValue(originalValue)
			return headerName + ": " + redactedValue + lineEnding
		})
	}

	return result
}

// createRedactedValue creates a redacted representation of a sensitive value
// showing a hint of its length for debugging purposes.
//
// For Authorization headers, it preserves the auth type (Basic/Bearer) but redacts
// the actual credentials. For other values, it shows the length to help with debugging
// truncation or encoding issues.
func createRedactedValue(value string) string {
	valueLen := len(strings.TrimSpace(value))

	if valueLen == 0 {
		return "[REDACTED]"
	}

	// For Authorization: Basic, show type but redact the base64 credentials
	if strings.HasPrefix(value, "Basic ") {
		return "Basic [REDACTED]"
	}

	// For Authorization: Bearer, show type but redact the token
	if strings.HasPrefix(value, "Bearer ") {
		return "Bearer [REDACTED]"
	}

	// For other values, show length hint to help debug truncation/encoding issues
	return fmt.Sprintf("[REDACTED-%d-chars]", valueLen)
}
