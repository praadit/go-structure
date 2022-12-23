package exception

type UnauthorizedError struct {
	Code    string
	Message string
}

func (e *UnauthorizedError) Error() string {
	return e.Message
}
