package exception

type ForbiddendError struct {
	Code    string
	Message string
}

func (e *ForbiddendError) Error() string {
	return e.Message
}
