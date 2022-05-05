package controller

import (
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"github.com/yuya-lock/contessa/api/conf"
	"github.com/yuya-lock/contessa/api/models"
	"io/ioutil"
	"net/http"
	"time"
)

func Index(c echo.Context) error {
	return c.String(http.StatusOK, "OK")
}

func FetchCocktails(c echo.Context) error {
	logrus.Info("Get cocktails")

	input := new(models.CocktailsInput)
	if err := c.Bind(input); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	req, err := http.NewRequest("GET", conf.COCKTAILS_API_URL, nil)
	if err != nil {
		logrus.Fatal(err)
	}

	q := req.URL.Query()
	q.Add("word", input.Word)
	q.Add("base", input.Base)
	q.Add("technique", input.Technique)
	q.Add("taste", input.Taste)
	q.Add("style", input.Style)
	q.Add("alcohol_from", input.AlcoholFrom)
	q.Add("alcohol_to", input.AlcoholTo)
	q.Add("top", input.Top)
	q.Add("page", input.Page)
	q.Add("limit", input.Limit)
	req.URL.RawQuery = q.Encode()

	timeout := time.Duration(5 * time.Second)
	client := &http.Client{
		Timeout: timeout,
	}

	resp, err := client.Do(req)
	if err != nil {
		logrus.Fatal(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logrus.Fatalf("ioutil.ReadAll err=%s", err.Error())
	}

	return c.JSON(http.StatusOK, string(body))
}

func FetchCocktailDetail(c echo.Context) error {
	id := c.Param("id")
	endpoint := conf.COCKTAILS_API_URL + "/" + id

	req, err := http.NewRequest("GET", endpoint, nil)
	if err != nil {
		logrus.Fatal(err)
	}

	timeout := time.Duration(5 * time.Second)
	client := &http.Client{
		Timeout: timeout,
	}

	resp, err := client.Do(req)
	if err != nil {
		logrus.Fatal(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logrus.Fatal(err)
	}

	return c.JSON(http.StatusOK, string(body))
}
