package presenter

type ErrorResponse struct {
	Error   *string `json:"error,omitempty"`
	Message *string `json:"message,omitempty"`
}

func NewErrorResponse(error string, message string) ErrorResponse {
	e := ErrorResponse{}
	if len(error) > 0 {
		e.Error = &error
	}

	if len(message) > 0 {
		e.Message = &message
	}

	return e
}
