package fake

import (
	"context"
	"github.com/google/uuid"
	"github.com/hell-kitchen/api-gateway/internal/model"
	"github.com/hell-kitchen/api-gateway/internal/service"
	"github.com/labstack/gommon/random"
	"math/rand"
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

func (t *tagsService) GetAll(_ context.Context) (*model.TagsGetManyResponse, error) {
	count := rand.Int() % 20
	var resp = make([]model.TagDTO, 0, count)
	for i := 0; i != count; i++ {
		resp = append(resp, model.TagDTO{
			ID:    uuid.NewString(),
			Name:  random.String(10, random.Alphabetic),
			Color: random.String(6, random.Hex),
			Slug:  uuid.NewString(),
		})
	}
	result := model.TagsGetManyResponse(resp)
	return &result, nil
}

func (t *tagsService) Get(_ context.Context, request model.TagsGetByIDRequest) (*model.TagsGetOneResponse, error) {
	return &model.TagsGetOneResponse{
		ID:    request.ID,
		Name:  random.String(uint8(rand.Int()%20), random.Alphabetic),
		Color: random.String(6, random.Hex),
		Slug:  uuid.NewString(),
	}, nil
}
