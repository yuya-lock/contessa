package models

import (
	"github.com/dgrijalva/jwt-go"
	"gorm.io/gorm"
)

type (
	User struct {
		gorm.Model
		Name     string `json:"name" form:"name" query:"name" gorm:"unique"`
		Password []byte `json:"password" form:"password" query:"password"`
		Age      string `json:"age" form:"age" query:"age"`
		Job      string `json:"job" form:"job" query:"job"`
		Sex      string `json:"sex" form:"sex" query:"sex"`
	}

	JwtCustomClaims struct {
		UID  uint   `json:"uid"`
		Name string `json:"name"`
		jwt.StandardClaims
	}
)
