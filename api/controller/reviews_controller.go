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
	db.Where("user_id = ? AND cocktail_id = ?", input.UserID, input.CocktailID).First(&comment)
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

func GetCocktailComment(c echo.Context) error {
	id := c.Param("id")

	db, sqlDB, err := database.Connect()
	if err != nil {
		return err
	}
	defer sqlDB.Close()

	comment := models.Comment{}
	db.Where("cocktail_id = ?", id).First(&comment)

	return c.JSON(http.StatusOK, comment)
}

func UpdateCocktailComment(c echo.Context) error {
	input := new(models.Comment)
	if err := c.Bind(input); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	db, sqlDB, err := database.Connect()
	if err != nil {
		return err
	}
	defer sqlDB.Close()

	comment := models.Comment{}
	db.Model(&comment).Where("user_id = ? AND cocktail_id = ?", input.UserID, input.CocktailID).Update("body", input.Body)

	return c.JSON(http.StatusOK, echo.Map{
		"message": "Your comment was updated successfully.",
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

func GetCocktailLike(c echo.Context) error {
	id := c.Param("id")

	db, sqlDB, err := database.Connect()
	if err != nil {
		return err
	}
	defer sqlDB.Close()

	like := models.Like{}
	db.Where("user_id = ?", id).First(&like)

	return c.JSON(http.StatusOK, like)
}

func UpdateCocktailLike(c echo.Context) error {
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

func GetCocktailRate(c echo.Context) error {
	id := c.Param("id")

	db, sqlDB, err := database.Connect()
	if err != nil {
		return err
	}
	defer sqlDB.Close()

	rate := models.Rate{}
	db.Where("cocktail_id = ?", id).First(&rate)

	return c.JSON(http.StatusOK, rate)
}

func UpdateCocktailRate(c echo.Context) error {
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
	db.Model(&rate).Where("user_id = ? AND cocktail_id = ?", input.UserID, input.CocktailID).Update("rating", input.Rating)

	return c.JSON(http.StatusOK, echo.Map{
		"message": "Your rate was updated successfully.",
	})
}
