package models

import (
	"gorm.io/gorm"
)

type (
	User struct {
		gorm.Model
		Name     string `json:"name" form:"name" query:"name" gorm:"unique;not null"`
		Password []byte `json:"password" form:"password" query:"password" gorm:"not null"`
		Age      string `json:"age" form:"age" query:"age"`
		Job      string `json:"job" form:"job" query:"job"`
		Sex      string `json:"sex" form:"sex" query:"sex"`
		Comments []Comment
		Likes    []Like
		Rates    []Rate
	}

	//JwtCustomClaims struct {
	//	UID  uint   `json:"uid"`
	//	Name string `json:"name"`
	//	jwt.StandardClaims
	//}
)
