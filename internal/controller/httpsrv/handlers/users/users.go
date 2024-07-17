package users

import (
	"context"
	"strconv"

	"github.com/0x16F/cloud-common/pkg/logger"
	"github.com/0x16F/cloud-users/internal/entity"
	"github.com/0x16F/cloud-users/pkg/codes"
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

type FeaturesService interface {
	IsFeatureEnabled(c *fiber.Ctx, log logger.Logger, handlerName string) error
}

type Handler struct {
	log             logger.Logger
	usersService    UsersService
	errorsService   ErrorsService
	featuresService FeaturesService
}

func NewHandler(
	log logger.Logger,
	usersService UsersService,
	errorsService ErrorsService,
	featuresService FeaturesService,
) *Handler {
	return &Handler{
		log:             log,
		usersService:    usersService,
		errorsService:   errorsService,
		featuresService: featuresService,
	}
}

func (h *Handler) CreateUser(c *fiber.Ctx) error {
	log := h.log.WithFields(logger.Fields{
		"method": "CreateUser",
	})

	if err := h.featuresService.IsFeatureEnabled(c, log, "create_user"); err != nil {
		return err
	}

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

func (h *Handler) GetUser(c *fiber.Ctx) error {
	log := h.log.WithFields(logger.Fields{
		"method": "GetUser",
	})

	if err := h.featuresService.IsFeatureEnabled(c, log, "get_user"); err != nil {
		return err
	}

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

func (h *Handler) GetUsers(c *fiber.Ctx) error {
	log := h.log.WithFields(logger.Fields{
		"method": "GetUsers",
	})

	if err := h.featuresService.IsFeatureEnabled(c, log, "get_users"); err != nil {
		return err
	}

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

func (h *Handler) UpdateEmail(c *fiber.Ctx) error {
	log := h.log.WithFields(logger.Fields{
		"method": "UpdateEmail",
	})

	if err := h.featuresService.IsFeatureEnabled(c, log, "update_email"); err != nil {
		return err
	}

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

func (h *Handler) UpdateUsername(c *fiber.Ctx) error {
	log := h.log.WithFields(logger.Fields{
		"method": "UpdateUsername",
	})

	if err := h.featuresService.IsFeatureEnabled(c, log, "update_username"); err != nil {
		return err
	}

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

func (h *Handler) UpdatePassword(c *fiber.Ctx) error {
	log := h.log.WithFields(logger.Fields{
		"method": "UpdatePassword",
	})

	if err := h.featuresService.IsFeatureEnabled(c, log, "update_password"); err != nil {
		return err
	}

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

func (h *Handler) DeleteUser(c *fiber.Ctx) error {
	log := h.log.WithFields(logger.Fields{
		"method": "DeleteUser",
	})

	if err := h.featuresService.IsFeatureEnabled(c, log, "delete_user"); err != nil {
		return err
	}

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
