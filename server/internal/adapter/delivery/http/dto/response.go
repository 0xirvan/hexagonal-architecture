package dto

type SingleResponse[T any] struct {
	Data T `json:"data"`
}

func ToSingleResponse[T any](item T) SingleResponse[T] {
	return SingleResponse[T]{Data: item}
}

type PaginationMeta struct {
	TotalItems  int  `json:"total_items"`
	TotalPages  int  `json:"total_pages"`
	CurrentPage int  `json:"current_page"`
	PageSize    int  `json:"page_size"`
	HasNext     bool `json:"has_next"`
	HasPrev     bool `json:"has_prev"`
}

type PaginatedResponse[T any] struct {
	Data       []T            `json:"data"`
	Pagination PaginationMeta `json:"pagination"`
}

func ToPaginatedResponse[T any](items []T, total, page, size int) PaginatedResponse[T] {
	totalPages := (total + size - 1) / size
	return PaginatedResponse[T]{
		Data: items,
		Pagination: PaginationMeta{
			TotalItems:  total,
			TotalPages:  totalPages,
			CurrentPage: page,
			PageSize:    size,
			HasNext:     page < totalPages,
			HasPrev:     page > 1,
		},
	}
}
