package controller

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"github.com/yuya-lock/contessa/api/database"
	"github.com/yuya-lock/contessa/api/models"
	"golang.org/x/crypto/bcrypt"
	"net/http"
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
	db.Where("name = ?", u.Name).Preload("Comments").Preload("Likes").Preload("Rates").First(&user)
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

	//return c.JSON(http.StatusOK, user)
}

func Restricted(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*models.JwtCustomClaims)
	name := claims.Name
	return c.JSON(http.StatusOK, echo.Map{
		"name": name,
	})
}
