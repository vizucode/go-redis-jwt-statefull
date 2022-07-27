package handlers

import (
	"go_jwt_statefull/exception"
	"go_jwt_statefull/models"
	services "go_jwt_statefull/services/interface"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Authentication struct {
	Services services.Authentication
}

func NewAuth(services services.Authentication) *Authentication {
	return &Authentication{
		Services: services,
	}
}

func (h *Authentication) Login(ctx *gin.Context) {
	authModel := new(models.AuthModel)
	err := ctx.ShouldBindJSON(authModel)
	if err != nil {
		panic(exception.BadRequestError(err.Error()))
	}
	userResult := h.Services.Login(ctx, authModel)
	ctx.JSON(http.StatusOK, gin.H{
		"message": "message oke",
		"data":    userResult,
	})
}
