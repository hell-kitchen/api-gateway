package ingredients

import (
	"context"
	"errors"

	"github.com/hell-kitchen/api-gateway/internal/convertor"
	"github.com/hell-kitchen/api-gateway/internal/model"
	pb "github.com/hell-kitchen/api-gateway/pkg/api/proto/ingredients"
	"go.uber.org/zap"
)

// GetAll return all ingredients by
func (s *Service) GetAll(ctx context.Context, request model.IngredientsGetAllRequest) (*model.IngredientsGetManyResponse, error) {
	resp, err := s.cli.GetAll(ctx, &pb.GetAllRequest{Name: request.Name})
	if err != nil {
		s.log.Error("got error while getting all ingredients from service", zap.Error(err))
		return nil, errors.New("Страница не найдена")
	}

	result := convertor.Ingredient().AllResponseToGetManyResponse(resp)
	return &result, nil
}

func (s *Service) GetByID(ctx context.Context, request model.IngredientsGetByIDRequest) (*model.IngredientsGetByIDResponse, error) {
	resp, err := s.cli.Get(ctx, &pb.GetRequest{Id: request.ID})
	if err != nil {
		return nil, err
	}

	res := convertor.Ingredient().ProtoToDTO(resp.GetIngredient())
	return &res, nil
}
