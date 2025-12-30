package response

import "github.com/gofiber/fiber/v2"

// Error sends a JSON error response
func Error(ctx *fiber.Ctx, status int, message string) error {
	return ctx.Status(status).JSON(fiber.Map{
		"error": message,
	})
}

// Success sends a JSON success response with data
func Success(ctx *fiber.Ctx, data any) error {
	return ctx.JSON(data)
}

// Created sends a 201 Created response with data
func Created(ctx *fiber.Ctx, data any) error {
	return ctx.Status(fiber.StatusCreated).JSON(data)
}

// NoContent sends a 204 No Content response
func NoContent(ctx *fiber.Ctx) error {
	return ctx.Status(fiber.StatusNoContent).Send(nil)
}

// BadRequest sends a 400 Bad Request error
func BadRequest(ctx *fiber.Ctx, message string) error {
	return Error(ctx, fiber.StatusBadRequest, message)
}

// NotFound sends a 404 Not Found error
func NotFound(ctx *fiber.Ctx, message string) error {
	return Error(ctx, fiber.StatusNotFound, message)
}

// Conflict sends a 409 Conflict error
func Conflict(ctx *fiber.Ctx, message string) error {
	return Error(ctx, fiber.StatusConflict, message)
}

// InternalError sends a 500 Internal Server Error
func InternalError(ctx *fiber.Ctx, message string) error {
	return Error(ctx, fiber.StatusInternalServerError, message)
}

// SuccessWithMeta sends a JSON response with data and metadata
func SuccessWithMeta(ctx *fiber.Ctx, data any, meta fiber.Map) error {
	return ctx.JSON(fiber.Map{
		"data": data,
		"meta": meta,
	})
}
