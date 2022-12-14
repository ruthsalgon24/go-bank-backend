package interfaces

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	Username string
	Email    string
	Password string
}

type Account struct {
	gorm.Model
	Type    string
	Name    string
	Balance float64
	UserID  uint
}

// Create transaction interface
type Transaction struct {
	gorm.Model
	From   uint
	To     uint
	Amount float64
}

type ResponseTransaction struct {
	ID     uint
	From   uint
	To     uint
	Amount float64
}

type ResponseAccount struct {
	ID      uint
	Name    string
	Balance float64
}

type ResponseUser struct {
	ID       uint
	Username string
	Email    string
	Accounts []ResponseAccount
}

// Create Validation interface
type Validation struct {
	Value string
	Valid string
}

type ErrResponse struct {
	Message string
}
