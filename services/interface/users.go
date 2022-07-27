package services

import (
	"go_jwt_statefull/models"

	"github.com/gin-gonic/gin"
)

type User interface {
	FindAll(ctx *gin.Context) []models.UsersResponseModel
}
