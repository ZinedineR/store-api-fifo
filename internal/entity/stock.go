package entity

import "time"

type Stock struct {
	ID        int       `gorm:"primaryKey;not null;autoIncrement" json:"id"`
	Price     float64   `json:"price" validate:"required,gt=0"`
	Quantity  int       `json:"quantity" validate:"required,gt=0"`
	ProductId int       `json:"product_id" validate:"required"`
	Product   *Product  `json:"product" gorm:"foreignKey:ProductId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	CreatedAt time.Time `gorm:"<-:create;" json:"created_at"`
}

func (Stock) TableName() string {
	return "stock"
}
