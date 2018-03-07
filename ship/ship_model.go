package ship

import (
	"time"

	"github.com/jinzhu/gorm"
)

// Ship model
type Ship struct {
	gorm.Model
	Status    byte
	OrderID   uint
	AddressID uint
	DeliverAT time.Time
}
