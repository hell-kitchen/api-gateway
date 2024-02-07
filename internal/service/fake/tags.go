package fake

import (
	"context"
	"github.com/hell-kitchen/api-gateway/internal/model"
	"github.com/hell-kitchen/api-gateway/internal/service"
)

var (
	_ service.TagsService = (*tagsService)(nil)
)

type tagsService struct {
	service *Service
}

func newTags(srv *Service) {
	srv.tags = &tagsService{
		service: srv,
	}
}

func (t *tagsService) GetAll(ctx context.Context) (*model.TagsGetManyResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (t *tagsService) Get(ctx context.Context, request model.TagsGetByIDRequest) (*model.TagsGetOneResponse, error) {
	//TODO implement me
	panic("implement me")
}
