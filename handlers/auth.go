package handlers

import (
	"go_jwt_statefull/config"
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
	ctx.SetCookie("token", userResult, config.JWT_DURATION*3600, "/", "", false, true)
	ctx.JSON(http.StatusOK, gin.H{
		"message": "message oke",
		"data":    userResult,
	})
}

func (h *Authentication) Logout(ctx *gin.Context) {
	cookies, err := ctx.Cookie("token")
	if err != nil {
		panic(exception.NotFoundError(err.Error()))
	}
	result := h.Services.Logout(ctx, cookies)
	ctx.JSON(http.StatusOK, gin.H{
		"message": "message ok",
		"data":    result,
	})
}
