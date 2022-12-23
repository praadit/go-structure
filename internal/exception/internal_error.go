package exception

type InternalError struct {
	Code            string
	InternalMessage string
}

func (e *InternalError) Error() string {
	return e.InternalMessage
}
