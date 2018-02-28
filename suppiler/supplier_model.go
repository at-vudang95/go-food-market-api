package supplier

import (
	"github.com/jinzhu/gorm"
)

// Supplier model
type Supplier struct {
	gorm.Model
	Amount float64
	UserID uint
}
