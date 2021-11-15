package middlewares

import (
	"kelasd/constant"
	"time"

	"github.com/golang-jwt/jwt"
)

type JwtCustomClaims struct {
	UserId int `json:"userId"`
	jwt.StandardClaims
}

func GenereteTokenJWT(userId int) (string, error) {

	// Set custom claims
	claims := &JwtCustomClaims{
		userId,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
		},
	}

	// Create token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte(constant.SECRET_JWT))

	return t, err
}
