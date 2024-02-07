//go:generate easyjson -all users.go
package model

type (
	UserInSubscriptions struct {
		Email        string `json:"email"`
		ID           string `json:"id"`
		Username     string `json:"username"`
		FirstName    string `json:"first_name"`
		LastName     string `json:"last_name"`
		IsSubscribed bool   `json:"is_subscribed"`
	}
	UserInRecipe struct {
		ID           string `json:"id"`
		Email        string `json:"email"`
		Username     string `json:"username"`
		FirstName    string `json:"first_name"`
		LastName     string `json:"last_name"`
		IsSubscribed bool   `json:"is_subscribed"`
	}
	UserInSubscriptionsResult struct {
		ID           string                       `json:"id"`
		Email        string                       `json:"email"`
		Username     string                       `json:"username"`
		FirstName    string                       `json:"first_name"`
		LastName     string                       `json:"last_name"`
		IsSubscribed bool                         `json:"is_subscribed"`
		Recipes      []RecipeInUsersSubscriptions `json:"recipes"`
		RecipesCount int                          `json:"recipes_count"`
	}
)
