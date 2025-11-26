package exceptions

type NotFoundError struct {
	ErrorMessage string
}

func NewNotFoundError(error string) NotFoundError {
	return NotFoundError{ErrorMessage: error}
}

func (e NotFoundError) Error() string {
	return e.ErrorMessage
}
