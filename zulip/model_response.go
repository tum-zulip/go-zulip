package zulip

const (
	ResponseSuccess = "success"
	ResponseError   = "error"
)

type Response struct {
	Result string `json:"result,omitempty"`
	Msg    string `json:"msg,omitempty"`
	// An array of any parameters sent in the request that are not supported by the endpoint.  See [error handling](zulip.com/api/rest-error-handling#ignored-parameters documentation for details on this and its change history.
	IgnoredParametersUnsupported []string `json:"ignored_parameters_unsupported,omitempty"`
}

func (o *Response) GetResult() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Result
}

func (o *Response) GetMsg() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Msg
}
