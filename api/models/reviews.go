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
		ID         uint
		UserID     uint
		CocktailID uint
	}

	Rate struct {
		ID         uint
		Rating     uint `json:"rating" form:"rating" query:"rating"`
		UserID     uint
		CocktailID uint
	}
)
