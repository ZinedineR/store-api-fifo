package repository

import (
	"boiler-plate-clean/internal/entity"
	"boiler-plate-clean/internal/model"
	"context"
	"gorm.io/gorm"
)

type SaleRepository interface {
	// Sale operations
	BaseRepository[entity.Sale]
	GetProfitReport(ctx context.Context, tx *gorm.DB, month, year int) (*model.ProfitReportRes, error)
}
