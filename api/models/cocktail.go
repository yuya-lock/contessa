package models

type (
	Cocktail struct {
		CocktailId          uint      `json:"cocktail_id" gorm:"primary_key"`
		CocktailName        string    `json:"cocktail_name"`
		CocktailNameEnglish string    `json:"cocktail_name_english"`
		BaseName            string    `json:"base_name"`
		TechniqueName       string    `json:"technique_name"`
		TasteName           string    `json:"taste_name"`
		StyleName           string    `json:"style_name"`
		Alcohol             uint      `json:"alcohol"`
		TopName             string    `json:"top_name"`
		GlassName           string    `json:"glass_name"`
		TypeName            string    `json:"type_name"`
		CocktailDigest      string    `json:"cocktail_digest"`
		CocktailDesc        string    `json:"cocktail_desc"`
		RecipeDesc          string    `json:"recipe_desc"`
		Recipes             []Recipe  `json:"recipes" gorm:"many2many:cocktail_recipes;"`
		Comments            []Comment `json:"comments" gorm:"many2many:cocktail_comments;"`
		Likes               []Like    `json:"likes" gorm:"many2many:cocktail_likes;"`
		Rates               []Rate    `json:"rates" gorm:"many2many:cocktail_rates;"`
	}

	Recipe struct {
		IngredientId   uint   `json:"ingredient_id" gorm:"primary_key"`
		IngredientName string `json:"ingredient_name"`
		Amount         string `json:"amount"`
		Unit           string `json:"unit"`
	}

	CocktailsInput struct {
		Word        string `query:"word"`
		Base        string `query:"base"`
		Technique   string `query:"technique"`
		Taste       string `query:"taste"`
		Style       string `query:"style"`
		AlcoholFrom string `query:"alcohol_from"`
		AlcoholTo   string `query:"alcohol_to"`
		Top         string `query:"top"`
		Page        string `query:"page"`
		Limit       string `query:"limit"`
	}

	CocktailsOutput struct {
		Cocktails   []Cocktail `json:"cocktails"`
		Status      string     `json:"status"`
		TotalPages  uint       `json:"total_pages"`
		CurrentPage uint       `json:"current_page"`
	}
)
