package user

import (
	"github.com/gofiber/fiber/v2"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

func RegisterRoutes(app *fiber.App, db *gorm.DB, redis *redis.Client) {
	service := NewService(db, redis)
	controller := NewController(service)

	userGroup := app.Group("/api/users")

	userGroup.Post("/", controller.CreateUser)
	userGroup.Get("/", controller.ListUsers)
	userGroup.Get("/:id", controller.GetUser)
	userGroup.Put("/:id", controller.UpdateUser)
	userGroup.Delete("/:id", controller.DeleteUser)
	userGroup.Get("/:id/verify", controller.VerifyUser)
}
