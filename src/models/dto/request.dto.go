package dto

type PaginationRequest struct {
	Page    int    `json:"page"`
	Perpage int    `json:"per_page"`
	Order   string `json:"order"`
	OrderBy string `json:"order_by"`
	Search  string `json:"search"`
}
