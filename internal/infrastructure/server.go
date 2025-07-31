package infrastructure

import (
	"gensmart/internal/delivery"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func InitializeServer(db *gorm.DB) *fiber.App {
	app := fiber.New(fiber.Config{})
	delivery.Router(app, db)
	return app
}
