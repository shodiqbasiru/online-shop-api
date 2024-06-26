package exception

type InternalServerError struct {
	Error string
}

func NewInternalServerError(error string) *InternalServerError {
	return &InternalServerError{Error: error}
}
