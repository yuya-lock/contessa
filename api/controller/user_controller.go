package controller

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"github.com/yuya-lock/contessa/api/database"
	"github.com/yuya-lock/contessa/api/models"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"strconv"
	"time"
)

func Signup(c echo.Context) error {
	logrus.Info("Register User")

	u := new(models.User)
	if err := c.Bind(u); err != nil {
		return nil
	}

	db, sqlDB, err := database.Connect()
	if err != nil {
		return c.String(http.StatusServiceUnavailable, "")
	}
	defer sqlDB.Close()
	database.InitDB()

	hashedPassword, _ := bcrypt.GenerateFromPassword(u.Password, 12)
	user := models.User{}
	db.Where("name = ?", u.Name).First(&user)
	if user.ID != 0 {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": "This account already exists.",
		})
	}
	user = models.User{
		Name:     u.Name,
		Password: hashedPassword,
		Job:      u.Job,
		Age:      u.Age,
		Sex:      u.Sex,
	}
	db.Create(&user)

	return c.JSON(http.StatusOK, user)
}

func Login(c echo.Context) error {
	u := new(models.User)
	if err := c.Bind(u); err != nil {
		return err
	}

	db, sqlDB, err := database.Connect()
	if err != nil {
		return c.String(http.StatusServiceUnavailable, "")
	}
	defer sqlDB.Close()

	user := models.User{}
	db.Where("name = ?", u.Name).First(&user)
	if user.ID == 0 {
		return c.JSON(http.StatusNotFound, echo.Map{
			"message": "User not found",
		})
	}

	if err := bcrypt.CompareHashAndPassword(user.Password, u.Password); err != nil {
		return c.JSON(http.StatusNotFound, echo.Map{
			"message": "Incorrect password",
		})
	}

	claims := jwt.StandardClaims{
		Issuer:    strconv.Itoa(int(user.ID)),
		ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
	}

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := jwtToken.SignedString([]byte("secret"))
	if err != nil {
		return err
	}

	cookie := new(http.Cookie)
	cookie.Name = "jwtToken"
	cookie.Value = token
	cookie.Expires = time.Now().Add(24 * time.Hour)
	c.SetCookie(cookie)

	return c.JSON(http.StatusOK, echo.Map{
		"token": token,
	})
}

func Mypage(c echo.Context) error {
	cookie, err := c.Cookie("jwtToken")
	if err != nil {
		return echo.ErrUnauthorized
	}
	token, err := jwt.ParseWithClaims(cookie.Value, &models.JwtCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte("secret"), nil
	})
	if err != nil || !token.Valid {
		return echo.ErrUnauthorized
	}

	claims := token.Claims.(*models.JwtCustomClaims)
	id := claims.Issuer

	db, sqlDB, err := database.Connect()
	if err != nil {
		return c.String(http.StatusServiceUnavailable, "")
	}
	defer sqlDB.Close()

	var user models.User
	db.Where("id = ?", id).First(&user)

	return c.JSON(http.StatusOK, user)
}

func Logout(c echo.Context) error {
	cookie := new(http.Cookie)
	cookie.Name = "jwtToken"
	cookie.Value = ""
	cookie.Expires = time.Now().Add(-time.Hour)
	c.SetCookie(cookie)

	return c.JSON(http.StatusOK, echo.Map{
		"message": "success",
	})
}
