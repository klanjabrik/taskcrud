package main

import (
	"taskcrud/models"
	"taskcrud/routes"

	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		return
	}

	db := models.SetupDB()
	db.AutoMigrate(
		&models.Task{},
		&models.User{})

	r := routes.SetupRoutes(db)
	r.Run()
}
