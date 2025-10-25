package zulip

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestRateLimitedErrorMarshal(t *testing.T) {
	err := RateLimitedError{
		CodedError: CodedError{
			Response: Response{
				Result: "error",
				Msg:    "test",
			},
			Code: "TEST",
		},
		RetryAfter: 100 * time.Millisecond,
	}

	// Test marshaling by value
	data, marshalErr := json.Marshal(err)
	assert.NoError(t, marshalErr)
	t.Logf("Marshaled JSON (by value): %s", string(data))

	// Test marshaling by pointer
	dataPtr, marshalErrPtr := json.Marshal(&err)
	assert.NoError(t, marshalErrPtr)
	t.Logf("Marshaled JSON (by pointer): %s", string(dataPtr))

	// Unmarshal it back
	var decoded RateLimitedError
	unmarshalErr := json.Unmarshal(dataPtr, &decoded)
	assert.NoError(t, unmarshalErr)
	t.Logf("Decoded RetryAfter: %v", decoded.RetryAfter)

	assert.Equal(t, 100*time.Millisecond, decoded.RetryAfter)
}
