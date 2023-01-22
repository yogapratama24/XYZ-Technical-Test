package helpers

import (
	"log"
	"os"
	"strconv"
	"time"
	"xyz_test/config"
	"xyz_test/model"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
)

var SecretKey = []byte(os.Getenv("SECRET_KEY"))

type JwtClaims struct {
	jwt.StandardClaims
}

func GenerateJWT(user model.User) (string, error) {
	// claims := jwt.StandardClaims{
	// 	ExpiresAt: time.Now().Add(30 * time.Hour).Unix(),
	// 	Issuer:    strconv.Itoa(user.Id),
	// }
	claims := JwtClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(30 * time.Hour).Unix(),
			Issuer:    strconv.Itoa(user.Id),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(SecretKey)
	if err != nil {
		log.Printf("Error signed string jwt with err: %s\n", err)
		return tokenString, err
	}
	return tokenString, err
}

func ParseJWT(c echo.Context) (model.User, error) {
	db := config.Connect()
	user := model.User{}
	var issuer interface{}

	token, err := c.Cookie("token")
	if err != nil {
		log.Printf("Error get cookie with err: %s", err)
		return user, err
	}
	claims := jwt.MapClaims{}
	res, err := jwt.ParseWithClaims(token.Value, claims, func(t *jwt.Token) (interface{}, error) {
		return SecretKey, nil
	})
	if err != nil {
		log.Printf("Error parse jwt token with err: %s", err)
		return user, err
	}
	resMap, ok := res.Claims.(jwt.MapClaims)
	if ok && res.Valid {
		issuer = resMap["iss"]
		issuers, err := strconv.Atoi(issuer.(string))
		if err != nil {
			log.Printf("Error conver issuer string to int with : %v", err)
			return user, err
		}
		db.Find(&user, issuers)
		return user, err
	} else {
		log.Printf("Invalid token jwt")
		return user, err
	}
}
