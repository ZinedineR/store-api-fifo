package model

import (
	"boiler-plate-clean/internal/entity"
)

type BaseProductReq struct {
	Name string `json:"name" validate:"required"`
}

func (req BaseProductReq) ToEntity() *entity.Product {
	return &entity.Product{
		Name: req.Name,
	}
}

// Requests
type CreateProductReq struct {
	BaseProductReq
}

type UpdateProductReq struct {
	ID int `json:"id"`
	BaseProductReq
}

type GetListProductReq struct {
	Page   PaginationParam
	Filter FilterParams
	Order  OrderParam
}

type GetProductByIdReq struct {
	ID int `json:"id"`
}

type DeleteProductReq struct {
	ID int `json:"id"`
}

// Responses
type CreateProductRes struct {
	*entity.Product
}

type UpdateProductRes struct {
	*entity.Product
}

type GetListProductRes struct {
	Data       []*entity.Product
	Pagination *Pagination
}

type GetProductByIdRes struct {
	*entity.Product
}

type DeleteProductRes struct {
	ID int `json:"id"`
}
