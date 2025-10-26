package clients

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"reflect"
	"strings"

	"github.com/tum-zulip/go-zulip/zulip"
	"github.com/tum-zulip/go-zulip/zulip/internal/utils"
)

func tryUnmarshalErrorModel[T any](data []byte) (T, error) {
	var model T
	dec := utils.NewStrictDecoder(data)
	err := dec.Decode(&model)
	if err != nil {
		return model, err
	}
	if reflect.ValueOf(model).IsZero() {
		return model, errors.New("no data")
	}
	return model, nil
}

func unmarshallAPIError(ctx context.Context, logger *slog.Logger, status int, body []byte) error {
	var model interface{}
	var err error

	contentType := http.DetectContentType(body)
	if strings.HasPrefix(contentType, "text/html") {
		return zulip.NewAPIError(body, fmt.Sprintf("status %d: %s", status, string(body)), nil)
	}

	model, err = tryUnmarshalErrorModel[zulip.CodedError](body)
	if err != nil {
		switch status {
		case http.StatusTooManyRequests:
			model, err = tryUnmarshalErrorModel[zulip.RateLimitedError](body)
		case http.StatusBadRequest:
			fallthrough
		case http.StatusNotFound:
			model, err = tryUnmarshalErrorModel[zulip.BadEventQueueIdError](body)
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
			model, err = tryUnmarshalErrorModel[zulip.NonExistingChannelIdError](body)
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
	}

	if err != nil {
		logger.WarnContext(ctx, "API returned an unknown error response", "status", status, "body", string(body), "model", model)
		return zulip.NewAPIError(body, fmt.Sprintf("status %d: %s", status, string(body)), nil)
	}
	return zulip.NewAPIError(body, fmt.Sprintf("status %d: %s", status, string(body)), model)

}
