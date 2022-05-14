package rest

import (
	"encoding/json"
	"fmt"
	"github.com/yuya-lock/contessa/api/conf"
	"github.com/yuya-lock/contessa/api/models"
	"io/ioutil"
	"net/http"
	"time"
)

func FetchCocktails(input *models.CocktailsInput) (*models.CocktailsOutput, error) {
	timeout := time.Duration(5 * time.Second)
	client := &http.Client{
		Timeout: timeout,
	}

	request, err := http.NewRequest("GET", conf.COCKTAILS_API_URL, nil)
	if err != nil {
		return nil, err
	}

	params := request.URL.Query()
	params.Add("word", input.Word)
	params.Add("base", input.Base)
	params.Add("technique", input.Technique)
	params.Add("taste", input.Taste)
	params.Add("style", input.Style)
	params.Add("alcohol_from", input.AlcoholFrom)
	params.Add("alcohol_to", input.AlcoholTo)
	params.Add("top", input.Top)
	params.Add("page", input.Page)
	params.Add("limit", input.Limit)
	request.URL.RawQuery = params.Encode()

	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	if response.StatusCode >= 400 {
		return nil, fmt.Errorf("Bad response status code %d.\n%s", response.StatusCode, body)
	}

	responseBody := new(models.CocktailsOutput)
	json.Unmarshal([]byte(body), responseBody)

	return responseBody, nil
}

func FetchCocktailData(page string) (*models.CocktailsOutput, error) {
	timeout := time.Duration(5 * time.Second)
	client := &http.Client{
		Timeout: timeout,
	}

	request, err := http.NewRequest("GET", conf.COCKTAILS_API_URL, nil)
	if err != nil {
		return nil, err
	}

	params := request.URL.Query()
	params.Add("page", page)
	params.Add("limit", "100")
	request.URL.RawQuery = params.Encode()

	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	if response.StatusCode >= 400 {
		return nil, fmt.Errorf("Bad response status code %d.\n%s", response.StatusCode, body)
	}

	responseBody := new(models.CocktailsOutput)
	json.Unmarshal([]byte(body), responseBody)

	return responseBody, nil
}
