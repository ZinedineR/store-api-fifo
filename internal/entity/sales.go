package entity

import "time"

type Sale struct {
	ID        int       `gorm:"primaryKey;not null;autoIncrement" json:"id"`
	ProductId int       `json:"product_id"`
	Product   *Product  `gorm:"foreignKey:ProductId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"product"`
	Quantity  int       `json:"quantity"`
	Price     float64   `json:"price"`
	TotalHPP  float64   `json:"total_hpp"`
	CreatedAt time.Time `gorm:"<-:create;" json:"created_at"`
}
