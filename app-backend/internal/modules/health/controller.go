package health

import "github.com/gofiber/fiber/v2"

type Controller struct {
	service Service
}

func NewController(s Service) *Controller {
	return &Controller{service: s}
}

func (c *Controller) Health(ctx *fiber.Ctx) error {
	return ctx.JSON(fiber.Map{
		"status":             c.service.Status(),
		"postgres_db_status": c.service.PostgresStatus(),
	})
}
