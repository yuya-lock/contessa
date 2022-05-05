package lib

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string `json:"name" form:"name" query:"name" gorm:"unique"`
	Password []byte `json:"password" form:"password" query:"password"`
	Age      string `json:"age" form:"age" query:"age"`
	Job      string `json:"job" form:"job" query:"job"`
	Sex      string `json:"sex" form:"sex" query:"sex"`
}
