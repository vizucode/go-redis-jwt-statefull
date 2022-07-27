package main

import (
	"go_jwt_statefull/config"
)

func main() {
	db := config.NewDB()

	config.AutoMigrate(db)
}
