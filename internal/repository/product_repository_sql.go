package repository

import (
	"boiler-plate-clean/internal/entity"
)

type ProductSQLRepo struct {
	BaseRepository[entity.Product]
}

func NewProductSQLRepository() ProductRepository {
	repo := NewBaseRepositoryImpl[entity.Product](nil)
	return &ProductSQLRepo{
		repo,
	}
}
