package health

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func RegisterRoutes(app *fiber.App, db *gorm.DB) {
	service := NewService(db)
	controller := NewController(service)

	app.Get("/health", controller.Health)
}
