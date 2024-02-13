//go:generate easyjson -all recipe.go
package model

type (
	RecipeDTO struct {
		// ID is string.
		//
		// It could be any string, but mostly is UUID v4 string.
		ID string `json:"id"`
		// Tags is array of tags in recipe.
		Tags []TagDTO `json:"tags"`
		// Author is info object of creator of recipe.
		Author           UserInRecipe            `json:"author"`
		Ingredients      []IngredientInRecipeDTO `json:"ingredients"`
		IsFavorited      bool                    `json:"is_favorited"`
		IsInShoppingCart bool                    `json:"is_in_shopping_cart"`
		Name             string                  `json:"name"`
		Image            string                  `json:"image"`
		Text             string                  `json:"text"`
		CookingTime      uint32                  `json:"cooking_time"`
	}
	RecipeInUsersSubscriptions struct {
		// ID is string.
		//
		// It could be any string, but mostly is UUID v4 string.
		ID          string `json:"id"`
		Name        string `json:"name"`
		Image       string `json:"image"`
		CookingTime uint32 `json:"cooking_time"`
	}
)
