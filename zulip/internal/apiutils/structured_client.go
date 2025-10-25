package apiutils

import (
	"context"
	"net/http"
)

type ResponseModel interface {
	GetIgnoredParametersUnsupported() []string
}

type StructuredClient interface {
	CallAPI(ctx context.Context, endpoint string, req *http.Request, responseModel ResponseModel) (httpResp *http.Response, err error)
	ServerURL() (string, error)
	GetUserAgent() string
}
