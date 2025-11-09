package dto

// PaginatedResponse represents a paginated response structure
// using generics to accommodate any data type
type PaginatedResponse[T any] struct {
	Data        []T  `json:"data"`
	TotalItems  int  `json:"total_items"`
	TotalPages  int  `json:"total_pages"`
	CurrentPage int  `json:"current_page"`
	PageSize    int  `json:"page_size"`
	HasNext     bool `json:"has_next"`
	HasPrev     bool `json:"has_prev"`
}
