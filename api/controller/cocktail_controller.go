package controller

import (
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"github.com/yuya-lock/contessa/api/database"
	"github.com/yuya-lock/contessa/api/models"
	"github.com/yuya-lock/contessa/api/rest"
	"net/http"
	"strconv"
)

func GetCocktails(c echo.Context) error {
	logrus.Info("Get cocktails")

	input := new(models.CocktailsInput)
	if err := c.Bind(input); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	responseBody, err := rest.FetchCocktails(input)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, responseBody)
}

func GetCocktailDetail(c echo.Context) error {
	id := c.Param("id")

	db, sqlDB, err := database.Connect()
	if err != nil {
		return err
	}
	defer sqlDB.Close()

	cocktail := models.Cocktail{}
	db.Where("cocktail_id = ?", id).Preload("Recipes").Preload("Comments").Preload("Likes").Preload("Rates").Find(&cocktail)

	return c.JSON(http.StatusOK, cocktail)
}

func InputCocktailData(c echo.Context) error {
	logrus.Info("Save cocktail info to database")

	db, sqlDB, err := database.Connect()
	if err != nil {
		return err
	}
	defer sqlDB.Close()
	database.InitDB()

	for page := 1; page <= 3; page++ {
		resp, err := rest.FetchCocktailData(strconv.Itoa(page))
		if err != nil {
			return err
		}
		for _, v := range resp.Cocktails {
			cocktail := models.Cocktail{
				CocktailId:          v.CocktailId,
				CocktailName:        v.CocktailName,
				CocktailNameEnglish: v.CocktailNameEnglish,
				BaseName:            v.BaseName,
				TechniqueName:       v.TechniqueName,
				TasteName:           v.TasteName,
				StyleName:           v.StyleName,
				Alcohol:             v.Alcohol,
				TopName:             v.TopName,
				GlassName:           v.GlassName,
				TypeName:            v.TypeName,
				CocktailDigest:      v.CocktailDigest,
				CocktailDesc:        v.CocktailDesc,
				RecipeDesc:          v.RecipeDesc,
				Recipes:             v.Recipes,
			}
			for _, u := range v.Recipes {
				recipe := models.Recipe{
					IngredientId:   u.IngredientId,
					IngredientName: u.IngredientName,
					Amount:         u.Amount,
					Unit:           u.Unit,
				}
				db.Create(&recipe)
			}
			db.Create(&cocktail)
		}
	}

	return c.JSON(http.StatusOK, echo.Map{
		"message": "Save cocktail info to database",
	})
}
