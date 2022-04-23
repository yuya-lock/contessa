package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func index(c echo.Context) error {
	return c.String(http.StatusOK, "healthy")
}

func getLargeArea(c echo.Context) error {
	return c.JSON(http.StatusOK, "Large Area")
}