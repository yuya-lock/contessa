package models

import "gorm.io/gorm"

type (
	Comment struct {
		gorm.Model
		body       string
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
		Rating     uint
		UserID     uint
		CocktailID uint
	}
)
