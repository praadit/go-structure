package exception

type BadRequestError struct {
	Code    string
	Message string
}

func (e *BadRequestError) Error() string {
	return e.Message
}
