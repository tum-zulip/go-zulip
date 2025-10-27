package clients

import (
	"context"
	"errors"
	"log/slog"
	"net/http"
	"reflect"
	"strings"

	"github.com/tum-zulip/go-zulip/zulip"
	strictdecoder "github.com/tum-zulip/go-zulip/zulip/internal/strict_decoder"
)

func tryUnmarshalErrorModel[T any](data []byte) (T, error) {
	var model T
	dec := strictdecoder.New(data)
	err := dec.Decode(&model)
	if err != nil {
		return model, err
	}
	if reflect.ValueOf(model).IsZero() {
		return model, errors.New("no data")
	}
	return model, nil
}

//nolint:funlen,nolintlint
func unmarshallAPIError(ctx context.Context, logger *slog.Logger, status int, body []byte) error {
	var model error
	var err error

	contentType := http.DetectContentType(body)
	if strings.HasPrefix(contentType, "text/html") {
		return zulip.NewAPIError(body, errors.New(http.StatusText(status)))
	}

	model, err = tryUnmarshalErrorModel[zulip.CodedError](body)
	if err == nil {
		// clean match, return early
		return zulip.NewAPIError(body, model)
	}

	// try more specific error models
	switch status {
	case http.StatusTooManyRequests:
		model, err = tryUnmarshalErrorModel[zulip.RateLimitedError](body)
	case http.StatusBadRequest:
		fallthrough
	case http.StatusNotFound:
		model, err = tryUnmarshalErrorModel[zulip.BadEventQueueIDError](body)
		if err == nil {
			break
		}
		model, err = tryUnmarshalErrorModel[zulip.BadNarrowError](body)
		if err == nil {
			break
		}
		model, err = tryUnmarshalErrorModel[zulip.IncompatibleParametersError](body)
		if err == nil {
			break
		}
		model, err = tryUnmarshalErrorModel[zulip.MissingArgumentError](body)
		if err == nil {
			break
		}
		model, err = tryUnmarshalErrorModel[zulip.NonExistingChannelIDError](body)
		if err == nil {
			break
		}
		model, err = tryUnmarshalErrorModel[zulip.InvitationFailedError](body)
		if err == nil {
			break
		}
		model, err = tryUnmarshalErrorModel[zulip.NonExistingChannelNameError](body)
		if err == nil {
			break
		}
	}

	if err == nil {
		// found a matching error model
		return zulip.NewAPIError(body, model)
	}

	logger.WarnContext(
		ctx,
		"API returned an unknown error Response",
		"status",
		status,
		"body",
		string(body),
		"model",
		model,
	)
	return zulip.NewAPIError(body, errors.New(http.StatusText(status)))
}
