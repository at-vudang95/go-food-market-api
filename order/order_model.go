package order

import (
	"time"

	"github.com/jinzhu/gorm"
)

// Order model
type Order struct {
	gorm.Model
	TransportAt time.Time
	Status      byte
	Node        string
	UserID      uint
	PromotionID uint
}

// Item model
type Item struct {
	gorm.Model
	PriceOffical float64
	Quantity     uint
	ItemID       uint
	OrderID      uint
}
