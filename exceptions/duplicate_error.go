package exceptions

type DuplicateError struct {
	ErrorMessage string
}

func NewDuplicateError(error string) DuplicateError {
	return DuplicateError{ErrorMessage: error}
}

func (e DuplicateError) Error() string {
	return e.ErrorMessage
}
