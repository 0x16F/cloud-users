package fflags

import (
	"context"

	"github.com/0x16F/cloud-users/internal/entity"
	of "github.com/open-feature/go-sdk/openfeature"
)

type Service struct {
	client *of.Client
}

func New(client *of.Client) *Service {
	return &Service{
		client: client,
	}
}

func (s *Service) IsFeatureEnabled(ctx context.Context, flag string, user entity.UserData) bool {
	return s.client.Boolean(ctx, flag, false, of.NewEvaluationContext(user.Login, map[string]interface{}{
		"login": user.Login,
		"role":  user.Role,
	}))
}
