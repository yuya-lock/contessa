package models

import "gorm.io/gorm"

type (
	Comment struct {
		gorm.Model
		Body       string `json:"body" form:"body" query:"body"`
		UserID     uint   `json:"user_id" form:"user_id" query:"user_id"`
		CocktailID uint   `json:"cocktail_id" form:"cocktail_id" query:"cocktail_id"`
	}

	Like struct {
		gorm.Model
		UserID     uint `json:"user_id" form:"user_id" query:"user_id"`
		CocktailID uint `json:"cocktail_id" form:"cocktail_id" query:"cocktail_id"`
	}

	Rate struct {
		gorm.Model
		Rating     uint `json:"rating" form:"rating" query:"rating"`
		UserID     uint `json:"user_id" form:"user_id" query:"user_id"`
		CocktailID uint `json:"cocktail_id" form:"cocktail_id" query:"cocktail_id"`
	}
)
