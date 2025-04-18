package service

import (
	"boiler-plate-clean/internal/model"
	"boiler-plate-clean/internal/repository"
	"boiler-plate-clean/pkg/exception"
	"boiler-plate-clean/pkg/xvalidator"
	"context"
	"gorm.io/gorm"
)

type StockServiceImpl struct {
	db          *gorm.DB
	repo        repository.StockRepository
	productRepo repository.ProductRepository
	validate    *xvalidator.Validator
}

func NewStockService(
	db *gorm.DB, repo repository.StockRepository,
	productRepo repository.ProductRepository,
	validate *xvalidator.Validator,
) StockService {
	return &StockServiceImpl{
		db:          db,
		repo:        repo,
		productRepo: productRepo,
		validate:    validate,
	}
}

// CreateStock creates a new campaign
func (s *StockServiceImpl) Create(
	ctx context.Context, req *model.CreateStockReq,
) (*model.CreateStockRes, *exception.Exception) {
	tx := s.db.Begin()
	defer tx.Rollback()
	if errs := s.validate.Struct(req); errs != nil {
		return nil, exception.InvalidArgument(errs)
	}
	body := req.ToEntity()
	product, err := s.productRepo.FindByID(ctx, s.db, req.ProductId)
	if err != nil {
		return nil, exception.Internal("failed to get product", err)
	}
	if product == nil {
		return nil, exception.NotFound("product not found")
	}
	if err := s.repo.CreateTx(ctx, tx, body); err != nil {
		return nil, exception.Internal("failed to create stock", err)
	}

	if err := tx.Commit().Error; err != nil {
		return nil, exception.Internal("failed to commit db", err)
	}
	return &model.CreateStockRes{Stock: body}, nil
}

func (s *StockServiceImpl) Update(
	ctx context.Context, req *model.UpdateStockReq,
) (*model.UpdateStockRes, *exception.Exception) {
	tx := s.db.Begin()
	defer tx.Rollback()
	if errs := s.validate.Struct(req); errs != nil {
		return nil, exception.InvalidArgument(errs)
	}
	body := req.ToEntity()
	result, err := s.repo.FindByID(ctx, s.db, req.ID)
	if err != nil {
		return nil, exception.Internal("failed to find stock", err)
	}

	if result == nil {
		return nil, exception.NotFound("stock not found")
	}
	body.ID = result.ID
	if err := s.repo.UpdateTx(ctx, tx, body); err != nil {
		return nil, exception.Internal("failed to update stock", err)
	}

	if err := tx.Commit().Error; err != nil {
		return nil, exception.Internal("failed to commit stock", err)
	}
	return &model.UpdateStockRes{Stock: body}, nil
}

func (s *StockServiceImpl) Delete(
	ctx context.Context, req *model.DeleteStockReq,
) (*model.DeleteStockRes, *exception.Exception) {
	tx := s.db.Begin()
	defer tx.Rollback()

	if err := s.repo.DeleteByIDTx(ctx, tx, req.ID); err != nil {
		return nil, exception.Internal("failed to delete stock", err)
	}

	if err := tx.Commit().Error; err != nil {
		return nil, exception.Internal("failed to commit db", err)
	}
	return &model.DeleteStockRes{ID: req.ID}, nil
}

func (s *StockServiceImpl) GetById(
	ctx context.Context, req *model.GetStockByIdReq,
) (*model.GetStockByIdRes, *exception.Exception) {
	result, err := s.repo.FindByID(ctx, s.db, req.ID)
	if err != nil {
		return nil, exception.Internal("failed to get stock error", err)
	}

	if result == nil {
		return nil, exception.NotFound("stock not found")
	}

	return &model.GetStockByIdRes{Stock: result}, nil
}

func (s *StockServiceImpl) GetList(
	ctx context.Context, req *model.GetListStockReq,
) (*model.GetListStockRes, *exception.Exception) {
	result, err := s.repo.FindByPagination(ctx, s.db, req.Page, req.Order, req.Filter)
	if err != nil {
		return nil, exception.Internal("failed to get stock error", err)
	}
	return &model.GetListStockRes{
		Data:       result.Data,
		Pagination: result.ToPagination(),
	}, nil
}
