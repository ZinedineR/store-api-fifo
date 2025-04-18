package service

import (
	"boiler-plate-clean/internal/model"
	"boiler-plate-clean/internal/repository"
	"boiler-plate-clean/pkg/exception"
	"boiler-plate-clean/pkg/xvalidator"
	"context"
	"gorm.io/gorm"
)

type SaleServiceImpl struct {
	db          *gorm.DB
	repo        repository.SaleRepository
	stockRepo   repository.StockRepository
	productRepo repository.ProductRepository
	validate    *xvalidator.Validator
}

func NewSaleService(
	db *gorm.DB, repo repository.SaleRepository,
	stockRepo repository.StockRepository,
	productRepo repository.ProductRepository,
	validate *xvalidator.Validator,
) SaleService {
	return &SaleServiceImpl{
		db:          db,
		repo:        repo,
		stockRepo:   stockRepo,
		productRepo: productRepo,
		validate:    validate,
	}
}

func (s *SaleServiceImpl) Create(
	ctx context.Context, req *model.CreateSaleReq,
) (*model.CreateSaleRes, *exception.Exception) {
	tx := s.db.Begin()
	defer tx.Rollback()

	if errs := s.validate.Struct(req); errs != nil {
		return nil, exception.InvalidArgument(errs)
	}

	product, err := s.productRepo.FindByID(ctx, tx, req.ProductId)
	if err != nil {
		return nil, exception.Internal("failed to get product", err)
	}
	if product == nil {
		return nil, exception.NotFound("product not found")
	}

	// Get FIFO stocks
	stocks, err := s.stockRepo.FindAvailableStockFIFO(ctx, tx, req.ProductId)
	if err != nil {
		return nil, exception.Internal("failed to get stock FIFO", err)
	}

	qtyLeft := req.Quantity
	var totalHPP float64
	for _, stock := range stocks {
		if qtyLeft <= 0 {
			break
		}
		takeQty := stock.Quantity
		if takeQty > qtyLeft {
			takeQty = qtyLeft
		}
		totalHPP += float64(takeQty) * stock.Price
		qtyLeft -= takeQty

		if err := s.stockRepo.DecreaseStockQtyTx(ctx, tx, stock.ID, takeQty); err != nil {
			return nil, exception.Internal("failed to update stock", err)
		}
	}

	if qtyLeft > 0 {
		return nil, exception.InvalidArgument("not enough stock")
	}

	sale := req.ToEntity()
	sale.TotalHPP = totalHPP

	if err := s.repo.CreateTx(ctx, tx, sale); err != nil {
		return nil, exception.Internal("failed to create sale", err)
	}

	if err := tx.Commit().Error; err != nil {
		return nil, exception.Internal("failed to commit sale", err)
	}

	return &model.CreateSaleRes{Sale: sale}, nil
}

func (s *SaleServiceImpl) GetProfitReport(
	ctx context.Context, req *model.ProfitReportReq,
) (*model.ProfitReportRes, *exception.Exception) {
	if errs := s.validate.Struct(req); errs != nil {
		return nil, exception.InvalidArgument(errs)
	}

	res, err := s.repo.GetProfitReport(ctx, s.db, req.Month, req.Year)
	if err != nil {
		return nil, exception.Internal("failed to get profit report", err)
	}
	return res, nil
}
