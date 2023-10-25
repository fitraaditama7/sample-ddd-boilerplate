package custom_error

type HTTPError struct {
	StatusCode   int
	ResponseCode string
	Message      string
}

func (e *HTTPError) Error() string {
	return e.Message
}
