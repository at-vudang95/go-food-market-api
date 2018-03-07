package promotion

import (
	"time"

	"github.com/jinzhu/gorm"
)

// Promotion model
type Promotion struct {
	gorm.Model
	Title  string
	FromAt time.Time
	EndAt  time.Time
}

// Item model
type Item struct {
	gorm.Model
	Percent     int
	PromotionID uint
	ItemID      uint
}
