package users

import (
	"context"
	"strconv"

	"github.com/0x16F/cloud/users/internal/entity"
	"github.com/0x16F/cloud/users/pkg/codes"
	"github.com/0x16F/cloud/users/pkg/logger"
	"github.com/gofiber/fiber/v2"
)

type UsersService interface {
	CreateUser(ctx context.Context, dto entity.UserCreateDTO) (entity.User, error)
	GetUser(ctx context.Context, id uint64) (entity.User, error)
	GetUsers(ctx context.Context, params entity.GetUsersParams) ([]entity.User, error)
	UpdateEmail(ctx context.Context, id uint64, email string) error
	UpdateUsername(ctx context.Context, id uint64, username string) error
	UpdatePassword(ctx context.Context, id uint64, oldPassword, newPassword string) error
	DeleteUser(ctx context.Context, id uint64) error
}

type ErrorsService interface {
	GetError(code int) error
}

type Handler struct {
	log           logger.Logger
	usersService  UsersService
	errorsService ErrorsService
}

func NewHandler(log logger.Logger, usersService UsersService, errorsService ErrorsService) *Handler {
	return &Handler{
		log:           log,
		usersService:  usersService,
		errorsService: errorsService,
	}
}

func (h *Handler) createUser(c *fiber.Ctx) error {
	log := h.log.WithFields(logger.Fields{
		"method": "CreateUser",
	})

	var req entity.UserCreateDTO

	if err := c.BodyParser(&req); err != nil {
		log.Errorf("failed to parse request body: %v", err)

		return h.errorsService.GetError(codes.InvalidBody)
	}

	user, err := h.usersService.CreateUser(c.Context(), req)
	if err != nil {
		log.Errorf("failed to create user: %v", err)

		return err
	}

	return c.JSON(user)
}

func (h *Handler) getUser(c *fiber.Ctx) error {
	log := h.log.WithFields(logger.Fields{
		"method": "GetUser",
	})

	id, err := strconv.ParseUint(c.Params("id"), 10, 64)
	if err != nil {
		log.Errorf("failed to parse id: %v", err)

		return h.errorsService.GetError(codes.InvalidID)
	}

	user, err := h.usersService.GetUser(c.Context(), id)
	if err != nil {
		log.Errorf("failed to get user: %v", err)

		return err
	}

	return c.JSON(user)
}

func (h *Handler) getUsers(c *fiber.Ctx) error {
	log := h.log.WithFields(logger.Fields{
		"method": "GetUsers",
	})

	var req GetUsersReq

	if err := c.QueryParser(&req); err != nil {
		log.Errorf("failed to parse query params: %v", err)

		return h.errorsService.GetError(codes.InvalidQuery)
	}

	params := entity.GetUsersParams{
		Limit:    req.Limit,
		LastID:   req.LastID,
		Username: req.Username,
		Email:    req.Email,
	}

	users, err := h.usersService.GetUsers(c.Context(), params)
	if err != nil {
		log.Errorf("failed to get users: %v", err)

		return err
	}

	return c.JSON(GetUsersResp{
		Users: users,
	})
}

func (h *Handler) updateEmail(c *fiber.Ctx) error {
	log := h.log.WithFields(logger.Fields{
		"method": "UpdateEmail",
	})

	id, err := strconv.ParseUint(c.Params("id"), 10, 64)
	if err != nil {
		log.Errorf("failed to parse id: %v", err)

		return h.errorsService.GetError(codes.InvalidID)
	}

	var req UpdateEmailReq

	if err = c.BodyParser(&req); err != nil {
		log.Errorf("failed to parse request body: %v", err)

		return h.errorsService.GetError(codes.InvalidBody)
	}

	if err = h.usersService.UpdateEmail(c.Context(), id, req.Email); err != nil {
		log.Errorf("failed to update email: %v", err)

		return err
	}

	return nil
}

func (h *Handler) updateUsername(c *fiber.Ctx) error {
	log := h.log.WithFields(logger.Fields{
		"method": "UpdateUsername",
	})

	id, err := strconv.ParseUint(c.Params("id"), 10, 64)
	if err != nil {
		log.Errorf("failed to parse id: %v", err)

		return h.errorsService.GetError(codes.InvalidID)
	}

	var req UpdateUsernameReq

	if err = c.BodyParser(&req); err != nil {
		log.Errorf("failed to parse request body: %v", err)

		return h.errorsService.GetError(codes.InvalidBody)
	}

	if err = h.usersService.UpdateUsername(c.Context(), id, req.Username); err != nil {
		log.Errorf("failed to update username: %v", err)

		return err
	}

	return nil
}

func (h *Handler) updatePassword(c *fiber.Ctx) error {
	log := h.log.WithFields(logger.Fields{
		"method": "UpdatePassword",
	})

	id, err := strconv.ParseUint(c.Params("id"), 10, 64)
	if err != nil {
		log.Errorf("failed to parse id: %v", err)

		return err
	}

	var req UpdatePasswordReq

	if err := c.BodyParser(&req); err != nil {
		log.Errorf("failed to parse request body: %v", err)

		return err
	}

	if err = h.usersService.UpdatePassword(c.Context(), id, req.OldPassword, req.NewPassword); err != nil {
		log.Errorf("failed to update password: %v", err)

		return err
	}

	return nil
}

func (h *Handler) deleteUser(c *fiber.Ctx) error {
	log := h.log.WithFields(logger.Fields{
		"method": "DeleteUser",
	})

	id, err := strconv.ParseUint(c.Params("id"), 10, 64)
	if err != nil {
		log.Errorf("failed to parse id: %v", err)

		return h.errorsService.GetError(codes.InvalidID)
	}

	if err = h.usersService.DeleteUser(c.Context(), id); err != nil {
		log.Errorf("failed to delete user: %v", err)

		return err
	}

	return nil
}

func (h *Handler) RegisterRoutes(group fiber.Router) {
	group.Get("/", h.getUsers)
	group.Get("/:id", h.getUser)
	group.Post("/", h.createUser)
	group.Patch("/:id/email", h.updateEmail)
	group.Patch("/:id/username", h.updateUsername)
	group.Patch("/:id/password", h.updatePassword)
	group.Delete("/:id", h.deleteUser)
}
