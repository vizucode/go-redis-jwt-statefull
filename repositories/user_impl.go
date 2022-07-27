package repositories

import (
	"go_jwt_statefull/entities"

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

func (r *UserImpl) FindAll(ctx *gin.Context) []entities.Users {
	users := []entities.Users{}
	if err := r.DB.Find(&users).Error; err != nil {
		panic(err.Error())
	}
	return users
}

func (r *UserImpl) Store(ctx *gin.Context, user *entities.Users) entities.Users {
	if err := r.DB.Create(&user).Error; err != nil {
		panic(err.Error())
	}
	return *user
}
