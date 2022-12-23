package exception

type NotFoundError struct {
	Code    string
	Message string
}

func (e *NotFoundError) Error() string {
	return e.Message
}
