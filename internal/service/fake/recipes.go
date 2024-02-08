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

func (r *recipesService) GetByID(_ context.Context, request model.RecipesGetByIDRequest) (*model.RecipesGetOneResponse, error) {
	return &model.RecipesGetOneResponse{
		ID: request.ID,
		Tags: []model.TagDTO{
			{
				ID:    uuid.NewString(),
				Name:  uuid.NewString(),
				Color: random.String(6, random.Hex),
				Slug:  uuid.NewString(),
			},
		},
		Author: model.UserInRecipe{
			ID:           uuid.NewString(),
			Email:        random.String(8, random.Alphabetic),
			Username:     random.String(8, random.Alphabetic),
			FirstName:    random.String(8, random.Alphabetic),
			LastName:     random.String(8, random.Alphabetic),
			IsSubscribed: rand.Int()%2 == 0,
		},
		Ingredients:      []model.IngredientInRecipeDTO{},
		IsFavorited:      rand.Int()%2 == 0,
		IsInShoppingCart: rand.Int()%2 == 0,
		Name:             random.String(8, random.Alphabetic),
		Image:            random.String(20, random.Alphanumeric),
		Text:             random.String(100, random.Alphabetic, " "),
		CookingTime:      rand.Int() % 100,
	}, nil
}

func (r *recipesService) GetAll(_ context.Context, req model.RecipesGetAllRequest) (*model.RecipesGetManyResponse, error) {
	return &model.RecipesGetManyResponse{
		Count:    0,
		Next:     "",
		Previous: "",
		Results: []model.RecipeDTO{
			{
				ID:               uuid.NewString(),
				Tags:             []model.TagDTO{},
				Author:           model.UserInRecipe{},
				Ingredients:      []model.IngredientInRecipeDTO{},
				IsFavorited:      rand.Int()%2 == 0 || req.IsFavorited == 1,
				IsInShoppingCart: rand.Int()%2 == 0 || req.IsInShoppingCart == 1,
				Name:             random.String(10, random.Alphabetic),
				Image:            random.String(10, random.Alphabetic),
				Text:             random.String(10, random.Alphabetic),
				CookingTime:      rand.Int()%100 + 1,
			},
		},
	}, nil
}

func (r *recipesService) Update(_ context.Context, request model.RecipesUpdateByIDRequest) (*model.RecipesUpdateOneResponse, error) {
	return &model.RecipesUpdateOneResponse{
		ID:               request.ID,
		Tags:             nil,
		Author:           model.UserInRecipe{},
		Ingredients:      nil,
		IsFavorited:      false,
		IsInShoppingCart: false,
		Name:             request.Name,
		Image:            request.Image,
		Text:             request.Text,
		CookingTime:      request.CookingTime,
	}, nil
}

func (r *recipesService) Delete(context.Context, model.RecipesDeleteByIDRequest) error {
	return nil
}

func (r *recipesService) AddToShoppingCart(_ context.Context, request model.RecipesAddRecipeToShoppingCartRequest) (*model.RecipesAddedToShoppingCartResponse, error) {
	return &model.RecipesAddedToShoppingCartResponse{
		ID:          request.ID,
		Name:        uuid.NewString(),
		Image:       random.String(uint8(rand.Int()%40+1), random.Alphabetic),
		CookingTime: rand.Int() % 100,
	}, nil
}

func (r *recipesService) RemoveFromShoppingCart(context.Context, model.RecipesRemoveRecipeFromShoppingCartRequest) error {
	return nil
}

func (r *recipesService) AddToFavorite(_ context.Context, request model.RecipesAddRecipeToFavoriteRequest) (*model.RecipesAddedToFavoriteResponse, error) {
	return &model.RecipesAddedToFavoriteResponse{
		ID:          request.ID,
		Name:        uuid.NewString(),
		Image:       random.String(uint8(rand.Int()%40+1), random.Alphabetic),
		CookingTime: rand.Int() % 100,
	}, nil
}

func (r *recipesService) RemoveFromFavorite(context.Context, model.RecipesRemoveRecipeFromFavoriteRequest) error {
	return nil
}
