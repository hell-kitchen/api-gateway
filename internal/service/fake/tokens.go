package fake

import (
	"context"
	"github.com/hell-kitchen/api-gateway/internal/model"
	"github.com/hell-kitchen/api-gateway/internal/service"
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

func (t *tokensService) Login(ctx context.Context, request model.TokensLoginRequest) (*model.TokensLoginResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (t *tokensService) Logout(ctx context.Context) error {
	//TODO implement me
	panic("implement me")
}
