package repositories

import (
	"go_jwt_statefull/entities"
	"go_jwt_statefull/exception"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type UserImpl struct {
	DB *gorm.DB
}

func NewUser(db *gorm.DB) *UserImpl {
	return &UserImpl{
		DB: db,
	}
}

func (r *UserImpl) FindCredential(ctx *gin.Context, email string) entities.Users {
	users := entities.Users{}
	if err := r.DB.WithContext(ctx).Where("email = ?", email).First(&users).Error; err != nil {
		panic(exception.NotFoundError(err.Error()))
	}
	return users
}
