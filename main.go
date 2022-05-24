package main

import (
	"deleteLater/models"
	rep "deleteLater/repository"
	"deleteLater/storage"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func main() {
	//	Load env
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading the env file")
	}

	// database config
	config := &storage.Config{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		Password: os.Getenv("DB_PASS"),
		User:     os.Getenv("DB_USER"),
		DBName:   os.Getenv("DB_NAME"),
		SSLMode:  os.Getenv("DB_SSL_MODE"),
	}

	db, err := storage.NewConnection(config)

	if err != nil {
		log.Fatal("could not load the DB")
	}

	// Make Migrations
	err = models.RunMigrations(db)

	if err != nil {
		log.Fatal("Error Creating migrations")
	}

	app := fiber.New()

	r := rep.Repository{
		DB: db,
	}

	r.SetupRoutes(app)

	app.Listen(":8000")
}
