package main

import (
	"io/ioutil"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

func index(c echo.Context) error {
	return c.String(http.StatusOK, "OK")
}

func fetchCocktails(c echo.Context) error {
	input := new(GetCocktailsInput)
	if err := c.Bind(input); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err := input.validator(); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	req, err := http.NewRequest("GET", COCKTAILS_API_URL, nil)
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
	if err != nil{
		logrus.Fatal(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logrus.Fatalf("ioutil.ReadAll err=%s", err.Error())
	}

	return c.JSON(http.StatusOK, string(body))
}

func fetchCocktailDetail(c echo.Context) error {
	id := c.Param("id")
	endpoint := COCKTAILS_API_URL + "/" + id
	
	req, err := http.NewRequest("GET", endpoint, nil)
	if err != nil {
		logrus.Fatal(err)
	}

	timeout := time.Duration(5 * time.Second)
	client := &http.Client{
        Timeout: timeout,
	}

	resp, err := client.Do(req)
	if err != nil{
		logrus.Fatal(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logrus.Fatal(err)
	}
	
	return c.JSON(http.StatusOK, string(body))
}
