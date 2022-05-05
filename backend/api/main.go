package main

import (
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.GET("/", index)
	e.GET("/cocktails", fetchCocktails)
	e.GET("/cocktails/:id", fetchCocktailDetail)

	//e.POST("/accounts", createUser)
	//e.POST("/login", Login)
	//e.GET("/user", getUser)
	//e.GET("logout", Logout)
	e.POST("/register", Register)

	//r := e.Group("/restricted")
	//r.Use(middleware.JWTWithConfig(Config))
	//r.GET("", Restricted)

	e.Logger.Fatal(e.Start(":8000"))
}
