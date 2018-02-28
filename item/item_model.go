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
	Avarar      string
	Status      bool
	Quantity    uint
	UnitID      uint
	SupplierID  uint
}
