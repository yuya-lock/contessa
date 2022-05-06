package controller

import (
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"github.com/yuya-lock/contessa/api/models"
	"github.com/yuya-lock/contessa/api/rest"
	"net/http"
)

func Index(c echo.Context) error {
	return c.String(http.StatusOK, "OK")
}

func GetCocktails(c echo.Context) error {
	logrus.Info("Get cocktails")

	input := new(models.CocktailsInput)
	if err := c.Bind(input); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	responseBody, err := rest.FetchCocktails(input)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, responseBody)
}

func GetCocktailDetail(c echo.Context) error {
	id := c.Param("id")

	responseBody, err := rest.FetchCocktailDetail(id)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, responseBody)
}
