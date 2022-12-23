package request

import (
	"encoding/json"
	"errors"
)

type PaginationRequest struct {
	Page    int    `json:"page"`
	Perpage int    `json:"per_page"`
	Order   string `json:"order"`
	OrderBy string `json:"order_by"`
	Search  string `json:"search"`
}

func (pagReq *PaginationRequest) ValidateOrder() error {
	message := []string{}
	if pagReq.Order == "" {
		message = append(message, "Order wajib diisi")
	}
	if pagReq.OrderBy == "" {
		message = append(message, "OrderBy wajib diisi")
	}

	if len(message) > 0 {
		messageString, err := json.Marshal(message)
		if err != nil {
			panic(err)
		}
		return errors.New(string(messageString))
	}

	return nil
}

type Pagination struct {
	TotalItem int64 `json:"total_item"`
	TotalPage int64 `json:"total_page"`
	Page      int   `json:"page"`
}
