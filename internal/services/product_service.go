package service

import (
	"boiler-plate-clean/internal/model"
	"boiler-plate-clean/internal/repository"
	"boiler-plate-clean/pkg/exception"
	"boiler-plate-clean/pkg/xvalidator"
	"context"
	"gorm.io/gorm"
)

type ProductServiceImpl struct {
	db       *gorm.DB
	repo     repository.ProductRepository
	validate *xvalidator.Validator
}

func NewProductService(
	db *gorm.DB, repo repository.ProductRepository,
	validate *xvalidator.Validator,
) ProductService {
	return &ProductServiceImpl{
		db:       db,
		repo:     repo,
		validate: validate,
	}
}

// CreateProduct creates a new campaign
func (s *ProductServiceImpl) Create(
	ctx context.Context, req *model.CreateProductReq,
) (*model.CreateProductRes, *exception.Exception) {
	tx := s.db.Begin()
	defer tx.Rollback()
	if errs := s.validate.Struct(req); errs != nil {
		return nil, exception.InvalidArgument(errs)
	}
	body := req.ToEntity()
	if err := s.repo.CreateTx(ctx, tx, body); err != nil {
		return nil, exception.Internal("failed to create product", err)
	}

	if err := tx.Commit().Error; err != nil {
		return nil, exception.Internal("failed to commit db", err)
	}
	return &model.CreateProductRes{Product: body}, nil
}

func (s *ProductServiceImpl) Update(
	ctx context.Context, req *model.UpdateProductReq,
) (*model.UpdateProductRes, *exception.Exception) {
	tx := s.db.Begin()
	defer tx.Rollback()
	if errs := s.validate.Struct(req); errs != nil {
		return nil, exception.InvalidArgument(errs)
	}
	body := req.ToEntity()
	result, err := s.repo.FindByID(ctx, s.db, req.ID)
	if err != nil {
		return nil, exception.Internal("failed to find product", err)
	}

	if result == nil {
		return nil, exception.NotFound("product not found")
	}
	body.ID = result.ID
	if err := s.repo.UpdateTx(ctx, tx, body); err != nil {
		return nil, exception.Internal("failed to update product", err)
	}

	if err := tx.Commit().Error; err != nil {
		return nil, exception.Internal("failed to commit product", err)
	}
	return &model.UpdateProductRes{Product: body}, nil
}

func (s *ProductServiceImpl) Delete(
	ctx context.Context, req *model.DeleteProductReq,
) (*model.DeleteProductRes, *exception.Exception) {
	tx := s.db.Begin()
	defer tx.Rollback()

	if err := s.repo.DeleteByIDTx(ctx, tx, req.ID); err != nil {
		return nil, exception.Internal("failed to delete product", err)
	}

	if err := tx.Commit().Error; err != nil {
		return nil, exception.Internal("failed to commit db", err)
	}
	return &model.DeleteProductRes{ID: req.ID}, nil
}

func (s *ProductServiceImpl) GetById(
	ctx context.Context, req *model.GetProductByIdReq,
) (*model.GetProductByIdRes, *exception.Exception) {
	result, err := s.repo.FindByID(ctx, s.db, req.ID)
	if err != nil {
		return nil, exception.Internal("failed to get product error", err)
	}

	if result == nil {
		return nil, exception.NotFound("product not found")
	}

	return &model.GetProductByIdRes{Product: result}, nil
}

func (s *ProductServiceImpl) GetList(
	ctx context.Context, req *model.GetListProductReq,
) (*model.GetListProductRes, *exception.Exception) {
	result, err := s.repo.FindByPagination(ctx, s.db, req.Page, req.Order, req.Filter)
	if err != nil {
		return nil, exception.Internal("failed to get product error", err)
	}
	return &model.GetListProductRes{
		Data:       result.Data,
		Pagination: result.ToPagination(),
	}, nil
}
