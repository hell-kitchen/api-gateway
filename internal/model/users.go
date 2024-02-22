//go:generate easyjson -all users.go
package model

type (
	UserInSubscriptions struct {
		// ID is string.
		//
		// It could be any string, but mostly is UUID v4 string.
		ID           string `json:"id"`
		Email        string `json:"email"`
		Username     string `json:"username"`
		FirstName    string `json:"first_name"`
		LastName     string `json:"last_name"`
		IsSubscribed bool   `json:"is_subscribed"`
	}
	UserInRecipe struct {
		// ID is string.
		//
		// It could be any string, but mostly is UUID v4 string.
		ID           string `json:"id"`
		Email        string `json:"email"`
		Username     string `json:"username"`
		FirstName    string `json:"first_name"`
		LastName     string `json:"last_name"`
		IsSubscribed bool   `json:"is_subscribed"`
	}
	UserInSubscriptionsResult struct {
		// ID is string.
		//
		// It could be any string, but mostly is UUID v4 string.
		ID           string                       `json:"id"`
		Email        string                       `json:"email"`
		Username     string                       `json:"username"`
		FirstName    string                       `json:"first_name"`
		LastName     string                       `json:"last_name"`
		IsSubscribed bool                         `json:"is_subscribed"`
		Recipes      []RecipeInUsersSubscriptions `json:"recipes"`
		RecipesCount uint32                       `json:"recipes_count"`
	}
)
