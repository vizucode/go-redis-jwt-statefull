package handlers

import (
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
