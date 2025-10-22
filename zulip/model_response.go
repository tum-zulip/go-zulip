package zulip

const (
	ResponseSuccess = "success"
	ResponseError   = "error"
)

type Response struct {
	Result string `json:"result,omitempty"`
	Msg    string `json:"msg,omitempty"`
	// An array of any parameters sent in the request that are not supported by the endpoint.  See [error handling] documentation for details on this and its change history.
	//
	// [error handling]: https://zulip.com/api/rest-error-handling#ignored-parameters
	IgnoredParametersUnsupported []string `json:"ignored_parameters_unsupported,omitempty"`
}

type responseModel interface {
	getIgnoredParametersUnsupported() []string
}

func (o *Response) getIgnoredParametersUnsupported() []string {
	if o == nil || o.IgnoredParametersUnsupported == nil {
		var empty []string
		return empty
	}
	return o.IgnoredParametersUnsupported
}
