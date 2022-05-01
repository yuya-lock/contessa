package lib

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name      string    `json:"name" gorm:"type:varchar(255);not null"`
	Age       uint      `json:"age" gorm:"not null;default:0"`
	Job       string    `json:"job"`
	Sex       string    `json:"sex"`
}
