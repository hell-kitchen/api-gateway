//go:generate easyjson -all recipe.go
package model

type (
	RecipeDTO struct {
		ID               string                  `json:"id"`
		Tags             []TagDTO                `json:"tags"`
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
		ID          string `json:"id"`
		Name        string `json:"name"`
		Image       string `json:"image"`
		CookingTime uint32 `json:"cooking_time"`
	}
)
