package exceptions

type ValidationError struct {
	ErrorMessage string
}

func NewValidationError(error string) *ValidationError {
	return &ValidationError{ErrorMessage: error}
}

func (e *ValidationError) Error() string {
	return e.ErrorMessage
}