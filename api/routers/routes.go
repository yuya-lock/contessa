package routers

import (
	"github.com/labstack/echo/v4"
	"github.com/yuya-lock/contessa/api/controller"
)

func Init() {
	e := echo.New()
	e.GET("/", controller.Index)
	e.GET("/cocktails", controller.FetchCocktails)
	e.GET("/cocktails/:id", controller.FetchCocktailDetail)

	//e.POST("/accounts", controller.createUser)
	//e.POST("/login", controller.Login)
	//e.GET("/user", controller.getUser)
	//e.GET("logout", controller.Logout)
	e.POST("/register", controller.Register)

	//r := e.Group("/restricted")
	//r.Use(middleware.JWTWithConfig(Config))
	//r.GET("", controller.Restricted)

	e.Logger.Fatal(e.Start(":8000"))
}
