package middlewares

import (
	"fmt"
	"go_jwt_statefull/config"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

func Authorization(ctx *gin.Context) {
	tokenString, err := ctx.Cookie("token")
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": err.Error(),
			"status":  "unauthorized",
		})
		return
	}
	rdb := config.NewRedisClient()
	defer rdb.Close()
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}
		return config.JWT_SECRET, nil
	})
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": err.Error(),
			"status":  "unauthorized",
		})
		return
	}
	if !token.Valid {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": err.Error(),
			"status":  "unauthorized",
		})
		return
	}
	result := rdb.Get(ctx, tokenString)
	if result.Err() != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": result.Err(),
			"status":  "unauthorized",
		})
		return
	}
}
