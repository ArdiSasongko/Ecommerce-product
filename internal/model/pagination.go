package model

type (
	PaginatinParams struct {
		Offset int
		Limit  int
	}

	CategoryWithPaginationResponse struct {
		Data       []CategoryResponse `json:"data"`
		TotalCount int                `json:"total_count"`
		Offset     int                `json:"offset"`
		Limit      int                `json:"limit"`
	}
)
