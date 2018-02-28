package order

import (
	"time"

	"github.com/jinzhu/gorm"
)

// Order model
type Order struct {
	gorm.Model
	TransportAt time.Time
	Address     string
	Status      byte
	Name        string
	Phone       string
	Node        string
	UserID      uint
	PromotionID uint
	ShipID      uint
}

// Item model
type Item struct {
	gorm.Model
	PriceOffical float64
	Quantity     uint
	ItemID       uint
	OrderID      uint
}
