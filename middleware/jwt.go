package middleware

import (
	"fmt"
	"github.com/RizkiMufrizal/gin-clean-architecture/common"
	"github.com/RizkiMufrizal/gin-clean-architecture/configuration"
	"github.com/RizkiMufrizal/gin-clean-architecture/model"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"strings"
)

func AuthenticateJWT(role string, config configuration.Config) gin.HandlerFunc {
	jwtSecret := config.Get("JWT_SECRET_KEY")
	return func(c *gin.Context) {
		authorizationHeader := c.GetHeader("Authorization")
		if !strings.Contains(authorizationHeader, "Bearer") {
			c.JSON(400, model.GeneralResponse{
				Code:    400,
				Message: "Bad Request",
				Data:    "Missing or malformed JWT",
			})
			return
		}

		tokenString := strings.Replace(authorizationHeader, "Bearer ", "", -1)
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}

			return jwtSecret, nil
		})
		if err != nil {
			c.JSON(400, model.GeneralResponse{
				Code:    400,
				Message: "Bad Request",
				Data:    "Missing or malformed JWT",
			})
			return
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			roles := claims["roles"].([]interface{})
			common.NewLogger().Info("role function ", role, " role user ", roles)
			for _, roleInterface := range roles {
				roleMap := roleInterface.(map[string]interface{})
				if roleMap["role"] == role {
					c.Next()
					return
				}
			}

			c.JSON(401, model.GeneralResponse{
				Code:    401,
				Message: "Unauthorized",
				Data:    "Invalid Role",
			})
			return
		} else {
			c.JSON(401, model.GeneralResponse{
				Code:    401,
				Message: "Unauthorized",
				Data:    "Invalid or expired JWT",
			})
			return
		}
	}
}
