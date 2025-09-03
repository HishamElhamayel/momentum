package pkg

import (
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var jwtKey = []byte(os.Getenv("JWT_SECRET"))

type Claims struct{
	UserID uint
	jwt.RegisteredClaims
}

func GenerateToken(userId uint) (string,error){
	expiresAt := time.Now().Add(time.Hour * 24)

	claims := Claims{
		UserID: userId,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expiresAt),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
}

func VerifyToken(tokenString string) (*Claims,error){
	claims := &Claims{}
	token,err := jwt.ParseWithClaims(tokenString,claims, func(token *jwt.Token)(any, error){
		return jwtKey,nil
	})

	if err != nil{
		return nil, err
	}

	if !token.Valid{
		return nil, errors.New("invalid token")
	}

	return claims, nil

}