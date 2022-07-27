package config

import (
	"go_jwt_statefull/entities"

	"gorm.io/gorm"
)

func AutoMigrate(db *gorm.DB) {
	db.AutoMigrate(new(entities.Users))
}
