package routers

import (
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"github.com/yuya-lock/contessa/api/controller"
)

func Init() {
	err := godotenv.Load(".env")
	if err != nil {
		logrus.Fatal("Error while loading .env file.\n%s", err)
	}

	e := echo.New()
	e.GET("/", controller.Index)
	e.GET("/cocktails", controller.GetCocktails)
	e.GET("/cocktails/:id", controller.GetCocktailDetail)

	e.POST("/signup", controller.Signup)
	e.POST("/login", controller.Login)
	e.GET("/user", controller.User)
	e.GET("logout", controller.Logout)

	e.Logger.Fatal(e.Start(":8000"))
}
