package main

import (
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

func init() {
	logrus.SetLevel(logrus.DebugLevel)
	logrus.SetFormatter(&logrus.JSONFormatter{})
}

func main() {
	e := echo.New()
	e.GET("/", index)
	e.GET("/cocktails", fetchCocktails)
	e.GET("/cocktails/:id", fetchCocktailDetail)

	e.Logger.Fatal(e.Start(":8000"))
}
