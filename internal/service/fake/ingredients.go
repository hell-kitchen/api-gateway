package fake

import (
	"context"
	"github.com/google/uuid"
	"github.com/hell-kitchen/api-gateway/internal/model"
	"github.com/hell-kitchen/api-gateway/internal/service"
	"math/rand"
)

var (
	_ service.IngredientsService = (*ingredientsService)(nil)
)

type ingredientsService struct {
	service service.Interface
}

func NewIngredients(srv service.Interface) {
	srv.ApplyIngredients(&ingredientsService{service: srv})
}

func (srv *ingredientsService) GetAll(ctx context.Context, request model.IngredientsGetAllRequest) (*model.IngredientsGetManyResponse, error) {
	count := rand.Int() % 40
	var ingredients []model.IngredientsGetByIDResponse
	for i := 0; i != count; i++ {
		if err := ctx.Err(); err != nil {
			return nil, err
		}
		ingredient := model.IngredientsGetByIDResponse{
			ID:              uuid.NewString(),
			Name:            uuid.NewString(),
			MeasurementUnit: uuid.NewString(),
		}
		if request.Name != "" {
			ingredient.Name = request.Name + uuid.NewString()
		}
		ingredients = append(ingredients, ingredient)
	}
	resp := model.IngredientsGetManyResponse(ingredients)
	return &resp, nil
}

func (srv *ingredientsService) GetByID(_ context.Context, request model.IngredientsGetByIDRequest) (*model.IngredientsGetByIDResponse, error) {
	ingredient := model.IngredientsGetByIDResponse{
		ID:              request.ID,
		Name:            uuid.NewString(),
		MeasurementUnit: uuid.NewString(),
	}
	return &ingredient, nil
}
