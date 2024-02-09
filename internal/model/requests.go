//go:generate easyjson -all requests.go
package model

type (
	UsersGetAllRequest struct {
		Page  int32 `json:"-" query:"page"`
		Limit int32 `json:"-" query:"limit"`
	}
	UsersCreateRequest struct {
		Email     string `json:"email"`
		Username  string `json:"username"`
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
		Password  string `json:"password"`
	}
	UsersGetByIDRequest struct {
		ID string `json:"-" param:"id"`
	}
	UsersSetPasswordRequest struct {
		NewPassword     string `json:"new_password"`
		CurrentPassword string `json:"current_password"`
	}
	UsersGetSubscriptionsRequest struct {
		Page         int32 `json:"-" query:"page"`
		Limit        int32 `json:"-" query:"limit"`
		RecipesLimit int32 `json:"-" query:"recipes_limit"`
	}
	UsersSubscribeRequest struct {
		ID           string `json:"-" param:"id"`
		RecipesLimit int32  `json:"-" query:"recipes_limit"`
	}
	UsersUnsubscribeRequest struct {
		ID string `param:"id" json:"-"`
	}
	TokensLoginRequest struct {
		Password string `json:"password"`
		Email    string `json:"email"`
	}
	TagsGetByIDRequest struct {
		ID string `param:"id" json:"-"`
	}
	// RecipesGetAllRequest godoc.
	//
	// TODO: check how to check was value in request or not
	RecipesGetAllRequest struct {
		Page             int32    `query:"page" json:"-"`
		Limit            int32    `query:"limit" json:"-"`
		IsFavorited      int32    `query:"is_favorited" json:"-"`
		IsInShoppingCart int32    `query:"is_in_shopping_cart" json:"-"`
		Author           string   `query:"author" json:"-"`
		Tags             []string `query:"tags" json:"-"`
	}
	RecipesCreateRequest struct {
		Ingredients []IngredientInRecipeCreationDTO `json:"ingredients"`
		Tags        []string                        `json:"tags"`
		Image       string                          `json:"image"`
		Name        string                          `json:"name"`
		Text        string                          `json:"text"`
		CookingTime uint32                          `json:"cooking_time"`
	}
	RecipesGetByIDRequest struct {
		ID string `param:"id" json:"-"`
	}
	RecipesUpdateByIDRequest struct {
		ID          string                          `json:"-" param:"id"`
		Ingredients []IngredientInRecipeCreationDTO `json:"ingredients"`
		Tags        []string                        `json:"tags"`
		Image       string                          `json:"image"`
		Name        string                          `json:"name"`
		Text        string                          `json:"text"`
		CookingTime int32                           `json:"cooking_time"`
	}
	RecipesDeleteByIDRequest struct {
		ID string `param:"id" json:"-"`
	}
	RecipesAddRecipeToShoppingCartRequest struct {
		ID string `param:"id" json:"-"`
	}
	RecipesRemoveRecipeFromShoppingCartRequest struct {
		ID string `param:"id" json:"-"`
	}
	RecipesAddRecipeToFavoriteRequest struct {
		ID string `param:"id" json:"-"`
	}
	RecipesRemoveRecipeFromFavoriteRequest struct {
		ID string `param:"id" json:"-"`
	}
	IngredientsGetAllRequest struct {
		Name string `json:"-" query:"name"`
	}
	IngredientsGetByIDRequest struct {
		ID string `json:"-" param:"id"`
	}
)
