package repositories

import (
	"go_jwt_statefull/entities"

	"github.com/gin-gonic/gin"
)

type User interface {
	FindAll(ctx *gin.Context) []entities.Users
}
