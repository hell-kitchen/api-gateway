//go:generate easyjson -all ingredients.go
package model

type IngredientInRecipeDTO struct {
	ID     string `json:"id"`
	Amount int    `json:"amount"`
}
