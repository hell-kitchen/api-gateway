package service

import (
	"context"

	"github.com/hell-kitchen/api-gateway/pkg/model"
)

type (
	// UsersService encapsulates all underlying logic of working with Users domain zone.
	UsersService interface {
		// Create registers user by provided request data and return created object.
		Create(ctx context.Context, request model.UsersCreateRequest) (*model.UsersCreateResponse, error)
		// GetByID returns user by request data.
		GetByID(ctx context.Context, request model.UsersGetByIDRequest) (*model.UsersGetByIDResponse, error)
		// GetAll return all users with pagination parameters.
		GetAll(ctx context.Context, request model.UsersGetAllRequest) (*model.UsersGetAllResponse, error)
		// GetMe return user if it is loggined.
		GetMe(ctx context.Context) (*model.UsersGetMeResponse, error)
		// SetPassword sets password of loggined user.
		SetPassword(ctx context.Context, request model.UsersSetPasswordRequest) error
		// GetSubscriptions returns subscriptions of current user.
		GetSubscriptions(ctx context.Context, request model.UsersGetSubscriptionsRequest) (*model.RecipesGetSubscribedResponse, error)
		// Subscribe creates subscription relation between logined user and user from request data.
		Subscribe(ctx context.Context, request model.UsersSubscribeRequest) (*model.UserSubscribedResponse, error)
		// Unsubscribe deletes subscription relation between users.
		Unsubscribe(ctx context.Context, request model.UsersUnsubscribeRequest) error
	}
	// IngredientsService encapsulates all neccessary logic to work with ingredients.
	IngredientsService interface {
		// GetAll returns all ingredients.
		GetAll(ctx context.Context, request model.IngredientsGetAllRequest) (*model.IngredientsGetManyResponse, error)
		// GetByID returns ingredient by id.
		GetByID(ctx context.Context, request model.IngredientsGetByIDRequest) (*model.IngredientsGetByIDResponse, error)
	}
	// RecipesService encapsulates all logic to work with recipes.
	RecipesService interface {
		// Create creates new recipe.
		Create(ctx context.Context, request model.RecipesCreateRequest) (*model.RecipesCreateResponse, error)
		// GetByID returns recipe by id.
		GetByID(ctx context.Context, request model.RecipesGetByIDRequest) (*model.RecipesGetOneResponse, error)
		// GetAll returns all recipes.
		GetAll(ctx context.Context, request model.RecipesGetAllRequest) (*model.RecipesGetManyResponse, error)
		// Update updates recipe by id with provided data.
		Update(ctx context.Context, request model.RecipesUpdateByIDRequest) (*model.RecipesUpdateOneResponse, error)
		// Delete deletes recipe by id.
		Delete(ctx context.Context, request model.RecipesDeleteByIDRequest) error
		// AddToShoppingCart adds recipe to shopping cart.
		AddToShoppingCart(ctx context.Context, request model.RecipesAddRecipeToShoppingCartRequest) (*model.RecipesAddedToShoppingCartResponse, error)
		// RemoveFromShoppingCart deletes recipe from shopping cart of current user.
		RemoveFromShoppingCart(ctx context.Context, request model.RecipesRemoveRecipeFromShoppingCartRequest) error
		// AddToFavorite adds recipe to favorites of current user.
		AddToFavorite(ctx context.Context, request model.RecipesAddRecipeToFavoriteRequest) (*model.RecipesAddedToFavoriteResponse, error)
		// RemoveFromFavorite removes recipe from current user's favorites.
		RemoveFromFavorite(ctx context.Context, request model.RecipesRemoveRecipeFromFavoriteRequest) error
	}
	// TagsService encapsulates all the logic to work with Tags domain zone
	TagsService interface {
		// GetAll returns all tags.
		GetAll(ctx context.Context) (*model.TagsGetManyResponse, error)
		// GetAll returns tags by id.
		Get(ctx context.Context, request model.TagsGetByIDRequest) (*model.TagsGetOneResponse, error)
	}
	TokensService interface {
		// Login creates token for user, based on login data.
		Login(ctx context.Context, request model.TokensLoginRequest) (*model.TokensLoginResponse, error)
		// Logout deletes token.
		Logout(ctx context.Context) error
	}
	Interface interface {
		// Recipes is accessor to RecipesService
		Recipes() RecipesService
		// Ingredients is accessor to IngredientsService
		Ingredients() IngredientsService
		// Users is accessor to UsersService
		Users() UsersService
		// Tokens is accessor to TokensService
		Tokens() TokensService
		// Tags is accessor to TagsService
		Tags() TagsService
	}
)
