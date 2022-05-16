package controller

import (
	"github.com/labstack/echo/v4"
	"github.com/yuya-lock/contessa/api/database"
	"github.com/yuya-lock/contessa/api/models"
	"net/http"
)

func CreateCocktailComment(c echo.Context) error {
	input := new(models.Comment)
	if err := c.Bind(input); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	db, sqlDB, err := database.Connect()
	if err != nil {
		return c.String(http.StatusServiceUnavailable, "")
	}
	defer sqlDB.Close()
	database.InitDB()

	comment := models.Comment{}
	db.Where("user_id = ? AND cocktail_id = ?", input.UserID, input.CocktailID).Find(&comment)
	if comment.ID != 0 {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": "This comment already exists.",
		})
	}
	comment = models.Comment{
		Body:       input.Body,
		UserID:     input.UserID,
		CocktailID: input.CocktailID,
	}
	db.Create(&comment)

	return c.JSON(http.StatusOK, echo.Map{
		"message": "Your comment was created successfully",
	})
}

func CreateCocktailLike(c echo.Context) error {
	input := new(models.Comment)
	if err := c.Bind(input); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	db, sqlDB, err := database.Connect()
	if err != nil {
		return err
	}
	defer sqlDB.Close()

	like := models.Like{}
	db.Where("user_id = ? AND cocktail_id = ?", input.UserID, input.CocktailID).First(&like)
	if like.ID != 0 {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": "This like already exists.",
		})
	}
	like = models.Like{
		UserID:     input.UserID,
		CocktailID: input.CocktailID,
	}
	db.Create(&like)

	return c.JSON(http.StatusOK, echo.Map{
		"message": "Your like was created successfully",
	})
}

func DeleteCocktailLike(c echo.Context) error {
	input := new(models.Comment)
	if err := c.Bind(input); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	db, sqlDB, err := database.Connect()
	if err != nil {
		return err
	}
	defer sqlDB.Close()

	like := models.Like{}
	db.Where("user_id = ? AND cocktail_id = ?", input.UserID, input.CocktailID).First(&like)
	db.Delete(&like)

	return c.JSON(http.StatusOK, echo.Map{
		"message": "Your like was deleted successfully.",
	})
}

func CreateCocktailRate(c echo.Context) error {
	input := new(models.Rate)
	if err := c.Bind(input); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	db, sqlDB, err := database.Connect()
	if err != nil {
		return err
	}
	defer sqlDB.Close()

	rate := models.Rate{}
	db.Where("user_id = ? AND cocktail_id = ?", input.UserID, input.CocktailID).First(&rate)
	if rate.ID != 0 {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": "This rate already exists.",
		})
	}
	rate = models.Rate{
		Rating:     input.Rating,
		UserID:     input.UserID,
		CocktailID: input.CocktailID,
	}
	db.Create(&rate)

	return c.JSON(http.StatusOK, echo.Map{
		"message": "Your rate was created successfully",
	})
}
