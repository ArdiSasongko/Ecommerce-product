package model

type (
	PaginatinParams struct {
		Offset int
		Limit  int
	}

	CategoryWithPaginationResponse struct {
		Categories []CategoryResponse `json:"categories"`
		TotalCount int                `json:"total_count"`
		Offset     int                `json:"offset"`
		Limit      int                `json:"limit"`
	}

	ProductsWithPaginationResponse struct {
		Products   []ProductsResponse `json:"products"`
		TotalCount int                `json:"total_count"`
		Offset     int                `json:"offset"`
		Limit      int                `json:"limit"`
	}
)
