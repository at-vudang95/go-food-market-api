package category

import (
	"github.com/jinzhu/gorm"
)

// Category model
type Category struct {
	gorm.Model
	Name     string
	Level    int
	ParentID uint
}
