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
	e.GET("/", controller.InputCocktailData)
	e.GET("/cocktails", controller.GetCocktails)
	e.GET("/cocktails/:id", controller.GetCocktailDetail)
	e.POST("/signup", controller.Signup)
	e.POST("/login", controller.Login)
	e.POST("/comment", controller.CreateCocktailComment)
	e.GET("/comment/:id", controller.GetCocktailComment)
	e.PUT("/comment", controller.UpdateCocktailComment)
	e.POST("like", controller.CreateCocktailLike)
	e.GET("/like/:id", controller.GetCocktailLike)
	e.PUT("/like", controller.UpdateCocktailLike)
	e.POST("/comment", controller.CreateCocktailRate)
	e.GET("/comment/:id", controller.GetCocktailRate)
	e.PUT("/comment", controller.UpdateCocktailRate)

	//config := middleware.JWTConfig{
	//	Claims:     &models.JwtCustomClaims{},
	//	SigningKey: []byte("secret"),
	//}
	//
	//r := e.Group("/restricted")
	//r.Use(middleware.JWTWithConfig(config))
	//r.GET("", controller.Restricted)

	e.Logger.Fatal(e.Start(":8000"))
}
