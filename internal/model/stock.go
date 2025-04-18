package model

import (
	"boiler-plate-clean/internal/entity"
)

type BaseStockReq struct {
	Price     float64 `json:"price" validate:"required,gt=0"`
	Quantity  int     `json:"quantity" validate:"required,gt=0"`
	ProductId int     `json:"product_id" validate:"required,gt=0"`
}

func (req BaseStockReq) ToEntity() *entity.Stock {
	return &entity.Stock{
		Price:     req.Price,
		Quantity:  req.Quantity,
		ProductId: req.ProductId,
	}
}

// Requests
type CreateStockReq struct {
	BaseStockReq
}

type UpdateStockReq struct {
	ID int `json:"id"`
	BaseStockReq
}

type GetListStockReq struct {
	Page   PaginationParam
	Filter FilterParams
	Order  OrderParam
}

type GetStockByIdReq struct {
	ID int `json:"id"`
}

type DeleteStockReq struct {
	ID int `json:"id"`
}

// Responses
type CreateStockRes struct {
	*entity.Stock
}

type UpdateStockRes struct {
	*entity.Stock
}

type GetListStockRes struct {
	Data       []*entity.Stock
	Pagination *Pagination
}

type GetStockByIdRes struct {
	*entity.Stock
}

type DeleteStockRes struct {
	ID int `json:"id"`
}
