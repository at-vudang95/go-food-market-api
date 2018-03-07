package supplier

import (
	"github.com/at-vudang95/go-food-market-api/user"
	"github.com/jinzhu/gorm"
)

// Supplier model
type Supplier struct {
	gorm.Model
	Amount  float64
	UserID  uint
	Account user.Account
}
