package common

import (
	"github.com/RizkiMufrizal/gin-clean-architecture/exception"
	"github.com/golang-jwt/jwt/v4"
	"time"
)

func GenerateToken(username string, roles []map[string]interface{}, jwtSecret string, jwtExpired int) string {
	claims := jwt.MapClaims{
		"username": username,
		"roles":    roles,
		"exp":      time.Now().Add(time.Minute * time.Duration(jwtExpired)).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenSigned, err := token.SignedString([]byte(jwtSecret))
	exception.PanicLogging(err)

	return tokenSigned
}
