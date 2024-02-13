//go:generate easyjson -all responses.go
package model

type (
	UsersGetAllResponse struct {
		Count        int32                 `json:"count"`
		NextPage     string                `json:"next"`
		PreviousPage string                `json:"previous"`
		Results      []UserInSubscriptions `json:"results"`
	}
	UsersGetByIDResponse struct {
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
	UsersCreateResponse struct {
		// ID is string.
		//
		// It could be any string, but mostly is UUID v4 string.
		ID        string `json:"id"`
		Email     string `json:"email"`
		Username  string `json:"username"`
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
	}
	BaseErrorResponse struct {
		Detail string `json:"detail"`
	}
	UsersGetMeResponse struct {
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
	TokensLoginResponse struct {
		AuthToken string `json:"auth_token"`
	}
	TagsGetManyResponse    []TagDTO
	TagsGetOneResponse     TagDTO
	RecipesGetManyResponse struct {
		Count    int32       `json:"count"`
		Next     string      `json:"next"`
		Previous string      `json:"previous"`
		Results  []RecipeDTO `json:"results"`
	}
	RecipesCreateResponse              RecipeDTO
	RecipesGetOneResponse              RecipeDTO
	RecipesUpdateOneResponse           RecipeDTO
	RecipesAddedToShoppingCartResponse struct {
		// ID is string.
		//
		// It could be any string, but mostly is UUID v4 string.
		ID          string `json:"id"`
		Name        string `json:"name"`
		Image       string `json:"image"`
		CookingTime uint32 `json:"cooking_time"`
	}
	RecipesAddedToFavoriteResponse struct {
		// ID is string.
		//
		// It could be any string, but mostly is UUID v4 string.
		ID          string `json:"id"`
		Name        string `json:"name"`
		Image       string `json:"image"`
		CookingTime uint32 `json:"cooking_time"`
	}
	RecipesGetSubscribedResponse struct {
		Count    uint32                      `json:"count"`
		Next     string                      `json:"next"`
		Previous string                      `json:"previous"`
		Results  []UserInSubscriptionsResult `json:"results"`
	}
	UserSubscribedResponse     UserInSubscriptionsResult
	IngredientsGetByIDResponse struct {
		// ID is string.
		//
		// It could be any string, but mostly is UUID v4 string.
		ID              string `json:"id"`
		Name            string `json:"name"`
		MeasurementUnit string `json:"measurement_unit"`
	}
	IngredientsGetManyResponse []IngredientsGetByIDResponse
)
