package migration

import (
	"boiler-plate-clean/internal/entity"
	"boiler-plate-clean/pkg/database"
)

func AutoMigration(CpmDB *database.Database) {
	CpmDB.MigrateDB(
		&entity.Product{},
		&entity.Stock{},
		&entity.Sale{})
	//&entity.SMSLog{}
}
