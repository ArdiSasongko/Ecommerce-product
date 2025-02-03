package service

import (
	"github.com/ArdiSasongko/Ecommerce-product/internal/model"
)

func ApplyPaginationCategoris(data []model.CategoryResponse, params model.PaginatinParams) []model.CategoryResponse {
	if params.Offset >= len(data) {
		return []model.CategoryResponse{}
	}

	end := params.Offset + params.Limit
	if end > len(data) {
		end = len(data)
	}

	return data[params.Offset:end]
}

func ApplyPaginationProducts(data []model.ProductsResponse, params model.PaginatinParams) []model.ProductsResponse {
	if params.Offset >= len(data) {
		return []model.ProductsResponse{}
	}

	end := params.Offset + params.Limit
	if end > len(data) {
		end = len(data)
	}

	return data[params.Offset:end]
}
