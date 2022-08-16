package routers

import (
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/sirupsen/logrus"
	"github.com/yuya-lock/contessa/api/controller"
	"github.com/yuya-lock/contessa/api/models"
)

func Init() {
	if err := godotenv.Load(".env"); err != nil {
		logrus.Fatal("Error while loading .env file.\n%s", err)
	}

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Use(middleware.CORS())
	e.GET("/", controller.InputCocktailData)
	e.GET("/cocktails", controller.GetCocktails)
	e.GET("/cocktails/:id", controller.GetCocktailDetail)
	e.POST("/signup", controller.Signup)
	e.POST("/login", controller.Login)
	e.POST("/comment", controller.CreateCocktailComment)
	e.POST("/createlike", controller.CreateCocktailLike)
	e.POST("/deletelike", controller.DeleteCocktailLike)
	e.POST("/rate", controller.CreateCocktailRate)

	config := middleware.JWTConfig{
		Claims:     &models.JwtCustomClaims{},
		SigningKey: []byte("secret"),
	}

	r := e.Group("/restricted")
	r.Use(middleware.JWTWithConfig(config))
	r.GET("/mypage", controller.GetUser)
	r.GET("/favoritecocktails", controller.GetFavoriteCocktails)
	r.GET("/mycomments", controller.GetMyComments)
	r.GET("/myrates", controller.GetMyRates)

	e.Logger.Fatal(e.Start(":8000"))
}
