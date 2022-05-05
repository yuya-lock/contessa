package controller

import (
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"github.com/yuya-lock/contessa/api/database"
	"github.com/yuya-lock/contessa/api/models"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

//var signingKey = []byte("secret")

//var Config = middleware.JWTConfig{
//	Claims:     &jwtCustomClaims{},
//	SigningKey: signingKey,
//}

func Register(c echo.Context) error {
	logrus.Info("Regist User")

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

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(u.Password), 12)
	user := models.User{
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
