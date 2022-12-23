package utilities

import "go-best-practice/internal/response"

func ResponseBuilder(message string, data interface{}, pagination interface{}) response.BaseResponse {
	if data == nil {
		return response.BaseResponse{
			Message: message,
		}
	}
	return response.BaseResponse{
		Message: message,
		Payload: &response.BaseData{
			Data:       data,
			Pagination: pagination,
		},
	}
}
