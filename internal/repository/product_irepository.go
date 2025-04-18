package repository

import (
	"boiler-plate-clean/internal/entity"
)

type ProductRepository interface {
	// Product operations
	BaseRepository[entity.Product]
}
