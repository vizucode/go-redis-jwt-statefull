package services

import (
	"go_jwt_statefull/entities"
	"go_jwt_statefull/exception"
	"go_jwt_statefull/models"
	repositories "go_jwt_statefull/repositories/interface"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
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
		userModelResponse = append(userModelResponse, models.UsersResponseModel{
			Id:    user.Id,
			Name:  user.Name,
			Email: user.Email,
		})
	}
	return userModelResponse
}

func (s *UsersImpl) Store(ctx *gin.Context, user *models.UsersRequestStoreModel) models.UsersResponseModel {
	result, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		panic(exception.BadRequestError(err.Error()))
	}
	user.Password = string(result)
	userResult := s.Repo.Store(ctx, &entities.Users{
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password,
	})
	return models.UsersResponseModel{
		Id:    userResult.Id,
		Name:  userResult.Name,
		Email: userResult.Email,
	}
}
