package features

import (
	"context"

	"github.com/0x16F/cloud-common/pkg/logger"
	"github.com/0x16F/cloud-users/internal/controller/httpsrv/extractor"
	"github.com/0x16F/cloud-users/internal/entity"
	"github.com/0x16F/cloud-users/pkg/codes"
	"github.com/gofiber/fiber/v2"
)

type FFlagsService interface {
	IsFeatureEnabled(ctx context.Context, flag string, user entity.UserData) bool
}

type ErrorsService interface {
	GetError(code int) error
}

type Service struct {
	fflagsService FFlagsService
	errorsService ErrorsService
}

func New(fflagsService FFlagsService, errorsService ErrorsService) *Service {
	return &Service{
		fflagsService: fflagsService,
		errorsService: errorsService,
	}
}

func (h *Service) IsFeatureEnabled(c *fiber.Ctx, log logger.Logger, handlerName string) error {
	userData := extractor.Extract(c)

	if !h.fflagsService.IsFeatureEnabled(c.Context(), handlerName, userData) {
		log.Warnf("feature is disabled for user %s with role %s", userData.Login, userData.Role)

		return h.errorsService.GetError(codes.FeatureIsDisabled)
	}

	return nil
}
