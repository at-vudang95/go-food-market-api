package image

import (
	"github.com/jinzhu/gorm"
)

// Image model
type Image struct {
	gorm.Model
	URL    string
	ItemID uint
}
