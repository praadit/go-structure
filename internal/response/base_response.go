package response

type BaseResponse struct {
	ErrorCode       string    `json:"error_code,omitempty"`
	Message         any       `json:"message,omitempty"`
	InternalMessage any       `json:"internal_message,omitempty"`
	Payload         *BaseData `json:"payload,omitempty"`
}

type BaseData struct {
	Data       any `json:"data,omitempty"`
	Pagination any `json:"pagination,omitempty"`
}
