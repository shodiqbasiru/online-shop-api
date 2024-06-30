package exception

type BadRequestError struct {
	Error string
}

func NewBadRequestError(error string) *BadRequestError {
	return &BadRequestError{Error: error}
}
