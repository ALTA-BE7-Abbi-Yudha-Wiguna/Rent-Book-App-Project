package entities

import "gorm.io/gorm"

type Book struct {
	gorm.Model
	UserID    uint   `json:"UserID" form:"UserID"`
	Title     string `json:"title" form:"title"`
	Author    string `json:"author" form:"author"`
	Publisher string `json:"publisher" form:"publisher"`
	Status    string `json:"status" from:"status" gorm:"default:Available"`
	Rents     []Rent
}
