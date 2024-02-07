package service

import (
	"context"
	"github.com/hell-kitchen/api-gateway/internal/model"
)

type (
	UsersService interface {
		Create(ctx context.Context, request model.UsersCreateRequest) (*model.UsersCreateResponse, error)
		GetByID(ctx context.Context, request model.UsersGetByIDRequest) (*model.UsersGetByIDResponse, error)
		GetAll(ctx context.Context, request model.UsersGetAllRequest) (*model.UsersGetAllResponse, error)
		GetMe(ctx context.Context) (*model.UsersGetMeResponse, error)
		SetPassword(ctx context.Context, request model.UsersSetPasswordRequest) error
		GetSubscriptions(ctx context.Context, request model.UsersGetSubscriptionsRequest) (*model.UsersGetSubscriptionsRequest, error)
		Subscribe(ctx context.Context, request model.UsersSubscribeRequest) (*model.UserSubscribedResponse, error)
		Unsubscribe(ctx context.Context, request model.UsersUnsubscribeRequest) error
	}
	IngredientsService interface {
		GetAll(ctx context.Context, request model.IngredientsGetAllRequest) (*model.IngredientsGetManyResponse, error)
		GetByID(ctx context.Context, request model.IngredientsGetByIDRequest) (*model.IngredientsGetByIDResponse, error)
	}
	RecipesService interface {
		Create(ctx context.Context, request model.RecipesCreateRequest) (*model.RecipesCreateResponse, error)
		GetByID(ctx context.Context, request model.RecipesGetByIDRequest) (*model.RecipesGetOneResponse, error)
		GetAll(ctx context.Context, request model.RecipesGetAllRequest) (*model.RecipesGetManyResponse, error)
		Update(ctx context.Context, request model.RecipesUpdateByIDRequest) (*model.RecipesUpdateOneResponse, error)
		Delete(ctx context.Context, request model.RecipesDeleteByIDRequest) error
		AddToShoppingCart(ctx context.Context, request model.RecipesAddRecipeToShoppingCartRequest) (*model.RecipesAddedToShoppingCartResponse, error)
		RemoveFromShoppingCart(ctx context.Context, request model.RecipesRemoveRecipeFromShoppingCartRequest) error
		AddToFavorite(ctx context.Context, request model.RecipesAddRecipeToFavoriteRequest) (*model.RecipesAddedToFavoriteResponse, error)
		RemoveFromFavorite(ctx context.Context, request model.RecipesRemoveRecipeFromFavoriteRequest) error
	}
	TagsService interface {
		GetAll(ctx context.Context) (*model.TagsGetManyResponse, error)
		Get(ctx context.Context, request model.TagsGetByIDRequest) (*model.TagsGetOneResponse, error)
	}
	TokensService interface {
		Login(ctx context.Context, request model.TokensLoginRequest) (*model.TokensLoginResponse, error)
		Logout(ctx context.Context) error
	}
	Interface interface {
		Recipes() RecipesService
		Ingredients() IngredientsService
		Users() UsersService
		Tokens() TokensService
		Tags() TagsService
	}
)
