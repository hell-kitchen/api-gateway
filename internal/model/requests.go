package model

type (
	UsersGetAllRequest struct {
		Page  int `query:"page"`
		Limit int `query:"limit"`
	}
	UsersCreateRequest struct {
		Email     string `json:"email"`
		Username  string `json:"username"`
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
		Password  string `json:"password"`
	}
	UsersGetByIDRequest struct {
		ID string `path:"id"`
	}
	UsersSetPasswordRequest struct {
		NewPassword     string `json:"new_password"`
		CurrentPassword string `json:"current_password"`
	}
	UsersGetSubscriptionsRequest struct {
		Page         int `query:"page"`
		Limit        int `query:"limit"`
		RecipesLimit int `query:"recipes_limit"`
	}
	UsersSubscribeRequest struct {
		ID           string `path:"id"`
		RecipesLimit int    `query:"recipes_limit"`
	}
	UsersUnsubscribeRequest struct {
		ID string `path:"id"`
	}
	TokensLoginRequest struct {
		Password string `json:"password"`
		Email    string `json:"email"`
	}
	TagsGetByIDRequest struct {
		ID string `path:"id"`
	}
	// RecipesGetAllRequest godoc.
	//
	// TODO: check how to check was value in request or not
	RecipesGetAllRequest struct {
		Page             int      `query:"page"`
		Limit            int      `query:"limit"`
		IsFavorited      int      `query:"is_favorited"`
		IsInShoppingCart int      `query:"is_in_shopping_cart"`
		Author           string   `query:"author"`
		Tags             []string `query:"tags"`
	}
	RecipesCreateRequest struct {
		Ingredients []IngredientInRecipeDTO `json:"ingredients"`
		Tags        []int                   `json:"tags"`
		Image       string                  `json:"image"`
		Name        string                  `json:"name"`
		Text        string                  `json:"text"`
		CookingTime int                     `json:"cooking_time"`
	}
	RecipesGetByIDRequest struct {
		ID string `path:"id"`
	}
	RecipesUpdateByIDRequest struct {
		ID          int                     `json:"-" path:"id"`
		Ingredients []IngredientInRecipeDTO `json:"ingredients"`
		Tags        []int                   `json:"tags"`
		Image       string                  `json:"image"`
		Name        string                  `json:"name"`
		Text        string                  `json:"text"`
		CookingTime int                     `json:"cooking_time"`
	}
	RecipesDeleteByIDRequest struct {
		ID string `path:"id"`
	}
	RecipesAddRecipeToShoppingCartRequest struct {
		ID string `path:"id"`
	}
	RecipesRemoveRecipeFromShoppingCartRequest struct {
		ID string `path:"id"`
	}
	RecipesAddRecipeToFavoriteRequest struct {
		ID string `path:"id"`
	}
	RecipesRemoveRecipeFromFavoriteRequest struct {
		ID string `path:"id"`
	}
	IngredientsGetAllRequest struct {
		Name string `json:"-" query:"name"`
	}
	IngredientsGetByIDRequest struct {
		ID int `json:"-" path:"id"`
	}
)
