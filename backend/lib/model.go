package lib

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name  string
	Email string
}
