package controller

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"github.com/yuya-lock/contessa/api/database"
	"github.com/yuya-lock/contessa/api/models"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"strings"
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

	return c.JSON(http.StatusOK, echo.Map{
		"message": "Your account was created successfully",
	})
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

	claims := &models.JwtCustomClaims{
		UID:  user.ID,
		Name: user.Name,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, echo.Map{
		"token": t,
	})
}

func GetUser(c echo.Context) error {
	auth := c.Request().Header.Get("Authorization")
	if auth == "" {
		return c.JSON(http.StatusUnauthorized, echo.Map{
			"Message": "Authorization Header Not Found",
		})
	}
	splitToken := strings.Split(auth, "Bearer ")
	auth = splitToken[1]
	u, err := jwt.ParseWithClaims(auth, &models.JwtCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte("secret"), nil
	})
	if err != nil {
		return c.JSON(http.StatusNotFound, echo.Map{
			"Message": "Token is wrong or Expire",
		})
	}
	claims := u.Claims.(*models.JwtCustomClaims)

	db, sqlDB, err := database.Connect()
	if err != nil {
		return c.String(http.StatusServiceUnavailable, "")
	}
	defer sqlDB.Close()

	user := models.User{}
	db.Where("name = ?", claims.Name).Preload("Comments").Preload("Likes").Preload("Rates").First(&user)
	if user.ID == 0 {
		return c.JSON(http.StatusNotFound, echo.Map{
			"message": "User not found",
		})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"user": user,
	})
}
