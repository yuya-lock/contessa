package routers

import (
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"github.com/yuya-lock/contessa/api/controller"
)

func Init() {
	if err := godotenv.Load(".env"); err != nil {
		logrus.Fatal("Error while loading .env file.\n%s", err)
	}

	e := echo.New()
	e.GET("/", controller.Index)
	e.GET("/cocktails", controller.GetCocktails)
	e.GET("/cocktails/:id", controller.GetCocktailDetail)
	e.POST("/signup", controller.Signup)
	e.POST("/login", controller.Login)
	e.GET("/mypage", controller.Mypage)
	e.GET("logout", controller.Logout)

	e.Logger.Fatal(e.Start(":8000"))
}
