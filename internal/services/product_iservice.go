package service

import (
	"boiler-plate-clean/internal/model"
	"boiler-plate-clean/pkg/exception"
	"context"
)

type ProductService interface {
	// CRUD operations for Product
	Create(ctx context.Context, req *model.CreateProductReq) (*model.CreateProductRes, *exception.Exception)
	GetById(ctx context.Context, req *model.GetProductByIdReq) (*model.GetProductByIdRes, *exception.Exception)
	GetList(ctx context.Context, req *model.GetListProductReq) (*model.GetListProductRes, *exception.Exception)
	Update(ctx context.Context, req *model.UpdateProductReq) (*model.UpdateProductRes, *exception.Exception)
	Delete(ctx context.Context, req *model.DeleteProductReq) (*model.DeleteProductRes, *exception.Exception)
}
