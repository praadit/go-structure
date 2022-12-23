package exception

type MaintenanceError struct {
	Code    string
	Message string
}

func (e *MaintenanceError) Error() string {
	return e.Message
}
