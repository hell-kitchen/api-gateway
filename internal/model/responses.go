//go:generate easyjson -all responses.go
package model

type (
	UsersGetAllResponse struct {
		Count        int                   `json:"count"`
		NextPage     string                `json:"next"`
		PreviousPage string                `json:"previous"`
		Results      []UserInSubscriptions `json:"results"`
	}
	UsersCreateResponse struct {
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
		Email        string `json:"email"`
		ID           string `json:"id"`
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
		Count    int         `json:"count"`
		Next     string      `json:"next"`
		Previous string      `json:"previous"`
		Results  []RecipeDTO `json:"results"`
	}
	RecipesCreateResponse              RecipeDTO
	RecipesGetOneResponse              RecipeDTO
	RecipesUpdateOneResponse           RecipeDTO
	RecipesAddedToShoppingCartResponse struct {
		ID          string `json:"id"`
		Name        string `json:"name"`
		Image       string `json:"image"`
		CookingTime int    `json:"cooking_time"`
	}
	RecipesGetSubscribedResponse struct {
		Count    int                         `json:"count"`
		Next     string                      `json:"next"`
		Previous string                      `json:"previous"`
		Results  []UserInSubscriptionsResult `json:"results"`
	}
	UserSubscribedResponse     UserInSubscriptionsResult
	IngredientsGetByIDResponse struct {
		Id              int    `json:"id"`
		Name            string `json:"name"`
		MeasurementUnit string `json:"measurement_unit"`
	}
	IngredientsGetManyResponse []IngredientsGetByIDResponse
)
