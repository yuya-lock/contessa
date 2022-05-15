package models

import "gorm.io/gorm"

type (
	Comment struct {
		gorm.Model
		Body       string `json:"body" form:"body" query:"body"`
		UserID     uint
		CocktailID uint
	}

	Like struct {
		gorm.Model
		UserID     uint
		CocktailID uint
	}

	Rate struct {
		gorm.Model
		Rating     uint `json:"rating" form:"rating" query:"rating"`
		UserID     uint
		CocktailID uint
	}
)
