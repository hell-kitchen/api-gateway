package fake

import (
	"context"
	"github.com/hell-kitchen/api-gateway/internal/model"
	"github.com/hell-kitchen/api-gateway/internal/service"
	"github.com/labstack/gommon/random"
)

var (
	_ service.TokensService = (*tokensService)(nil)
)

type tokensService struct {
	service *Service
}

func newTokens(srv *Service) {
	srv.tokens = &tokensService{
		service: srv,
	}
}

func (t *tokensService) Login(context.Context, model.TokensLoginRequest) (*model.TokensLoginResponse, error) {
	resp := model.TokensLoginResponse{
		AuthToken: random.String(20, random.Alphanumeric),
	}
	return &resp, nil
}

func (t *tokensService) Logout(context.Context) error {
	return nil
}
