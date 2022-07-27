package services

import (
	"go_jwt_statefull/models"

	"github.com/gin-gonic/gin"
)

type User interface {
	FindAll(ctx *gin.Context) []models.UsersResponseModel
	Store(ctx *gin.Context, user *models.UsersRequestStoreModel) models.UsersResponseModel
}
