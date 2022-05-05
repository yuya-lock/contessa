package main

import (
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"github.com/yuya-lock/contessa/backend/lib"
	"golang.org/x/crypto/bcrypt"
	"io/ioutil"
	"net/http"
	"time"
)

func index(c echo.Context) error {
	return c.String(http.StatusOK, "OK")
}

func fetchCocktails(c echo.Context) error {
	logrus.Info("Get cocktails")

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
	if err != nil {
		logrus.Fatal(err)
	}

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

//var signingKey = []byte("secret")

//var Config = middleware.JWTConfig{
//	Claims:     &jwtCustomClaims{},
//	SigningKey: signingKey,
//}

func Register(c echo.Context) error {
	logrus.Info("Regist User")

	u := new(UserInput)
	if err := c.Bind(u); err != nil {
		return nil
	}

	db, sqlDB, err := lib.Connect()
	if err != nil {
		return c.String(http.StatusServiceUnavailable, "")
	}
	defer sqlDB.Close()
	lib.InitDB()

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(u.Password), 12)
	user := lib.User{
		Name:     u.Name,
		Password: hashedPassword,
		Job:      u.Job,
		Age:      u.Age,
		Sex:      u.Sex,
	}
	db.Create(&user)

	return c.JSON(http.StatusOK, user)
}

//func Login(c echo.Context) error {
//	u := new(UserInput)
//	if err := c.Bind(u); err != nil {
//		return err
//	}
//
//	user := lib.User{}
//	user.FirstByName(u.Name)
//	if user.ID == 0 {
//		return c.JSON(http.StatusNotFound, echo.Map{
//			"message": "User not found",
//		})
//	}
//
//	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(u.Password))
//	if err != nil {
//		return echo.ErrUnauthorized
//	}
//
//	claims := &jwtCustomClaims{
//		user.ID,
//		user.Name,
//		jwt.StandardClaims{
//			Issuer:    strconv.Itoa(int(user.ID)),
//			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
//		},
//	}
//
//	jwtToken := jwt.NewWithClaims(jwt.SigningMethodES256, claims)
//
//	token, err := jwtToken.SignedString(signingKey)
//	if err != nil {
//		return err
//	}
//
//	cookie := new(http.Cookie)
//	cookie.Name = "jwt"
//	cookie.Value = token
//	cookie.Expires = time.Now().Add(24 * time.Hour)
//	c.SetCookie(cookie)
//
//	return c.JSON(http.StatusOK, echo.Map{
//		"jwt": token,
//	})
//}
//
//type Claims struct {
//	jwt.StandardClaims
//}
//
//func Logout(c echo.Context) error {
//	cookie := new(http.Cookie)
//	cookie.Name = "jwt"
//	cookie.Value = ""
//	cookie.Expires = time.Now().Add(-time.Hour)
//	c.SetCookie(cookie)
//
//	return c.JSON(http.StatusOK, echo.Map{
//		"message": "success",
//	})
//}
//
//func getUser(c echo.Context) error {
//	cookie, err := c.Cookie("jwt")
//	if err != nil {
//		return err
//	}
//
//	token, err := jwt.ParseWithClaims(cookie.String(), &Claims{}, func(token *jwt.Token) (interface{}, error) {
//		return []byte("secret"), nil
//	})
//	if err != nil || !token.Valid {
//		return echo.ErrUnauthorized
//	}
//
//	claims := token.Claims.(*Claims)
//	i, _ := strconv.Atoi(claims.Issuer)
//	id := uint(i)
//
//	var user lib.User
//	user.FirstById(id)
//
//	return c.JSON(http.StatusOK, user)
//}
//
////func getUser(c echo.Context) error {
////	i, _ := strconv.Atoi(c.Param("id"))
////	id := uint(i)
////	user := lib.User{}
////	user.FirstById(id)
////
////	return c.JSON(http.StatusOK, user)
////}
//
//func createUser(c echo.Context) error {
//	u := new(UserInput)
//	if err := c.Bind(u); err != nil {
//		return err
//	}
//
//	hashed, _ := bcrypt.GenerateFromPassword([]byte(u.Password), 12)
//	hashedPassword := string(hashed)
//
//	user := lib.User{
//		Name:     u.Name,
//		Password: hashedPassword,
//		Age:      u.Age,
//		Job:      u.Job,
//		Sex:      u.Sex,
//	}
//	user.Create()
//
//	return c.JSON(http.StatusOK, user)
//}
//
////func updateUser(c echo.Context) error {
////	i, _ := strconv.Atoi(c.Param("id"))
////	id := uint(i)
////	name := c.FormValue("name")
////	age := c.FormValue("age")
////	job := c.FormValue("job")
////	sex := c.FormValue("sex")
////
////	user := lib.User{
////		ID:   id,
////		Name: name,
////		Age:  age,
////		Job:  job,
////		Sex:  sex,
////	}
////	user.Updates()
////
////	return c.JSON(http.StatusOK, user)
////}
//
//func deleteUser(c echo.Context) error {
//	i, _ := strconv.Atoi(c.Param("id"))
//	id := uint(i)
//	user := lib.User{}
//	user.DeleteById(id)
//
//	return c.JSON(http.StatusOK, user)
//}
//
//func Restricted(c echo.Context) error {
//	user := c.Get("user").(*jwt.Token)
//	claims := user.Claims.(*jwtCustomClaims)
//	name := claims.Name
//	return c.JSON(http.StatusOK, echo.Map{
//		"username": name,
//	})
//}
