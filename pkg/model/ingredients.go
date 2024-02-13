//go:generate easyjson -all ingredients.go
package model

type (
	// IngredientInRecipeCreationDTO godoc.
	IngredientInRecipeCreationDTO struct {
		// ID is string.
		//
		// It could be any string, but mostly is UUID v4 string.
		ID string `json:"id"`
		// Amount is amount of ingredient in it's measurement unit.
		Amount uint32 `json:"amount"`
	}
	IngredientInRecipeDTO struct {
		// ID is string.
		//
		// It could be any string, but mostly is UUID v4 string.
		ID string `json:"id"`
		// Name is public name of recipe.
		Name string `json:"name"`
		// MeasurementUnit is measurement unit of ingredient.
		MeasurementUnit string `json:"measurement_unit"`
		// Amount is amount of ingredient in recipe.
		Amount uint32 `json:"amount"`
	}
)
