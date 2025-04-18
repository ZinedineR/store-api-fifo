package entity

import "time"

type Product struct {
	ID        int       `gorm:"primaryKey;not null;autoIncrement" json:"id"`
	Name      string    `json:"name" validate:"required"`
	CreatedAt time.Time `gorm:"<-:create;" json:"created_at"`
}

func (Product) TableName() string {
	return "product"
}
