package exceptions

type BadRequestError struct {
	ErrorMessage string
}

func NewBadRequestError(error string) BadRequestError {
	return BadRequestError{ErrorMessage: error}
}

func (e BadRequestError) Error() string {
	return e.ErrorMessage
}