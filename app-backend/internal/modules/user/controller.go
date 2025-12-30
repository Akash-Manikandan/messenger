package user

import (
	"github.com/Akash-Manikandan/app-backend/internal/models"
	"github.com/Akash-Manikandan/app-backend/pkg/response"
	"github.com/gofiber/fiber/v2"
)

type Controller struct {
	service Service
}

func NewController(s Service) *Controller {
	return &Controller{service: s}
}

func (c *Controller) CreateUser(ctx *fiber.Ctx) error {
	var req CreateUserRequest
	if err := ctx.BodyParser(&req); err != nil {
		return response.BadRequest(ctx, "Invalid request body")
	}

	// Validate request using proto validation
	if err := req.Validate(); err != nil {
		return response.BadRequest(ctx, err.Error())
	}

	user := &models.User{
		Username:  req.Username,
		Email:     req.Email,
		Password:  req.Password,
		FirstName: req.FirstName,
		LastName:  req.LastName,
	}

	if err := c.service.CreateUser(user); err != nil {
		return response.InternalError(ctx, err.Error())
	}

	return response.Created(ctx, user.Redact())
}

func (c *Controller) GetUser(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	if id == "" {
		return response.BadRequest(ctx, "Invalid user ID")
	}

	user, err := c.service.GetUserByID(id)
	if err != nil {
		return response.NotFound(ctx, "User not found")
	}

	return response.Success(ctx, user.Redact())
}

func (c *Controller) UpdateUser(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	if id == "" {
		return response.BadRequest(ctx, "Invalid user ID")
	}

	var user models.User
	if err := ctx.BodyParser(&user); err != nil {
		return response.BadRequest(ctx, "Invalid request body")
	}

	user.ID = id
	if err := c.service.UpdateUser(&user); err != nil {
		return response.InternalError(ctx, err.Error())
	}

	return response.Success(ctx, user.Redact())
}

func (c *Controller) DeleteUser(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	if id == "" {
		return response.BadRequest(ctx, "Invalid user ID")
	}

	if err := c.service.DeleteUser(id); err != nil {
		return response.InternalError(ctx, err.Error())
	}

	return response.NoContent(ctx)
}

func (c *Controller) ListUsers(ctx *fiber.Ctx) error {
	limit := ctx.QueryInt("limit", 10)
	offset := ctx.QueryInt("offset", 0)

	users, err := c.service.ListUsers(limit, offset)
	if err != nil {
		return response.InternalError(ctx, err.Error())
	}

	// Redact sensitive fields from all users
	for i := range users {
		users[i].Redact()
	}

	return response.SuccessWithMeta(ctx, users, fiber.Map{
		"limit":  limit,
		"offset": offset,
	})
}

func (c *Controller) VerifyUser(ctx *fiber.Ctx) error {
	userID := ctx.Params("id")
	if userID == "" {
		return response.BadRequest(ctx, "Invalid user ID")
	}

	if err := c.service.VerifyUser(userID); err != nil {
		if err.Error() == ErrUserNotFound {
			return response.NotFound(ctx, ErrUserNotFound)
		}
		if err.Error() == ErrUserAlreadyVerified {
			return response.BadRequest(ctx, ErrUserAlreadyVerified)
		}
		return response.InternalError(ctx, err.Error())
	}

	// Get updated user
	user, err := c.service.GetUserByID(userID)
	if err != nil {
		return response.InternalError(ctx, "Failed to fetch user")
	}

	return response.Success(ctx, fiber.Map{
		"success": true,
		"message": "User verified successfully",
		"user":    user.Redact(),
	})
}
