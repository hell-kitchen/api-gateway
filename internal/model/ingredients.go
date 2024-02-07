//go:generate easyjson -all ingredients.go
package model

type (
	IngredientInRecipeCreationDTO struct {
		ID     string `json:"id"`
		Amount int    `json:"amount"`
	}
	IngredientInRecipeDTO struct {
		ID              string `json:"id"`
		Name            string `json:"name"`
		MeasurementUnit string `json:"measurement_unit"`
		Amount          int    `json:"amount"`
	}
)
