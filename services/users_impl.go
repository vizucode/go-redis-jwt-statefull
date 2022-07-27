package services

import (
	"go_jwt_statefull/models"
	repositories "go_jwt_statefull/repositories/interface"

	"github.com/gin-gonic/gin"
)

type UsersImpl struct {
	Repo repositories.User
}

func NewUserImpl(repo repositories.User) *UsersImpl {
	return &UsersImpl{
		Repo: repo,
	}
}

func (s *UsersImpl) FindAll(ctx *gin.Context) []models.UsersResponseModel {
	userModelResponse := []models.UsersResponseModel{}
	userResults := s.Repo.FindAll(ctx)
	for _, user := range userResults {
		userModelResponse = append(userModelResponse, (models.UsersResponseModel)(user))
	}
	return userModelResponse
}
