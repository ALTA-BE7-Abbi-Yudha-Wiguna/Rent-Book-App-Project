package entities

import (
	"gorm.io/gorm"
	"time"
)

type Rent struct {
	gorm.Model
	Id                  int       `json:"id"`
	UserID              uint      `json:"userID" form:"userID"`
	BookID              uint      `json:"bookID" form:"bookID"`
	StatusRent          string    `json:"statusRent" form:"statusRent" gorm:"default:Belum Dikembalikan" `
	TanggalPengembalian time.Time `json:"tanggalPengembalian" form:"tanggalPengembalian"`
}
