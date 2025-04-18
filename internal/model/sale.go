package model

import "boiler-plate-clean/internal/entity"

type BaseSaleReq struct {
	ProductId int     `json:"product_id" validate:"required"`
	Quantity  int     `json:"quantity" validate:"required,gt=0"`
	Price     float64 `json:"price" validate:"required,gt=0"`
}
type CreateSaleReq struct {
	BaseSaleReq
}

func (r BaseSaleReq) ToEntity() *entity.Sale {
	return &entity.Sale{
		ProductId: r.ProductId,
		Quantity:  r.Quantity,
		Price:     r.Price,
	}
}

type UpdateSaleReq struct {
	ID int `json:"id"`
	CreateSaleReq
}

type GetSaleByIdReq struct {
	ID int `json:"id"`
}

type DeleteSaleReq struct {
	ID int `json:"id"`
}

type GetListSaleReq struct {
	Page   PaginationParam
	Filter FilterParams
	Order  OrderParam
}

type CreateSaleRes struct {
	Sale *entity.Sale
}

type UpdateSaleRes struct {
	Sale *entity.Sale
}

type GetSaleByIdRes struct {
	Sale *entity.Sale
}

type DeleteSaleRes struct {
	ID int `json:"id"`
}

type GetListSaleRes struct {
	Data       []*entity.Sale
	Pagination *Pagination
}

type ProfitReportReq struct {
	Month int `form:"month" validate:"required,min=1,max=12"`
	Year  int `form:"year" validate:"required"`
}

type ProfitReportRes struct {
	Month          int     `json:"month"`
	Year           int     `json:"year"`
	TotalPenjualan float64 `json:"total_penjualan"`
	TotalHPP       float64 `json:"total_hpp"`
	Laba           float64 `json:"laba"`
}
