package repository

import (
	"boiler-plate-clean/internal/entity"
	"context"
	"gorm.io/gorm"
)

type StockRepository interface {
	// Stock operations
	BaseRepository[entity.Stock]
	FindAvailableStockFIFO(ctx context.Context, tx *gorm.DB, productId int) ([]*entity.Stock, error)
	DecreaseStockQtyTx(ctx context.Context, tx *gorm.DB, stockId int, qty int) error
}
