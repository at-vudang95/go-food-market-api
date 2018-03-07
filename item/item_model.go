package item

import (
	"github.com/jinzhu/gorm"
)

// Item struct
type Item struct {
	gorm.Model
	Name        string
	Description string
	Price       float64
	Avatar      string
	Status      bool
	Quantity    uint
	SupplierID  uint
	CategoryID  uint
}
