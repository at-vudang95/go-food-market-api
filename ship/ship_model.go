package ship

import (
	"time"

	"github.com/jinzhu/gorm"
)

// Ship model
type Ship struct {
	gorm.Model
	Price     uint64
	UserID    uint
	Address   string
	DeliverAT time.Time
}
