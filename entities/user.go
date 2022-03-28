package entities

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name     string `json:"name" form:"name"`
	Email    string `json:"email" form:"email" gorm:"unique"`
	Password string `json:"password" form:"password"`
	Adress   string `json:"adress" form:"adress"`
	Books    []Book
	Rents    []Rent
}
