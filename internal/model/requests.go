//go:generate easyjson -all requests.go
package model

type (
	UsersGetAllRequest struct {
		Page  int `json:"-" query:"page"`
		Limit int `json:"-" query:"limit"`
	}
	UsersCreateRequest struct {
		Email     string `json:"email"`
		Username  string `json:"username"`
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
		Password  string `json:"password"`
	}
	UsersGetByIDRequest struct {
		ID string `json:"-" path:"id"`
	}
	UsersSetPasswordRequest struct {
		NewPassword     string `json:"new_password"`
		CurrentPassword string `json:"current_password"`
	}
	UsersGetSubscriptionsRequest struct {
		Page         int `json:"-" query:"page"`
		Limit        int `json:"-" query:"limit"`
		RecipesLimit int `json:"-" query:"recipes_limit"`
	}
	UsersSubscribeRequest struct {
		ID           string `json:"-" path:"id"`
		RecipesLimit int    `json:"-" query:"recipes_limit"`
	}
	UsersUnsubscribeRequest struct {
		ID string `path:"id" json:"-"`
	}
	TokensLoginRequest struct {
		Password string `json:"password"`
		Email    string `json:"email"`
	}
	TagsGetByIDRequest struct {
		ID string `path:"id" json:"-"`
	}
	// RecipesGetAllRequest godoc.
	//
	// TODO: check how to check was value in request or not
	RecipesGetAllRequest struct {
		Page             int      `query:"page" json:"-"`
		Limit            int      `query:"limit" json:"-"`
		IsFavorited      int      `query:"is_favorited" json:"-"`
		IsInShoppingCart int      `query:"is_in_shopping_cart" json:"-"`
		Author           string   `query:"author" json:"-"`
		Tags             []string `query:"tags" json:"-"`
	}
	RecipesCreateRequest struct {
		Ingredients []IngredientInRecipeCreationDTO `json:"ingredients"`
		Tags        []int                           `json:"tags"`
		Image       string                          `json:"image"`
		Name        string                          `json:"name"`
		Text        string                          `json:"text"`
		CookingTime int                             `json:"cooking_time"`
	}
	RecipesGetByIDRequest struct {
		ID string `path:"id" json:"-"`
	}
	RecipesUpdateByIDRequest struct {
		ID          int                             `json:"-" path:"id"`
		Ingredients []IngredientInRecipeCreationDTO `json:"ingredients"`
		Tags        []int                           `json:"tags"`
		Image       string                          `json:"image"`
		Name        string                          `json:"name"`
		Text        string                          `json:"text"`
		CookingTime int                             `json:"cooking_time"`
	}
	RecipesDeleteByIDRequest struct {
		ID string `path:"id" json:"-"`
	}
	RecipesAddRecipeToShoppingCartRequest struct {
		ID string `path:"id" json:"-"`
	}
	RecipesRemoveRecipeFromShoppingCartRequest struct {
		ID string `path:"id" json:"-"`
	}
	RecipesAddRecipeToFavoriteRequest struct {
		ID string `path:"id" json:"-"`
	}
	RecipesRemoveRecipeFromFavoriteRequest struct {
		ID string `path:"id" json:"-"`
	}
	IngredientsGetAllRequest struct {
		Name string `json:"-" query:"name"`
	}
	IngredientsGetByIDRequest struct {
		ID int `json:"-" path:"id"`
	}
)
