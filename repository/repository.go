package repository

import (
	"deleteLater/models"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"net/http"
)

type Repository struct {
	DB *gorm.DB
}

func (r *Repository) SetupRoutes(app *fiber.App) {
	api := app.Group("/api")
	api.Post("/create_books", r.CreateBook)
	api.Delete("/delete_books/:id", r.DeleteBook)
	api.Post("/get_book/:id", r.GetBook)
	api.Get("/books", r.GetBooks)
}

func (r *Repository) CreateBook(context *fiber.Ctx) error {
	book := models.Books{}
	err := context.BodyParser(&book)

	if err != nil {
		err := context.Status(http.StatusUnprocessableEntity).JSON(
			&fiber.Map{"message": " Invalid input"})
		if err != nil {
			return err
		}
		return err
	}

	err = r.DB.Create(&book).Error
	if err != nil {
		err := context.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message": "could not create the book."})
		if err != nil {
			return err
		}
		return err
	}

	context.Status(http.StatusCreated).JSON(
		&fiber.Map{"message": "book created."})
	if err != nil {
		return err
	}
	return nil

}

func (r *Repository) DeleteBook(context *fiber.Ctx) error {
	return nil
}

func (r *Repository) GetBook(context *fiber.Ctx) error {
	return nil
}

func (r *Repository) GetBooks(context *fiber.Ctx) error {
	bookModels := &[]models.Books{}

	err := r.DB.Find(bookModels).Error
	if err != nil {
		context.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message": "book not found."})
		return err
	}
	context.Status(http.StatusOK).JSON(
		&fiber.Map{"message": "books fetched successfully!",
			"data": bookModels})

	return nil
}
