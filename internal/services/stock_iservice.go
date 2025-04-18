package service

import (
	"boiler-plate-clean/internal/model"
	"boiler-plate-clean/pkg/exception"
	"context"
)

type StockService interface {
	// CRUD operations for Stock
	Create(ctx context.Context, req *model.CreateStockReq) (*model.CreateStockRes, *exception.Exception)
	GetById(ctx context.Context, req *model.GetStockByIdReq) (*model.GetStockByIdRes, *exception.Exception)
	GetList(ctx context.Context, req *model.GetListStockReq) (*model.GetListStockRes, *exception.Exception)
	Update(ctx context.Context, req *model.UpdateStockReq) (*model.UpdateStockRes, *exception.Exception)
	Delete(ctx context.Context, req *model.DeleteStockReq) (*model.DeleteStockRes, *exception.Exception)
}
