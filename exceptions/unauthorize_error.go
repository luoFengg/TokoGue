package exceptions

type UnauthorizedError struct {
	ErrorMessage string
}

func NewUnauthorizedError(error string) UnauthorizedError {
	return UnauthorizedError{ErrorMessage: error}
}

func (e UnauthorizedError) Error() string {
	return e.ErrorMessage
}