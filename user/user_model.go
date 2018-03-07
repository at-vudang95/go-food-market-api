package user

import (
	"github.com/jinzhu/gorm"
)

// User struct
type User struct {
	gorm.Model
	Name      string
	Phone     string
	Email     string
	AccountID uint
	Account   Account
}

// Account struct
type Account struct {
	gorm.Model
	Username string
	Password string
}
