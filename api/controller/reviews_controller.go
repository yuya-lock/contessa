package controller

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"github.com/yuya-lock/contessa/api/database"
	"github.com/yuya-lock/contessa/api/models"
	"net/http"
	"strings"
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

	cocktail := models.Cocktail{}
	db.Where("cocktail_id = ?", input.CocktailID).First(&cocktail)

	comment := models.Comment{}
	db.Where("user_id = ? AND cocktail_id = ?", input.UserID, input.CocktailID).Find(&comment)
	if comment.ID != 0 {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": "Your comment already exists.",
		})
	}
	comment = models.Comment{
		Body:       input.Body,
		UserID:     input.UserID,
		CocktailID: input.CocktailID,
	}
	db.Model(&cocktail).Association("Comments").Append(&comment)

	return c.JSON(http.StatusOK, echo.Map{
		"message": "Your comment was created successfully",
	})
}

func GetMyComments(c echo.Context) error {
	auth := c.Request().Header.Get("Authorization")
	if auth == "" {
		return c.JSON(http.StatusUnauthorized, echo.Map{
			"Message": "Authorization Header Not Found",
		})
	}
	splitToken := strings.Split(auth, "Bearer ")
	auth = splitToken[1]
	u, err := jwt.ParseWithClaims(auth, &models.JwtCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte("secret"), nil
	})
	if err != nil {
		return c.JSON(http.StatusNotFound, echo.Map{
			"Message": "Token is wrong or Expire",
		})
	}
	claims := u.Claims.(*models.JwtCustomClaims)

	db, sqlDB, err := database.Connect()
	if err != nil {
		return c.String(http.StatusServiceUnavailable, "")
	}
	defer sqlDB.Close()

	user := models.User{}
	db.Where("name = ?", claims.Name).Preload("Comments").Preload("Likes").Preload("Rates").First(&user)
	if user.ID == 0 {
		return c.JSON(http.StatusNotFound, echo.Map{
			"message": "User not found",
		})
	}

	comments := user.Comments
	var cocktails []models.Cocktail
	var myCommentCocktailsId []uint
	for _, comment := range comments {
		myCommentCocktailsId = append(myCommentCocktailsId, comment.CocktailID)
	}
	db.Where("cocktail_id in ?", myCommentCocktailsId).Find(&cocktails)

	return c.JSON(http.StatusOK, echo.Map{
		"myCommentCocktails": cocktails,
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

	cocktail := models.Cocktail{}
	db.Where("cocktail_id = ?", input.CocktailID).First(&cocktail)

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
	db.Model(&cocktail).Association("Likes").Append(&like)

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

func GetFavoriteCocktails(c echo.Context) error {
	auth := c.Request().Header.Get("Authorization")
	if auth == "" {
		return c.JSON(http.StatusUnauthorized, echo.Map{
			"Message": "Authorization Header Not Found",
		})
	}
	splitToken := strings.Split(auth, "Bearer ")
	auth = splitToken[1]
	u, err := jwt.ParseWithClaims(auth, &models.JwtCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte("secret"), nil
	})
	if err != nil {
		return c.JSON(http.StatusNotFound, echo.Map{
			"Message": "Token is wrong or Expire",
		})
	}
	claims := u.Claims.(*models.JwtCustomClaims)

	db, sqlDB, err := database.Connect()
	if err != nil {
		return c.String(http.StatusServiceUnavailable, "")
	}
	defer sqlDB.Close()

	user := models.User{}
	db.Where("name = ?", claims.Name).Preload("Comments").Preload("Likes").Preload("Rates").First(&user)
	if user.ID == 0 {
		return c.JSON(http.StatusNotFound, echo.Map{
			"message": "User not found",
		})
	}

	likes := user.Likes
	var cocktails []models.Cocktail
	var favoriteCocktailsId []uint
	for _, like := range likes {
		favoriteCocktailsId = append(favoriteCocktailsId, like.CocktailID)
	}
	db.Where("cocktail_id in ?", favoriteCocktailsId).Find(&cocktails)

	return c.JSON(http.StatusOK, echo.Map{
		"favoriteCocktails": cocktails,
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

	cocktail := models.Cocktail{}
	db.Where("cocktail_id = ?", input.CocktailID).First(&cocktail)

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
	db.Model(&cocktail).Association("Rates").Append(&rate)

	return c.JSON(http.StatusOK, echo.Map{
		"message": "Your rate was created successfully",
	})
}

func GetMyRates(c echo.Context) error {
	auth := c.Request().Header.Get("Authorization")
	if auth == "" {
		return c.JSON(http.StatusUnauthorized, echo.Map{
			"Message": "Authorization Header Not Found",
		})
	}
	splitToken := strings.Split(auth, "Bearer ")
	auth = splitToken[1]
	u, err := jwt.ParseWithClaims(auth, &models.JwtCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte("secret"), nil
	})
	if err != nil {
		return c.JSON(http.StatusNotFound, echo.Map{
			"Message": "Token is wrong or Expire",
		})
	}
	claims := u.Claims.(*models.JwtCustomClaims)

	db, sqlDB, err := database.Connect()
	if err != nil {
		return c.String(http.StatusServiceUnavailable, "")
	}
	defer sqlDB.Close()

	user := models.User{}
	db.Where("name = ?", claims.Name).Preload("Comments").Preload("Likes").Preload("Rates").First(&user)
	if user.ID == 0 {
		return c.JSON(http.StatusNotFound, echo.Map{
			"message": "User not found",
		})
	}

	rates := user.Rates
	var cocktails []models.Cocktail
	var myRateCocktailsId []uint
	for _, rate := range rates {
		myRateCocktailsId = append(myRateCocktailsId, rate.CocktailID)
	}
	db.Where("cocktail_id in ?", myRateCocktailsId).Find(&cocktails)

	return c.JSON(http.StatusOK, echo.Map{
		"myRateCocktails": cocktails,
	})
}
