package repositories

import (
	"go_jwt_statefull/entities"

	"github.com/gin-gonic/gin"
)

type User interface {
	FindCredential(ctx *gin.Context, email string) entities.Users
}
