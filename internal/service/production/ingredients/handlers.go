package ingredients

import (
	"context"
	"github.com/hell-kitchen/api-gateway/internal/convertor"
	"github.com/hell-kitchen/api-gateway/internal/model"
	pb "github.com/hell-kitchen/api-gateway/pkg/api/proto/ingredients"
)

func (s Service) GetAll(ctx context.Context, request model.IngredientsGetAllRequest) (*model.IngredientsGetManyResponse, error) {
	resp, err := s.cli.GetAll(ctx, &pb.GetAllRequest{Name: request.Name})
	if err != nil {
		return nil, err
	}

	result := convertor.Ingredient().AllResponseToGetManyResponse(resp)
	return &result, nil
}

func (s Service) GetByID(ctx context.Context, request model.IngredientsGetByIDRequest) (*model.IngredientsGetByIDResponse, error) {
	resp, err := s.cli.Get(ctx, &pb.GetRequest{Id: request.ID})
	if err != nil {
		return nil, err
	}

	res := convertor.Ingredient().ProtoToDTO(resp.GetIngredient())
	return &res, nil
}
