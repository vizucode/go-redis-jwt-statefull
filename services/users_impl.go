package services

import (
	"go_jwt_statefull/config"
	"go_jwt_statefull/exception"
	"go_jwt_statefull/models"
	repositories "go_jwt_statefull/repositories/interface"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

type AuthenticationImpl struct {
	Repo repositories.User
}

func NewAuthentication(repo repositories.User) *AuthenticationImpl {
	return &AuthenticationImpl{
		Repo: repo,
	}
}

func (s *AuthenticationImpl) Login(ctx *gin.Context, user *models.AuthModel) string {
	redisClient := config.NewRedisClient()
	userResult := s.Repo.FindCredential(ctx, user.Email)
	err := bcrypt.CompareHashAndPassword([]byte(userResult.Password), []byte(user.Password))
	if err != nil {
		panic(exception.NotFoundError(err.Error()))
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, config.JWTClaims{
		Name:  userResult.Name,
		Email: userResult.Email,
		RegisteredClaims: jwt.RegisteredClaims{
			ID:        strconv.Itoa(userResult.Id),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(config.JWT_DURATION) * time.Minute)),
		},
	})
	tokenString, err := token.SignedString(config.JWT_SECRET)
	if err != nil {
		panic(exception.InternalServerError(err.Error()))
	}
	redisError := redisClient.Set(ctx, tokenString, tokenString, (time.Duration(config.JWT_DURATION) * time.Minute))
	if redisError.Err() != nil {
		panic(exception.InternalServerError(redisError.Err().Error()))
	}
	defer redisClient.Close()
	return tokenString
}
