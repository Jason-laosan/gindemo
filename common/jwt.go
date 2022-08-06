package common

import (
	"gindemo/model"
	"time"

	"github.com/dgrijalva/jwt-go/v4"
)

type Claims struct {
	UserId uint
	jwt.StandardClaims
}

var jwtKey = []byte("a_secret_crect")

func ReleaseToken(user model.User) (string, error) {
	// expirationTime :=
	expirationTime := jwt.NewTime(float64(time.Now().Add(7 * 24 * time.Hour).Unix()))

	claims := &Claims{
		UserId: user.ID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime,
			IssuedAt:  jwt.Now(),
			Issuer:    "oceanlearn.tech",
			Subject:   "user token",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenstring, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}
	return tokenstring, nil
}

func ParseToken(tokenString string) (*jwt.Token, *Claims, error) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (i interface{}, err error) {
		return jwtKey, nil
	})
	return token, claims, err
}
