package handlers

import (
	"go_jwt_statefull/exception"
	"go_jwt_statefull/models"
	services "go_jwt_statefull/services/interface"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Users struct {
	Services services.User
}

func NewUsers(services services.User) *Users {
	return &Users{
		Services: services,
	}
}

func (h *Users) FindAll(ctx *gin.Context) {
	userResults := h.Services.FindAll(ctx)
	ctx.JSON(http.StatusOK, gin.H{
		"message": "message oke",
		"data":    userResults,
	})
}

func (h *Users) Store(ctx *gin.Context) {
	usersRequest := new(models.UsersRequestStoreModel)
	err := ctx.ShouldBindJSON(usersRequest)
	if err != nil {
		panic(exception.BadRequestError(err.Error()))
	}
	userResult := h.Services.Store(ctx, usersRequest)
	ctx.JSON(http.StatusOK, gin.H{
		"message": "message oke",
		"data":    userResult,
	})
}
