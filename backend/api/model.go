package main

import (
	"fmt"
	"unicode/utf8"
)

type (
	GetCocktailsInput struct {
		Word        string   `query:"word"`
		Base        string   `query:"base"`
		Technique   string   `query:"technique"`
		Taste       string   `query:"taste"`
		Style       string   `query:"style"`
		AlcoholFrom string   `query:"alcohol_from"`
		AlcoholTo  string    `query:"alcohol_to"`
		Top         string   `query:"top"`
		Page        string   `query:"page"`
		Limit       string   `query:"limit"`
	}

	// GetCocktailsOutputItemRecipes struct {
	// 	Ingredient_id   uint   `json:"ingredient_id"`
	// 	Ingredient_name string `json:"ingredient_name"`
	// 	Amount          string `json:"amount"`
	// 	Unit            string `json:"unit"`
	// }

	// GetCocktailsOutputItem struct {
	// 	Cocktail_id           uint                            `json:"cocktail_id"`
	// 	Cocktail_name         string                          `json:"cocktail_name"`
	// 	Cocktail_name_english string                          `json:"cocktail_name_english"`
	// 	Base_name             string                          `json:"base_name"`
	// 	Technique_name        string                          `json:"technique_name"`
	// 	Taste_name            string                          `json:"taste_name"`
	// 	Style_name            string                          `json:"style_name"`
	// 	Alcohol               uint                            `json:"alcohol"`
	// 	Top_name              string                          `json:"top_name"`
	// 	Glass_name            string                          `json:"glass_name"`
	// 	Type_name             string                          `json:"type_name"`
	// 	Cocktail_digest       string                          `json:"cocktail_digest"`
	// 	Cocktail_desc         string                          `json:"cocktail_desc"`
	// 	Recipe_desc           string                          `json:"recipe_desc"`
	// 	Recipes               []GetCocktailsOutputItemRecipes `json:"recipes"`
	// }

	// GetCocktailsOutput struct {
	// 	Cocktails     []GetCocktailsOutputItem `json:"cocktails"`
	// 	Status        string                   `json:"status"`
	// 	Total_pages   uint                     `json:"total_pages"`
	// 	Current_page  uint                     `json:"current_page"`
	// }
)

func (input *GetCocktailsInput) validator() error {
	if utf8.RuneCountInString(input.Word) > 20{
		return fmt.Errorf("Character count exceeds the limit of 20 characters")
	}
	return nil
}
