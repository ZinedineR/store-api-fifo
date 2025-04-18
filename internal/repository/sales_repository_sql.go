package repository

import (
	"boiler-plate-clean/internal/entity"
	"boiler-plate-clean/internal/model"
	"context"
	"gorm.io/gorm"
)

type SaleSQLRepo struct {
	BaseRepository[entity.Sale]
}

func NewSaleSQLRepository() SaleRepository {
	relationField := []RelationField{
		{
			Name: "Product",
			Func: func(db *gorm.DB) *gorm.DB {
				return db
			},
		},
	}
	repo := NewBaseRepositoryImpl[entity.Sale](relationField)
	return &SaleSQLRepo{
		BaseRepository: repo,
	}
}
func (r *SaleSQLRepo) GetProfitReport(
	ctx context.Context, tx *gorm.DB, month, year int,
) (*model.ProfitReportRes, error) {
	var res model.ProfitReportRes
	query := `
		SELECT
			EXTRACT(MONTH FROM created_at) AS month,
			EXTRACT(YEAR FROM created_at) AS year,
			SUM(quantity * price) AS total_penjualan,
			SUM(total_hpp) AS total_hpp,
			SUM(quantity * price) - SUM(total_hpp) AS laba
		FROM sales
		WHERE EXTRACT(MONTH FROM created_at) = ? AND EXTRACT(YEAR FROM created_at) = ?
		GROUP BY year, month
	`

	if err := tx.WithContext(ctx).Raw(query, month, year).Scan(&res).Error; err != nil {
		return nil, err
	}
	return &res, nil
}
