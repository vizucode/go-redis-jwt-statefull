package middlewares

import (
	"fmt"
	"go_jwt_statefull/config"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

func Authorize(ctx *gin.Context) {
	tokenString, err := ctx.Cookie("token")
	rdb := config.RedisClient()
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": err.Error(),
		})
		return
	}
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return config.JWT_TOP_SECRET, nil
	})
	result := rdb.Get(ctx, tokenString)
	if result.Err() != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": "status unathorized",
		})
		return
	}
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": "status unathorized",
		})
		return
	}
	if !token.Valid {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": "status unathorized",
		})
		return
	}
}
