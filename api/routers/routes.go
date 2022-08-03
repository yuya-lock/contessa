package routers

import (
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/sirupsen/logrus"
	"github.com/yuya-lock/contessa/api/controller"
	"github.com/yuya-lock/contessa/api/models"
)

func MethodOverride(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		if c.Request().Method == "POST" {
			method := c.Request().PostFormValue("_method")
			if method == "PUT" || method == "PATCH" || method == "DELETE" {
				c.Request().Method = method
			}
		}
		return next(c)
	}
}

func Init() {
	if err := godotenv.Load(".env"); err != nil {
		logrus.Fatal("Error while loading .env file.\n%s", err)
	}

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Use(middleware.CORS())
	e.Pre(MethodOverride)
	e.GET("/", controller.InputCocktailData)
	e.GET("/cocktails", controller.GetCocktails)
	e.GET("/cocktails/:id", controller.GetCocktailDetail)
	e.POST("/signup", controller.Signup)
	e.POST("/login", controller.Login)
	e.POST("/comment", controller.CreateCocktailComment)
	e.POST("/like", controller.CreateCocktailLike)
	e.DELETE("/like", controller.DeleteCocktailLike)
	e.POST("/rate", controller.CreateCocktailRate)

	config := middleware.JWTConfig{
		Claims:     &models.JwtCustomClaims{},
		SigningKey: []byte("secret"),
	}

	r := e.Group("/restricted")
	r.Use(middleware.JWTWithConfig(config))
	r.GET("", controller.Restricted)

	e.Logger.Fatal(e.Start(":8080"))
}
