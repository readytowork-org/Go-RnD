package interfaces

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Username string
	Password string
	Email    string
}

type Account struct {
	gorm.Model
	Type    string
	Name    string
	Balance uint
	UserId  uint
}

type ResponseAccount struct {
	ID      uint
	Name    string
	Balance int
}

type ResonseUser struct {
	ID       uint
	Username string
	Email    string
	Accounts []ResponseAccount
}
