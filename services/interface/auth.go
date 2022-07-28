package services

import (
	"go_jwt_statefull/models"

	"github.com/gin-gonic/gin"
)

type Authentication interface {
	Login(ctx *gin.Context, users *models.AuthModel) string
	Logout(ctx *gin.Context, token string) bool
}
