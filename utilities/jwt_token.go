package utilities

import (
	"fmt"
	"go-todo-api/config"
	"go-todo-api/models"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func GenerateToken(user models.User) (string, error) {
	conf := config.New()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(time.Minute * 60 * 24 * 7).Unix(),
		"iat": time.Now().Unix(),
	})
	return token.SignedString([]byte(conf.Secret.JWT))
}

func UserIdFromHeader(c *gin.Context) {
	conf := config.New()
	userId := 0
	authorization := c.Request.Header["Authorization"]
	if len(authorization) == 0 {
		return userId
	}
	tokenString := strings.Split(authorization[0], " ")[1]
	if tokenString != "" {
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return []byte(conf.Secret.JWT), nil
		})
		if err != nil {
			fmt.Println(err)
			return userId
		}
		if !token.Valid {
			fmt.Println("invalid")
			return userId
		}
		claims, ok := token.Claims.(jwt.MapClaims)
		if ok {
			sub := claims["sub"].(float64)
			userId = int(sub)
		}
	}
	return userId
}
