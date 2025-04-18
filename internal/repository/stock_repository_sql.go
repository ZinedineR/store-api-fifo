package repository

import (
	"boiler-plate-clean/internal/entity"
	"boiler-plate-clean/internal/model"
	"context"
	"gorm.io/gorm"
	"strconv"
)

type StockSQLRepo struct {
	BaseRepository[entity.Stock]
}

func NewStockSQLRepository() StockRepository {
	relationField := []RelationField{
		{
			Name: "Product",
			Func: func(db *gorm.DB) *gorm.DB {
				return db
			},
		},
	}
	repo := NewBaseRepositoryImpl[entity.Stock](relationField)
	return &StockSQLRepo{
		repo,
	}
}

func (r *StockSQLRepo) FindAvailableStockFIFO(
	ctx context.Context, tx *gorm.DB, productId int,
) ([]*entity.Stock, error) {
	return r.Find(ctx, tx, model.OrderParam{
		Order:   "asc",
		OrderBy: "created_at",
	},
		model.FilterParams{
			{
				Field:    "product_id",
				Value:    strconv.Itoa(productId),
				Operator: "=",
			},
			{
				Field:    "quantity",
				Value:    "0",
				Operator: ">",
			},
		})
}

func (r *StockSQLRepo) DecreaseStockQtyTx(
	ctx context.Context, tx *gorm.DB, stockId int, qty int,
) error {
	return tx.WithContext(ctx).
		Model(&entity.Stock{}).
		Where("id = ? AND quantity >= ?", stockId, qty).
		Update("quantity", gorm.Expr("quantity - ?", qty)).Error
}
