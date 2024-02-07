package fake

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/hell-kitchen/api-gateway/internal/model"
	"github.com/hell-kitchen/api-gateway/internal/service"
	"github.com/labstack/gommon/random"
	"math/rand"
)

var (
	_ service.RecipesService = (*recipesService)(nil)
)

func newRecipes(srv *Service) {
	srv.recipes = &recipesService{
		service: srv,
	}
}

type recipesService struct {
	service *Service
}

func (r *recipesService) Create(ctx context.Context, request model.RecipesCreateRequest) (*model.RecipesCreateResponse, error) {
	tags := make([]model.TagDTO, 0, len(request.Tags))
	ingredients := make([]model.IngredientInRecipeDTO, 0, len(request.Ingredients))
	for _, id := range request.Tags {

		if err := ctx.Err(); err != nil {
			return nil, err
		}

		tags = append(tags, model.TagDTO{
			ID:    id,
			Name:  random.String(7, random.Alphabetic),
			Color: random.String(6, random.Hex),
			Slug:  random.String(15, random.Alphanumeric),
		})
	}

	for _, ingredient := range request.Ingredients {
		if err := ctx.Err(); err != nil {
			return nil, err
		}

		ingredients = append(ingredients, model.IngredientInRecipeDTO{
			ID:              ingredient.ID,
			Name:            random.String(8, random.Alphabetic),
			MeasurementUnit: random.String(3, random.Alphabetic),
			Amount:          ingredient.Amount,
		})
	}

	return &model.RecipesCreateResponse{
		ID:   uuid.NewString(),
		Tags: tags,
		Author: model.UserInRecipe{
			ID:           uuid.NewString(),
			Email:        fmt.Sprintf("%s@ya.ru", random.String(8, random.Alphanumeric, "._")),
			Username:     random.String(7, random.Alphabetic, "_-"),
			FirstName:    random.String(7, random.Alphabetic),
			LastName:     random.String(7, random.Alphabetic),
			IsSubscribed: rand.Int()%2 == 0,
		},
		Ingredients:      ingredients,
		IsFavorited:      rand.Int()%2 == 0,
		IsInShoppingCart: rand.Int()%2 == 0,
		Name:             request.Name,
		Image:            request.Image,
		Text:             request.Text,
		CookingTime:      request.CookingTime,
	}, nil
}

func (r *recipesService) GetByID(ctx context.Context, request model.RecipesGetByIDRequest) (*model.RecipesGetOneResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (r *recipesService) GetAll(ctx context.Context, request model.RecipesGetAllRequest) (*model.RecipesGetManyResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (r *recipesService) Update(ctx context.Context, request model.RecipesUpdateByIDRequest) (*model.RecipesUpdateOneResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (r *recipesService) Delete(ctx context.Context, request model.RecipesDeleteByIDRequest) error {
	//TODO implement me
	panic("implement me")
}

func (r *recipesService) AddToShoppingCart(ctx context.Context, request model.RecipesAddRecipeToShoppingCartRequest) (*model.RecipesAddedToShoppingCartResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (r *recipesService) RemoveFromShoppingCart(ctx context.Context, request model.RecipesRemoveRecipeFromShoppingCartRequest) error {
	//TODO implement me
	panic("implement me")
}

func (r *recipesService) AddToFavorite(ctx context.Context, request model.RecipesAddRecipeToFavoriteRequest) (*model.RecipesAddedToFavoriteResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (r *recipesService) RemoveFromFavorite(ctx context.Context, request model.RecipesRemoveRecipeFromFavoriteRequest) error {
	//TODO implement me
	panic("implement me")
}
