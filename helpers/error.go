package helpers

type CustomError struct {
	Message       string `json:"message"`
	OriginalError error  `json:"original_error,omitempty"`
}

func (e *CustomError) Error() string {
	return e.Message
}

func NewCustomError(message string, err error) *CustomError {
	return &CustomError{Message: message, OriginalError: err}
}
