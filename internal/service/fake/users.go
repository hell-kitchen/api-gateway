package fake

import (
	"context"
	"github.com/google/uuid"
	"github.com/hell-kitchen/api-gateway/internal/model"
	"github.com/hell-kitchen/api-gateway/internal/service"
	"github.com/labstack/gommon/random"
	"math/rand"
)

var (
	_ service.UsersService = (*usersService)(nil)
)

type usersService struct {
	service *Service
}

func newUsers(srv *Service) {
	srv.users = &usersService{
		service: srv,
	}
}

func (u *usersService) Create(_ context.Context, request model.UsersCreateRequest) (*model.UsersCreateResponse, error) {
	resp := &model.UsersCreateResponse{
		ID:        uuid.NewString(),
		Email:     request.Email,
		Username:  request.Username,
		FirstName: request.FirstName,
		LastName:  request.LastName,
	}
	return resp, nil
}

func (u *usersService) GetByID(_ context.Context, request model.UsersGetByIDRequest) (*model.UsersGetByIDResponse, error) {
	resp := model.UsersGetByIDResponse{
		Email:        random.String(8, random.Alphanumeric),
		ID:           request.ID,
		Username:     random.String(8, random.Alphabetic),
		FirstName:    random.String(8, random.Alphabetic),
		LastName:     random.String(8, random.Alphabetic),
		IsSubscribed: rand.Int()%2 == 0,
	}
	return &resp, nil
}

func (u *usersService) GetAll(_ context.Context, _ model.UsersGetAllRequest) (*model.UsersGetAllResponse, error) {
	count := rand.Int() % 200
	var resp []model.UserInSubscriptions
	for i := 0; i != count; i++ {
		resp = append(resp, model.UserInSubscriptions{
			Email:        random.String(10, random.Alphanumeric, "@"),
			ID:           uuid.NewString(),
			Username:     random.String(10, random.Alphanumeric),
			FirstName:    random.String(10, random.Alphanumeric),
			LastName:     random.String(10, random.Alphanumeric),
			IsSubscribed: rand.Int()%2 == 0,
		})
	}
	return &model.UsersGetAllResponse{
		Count:        int32(count),
		NextPage:     "",
		PreviousPage: "",
		Results:      resp,
	}, nil
}

func (u *usersService) GetMe(context.Context) (*model.UsersGetMeResponse, error) {
	return &model.UsersGetMeResponse{
		Email:        uuid.NewString(),
		ID:           uuid.NewString(),
		Username:     uuid.NewString(),
		FirstName:    uuid.NewString(),
		LastName:     uuid.NewString(),
		IsSubscribed: rand.Int()%2 == 0,
	}, nil
}

func (u *usersService) SetPassword(context.Context, model.UsersSetPasswordRequest) error {
	return nil
}

func (u *usersService) GetSubscriptions(_ context.Context, request model.UsersGetSubscriptionsRequest) (*model.RecipesGetSubscribedResponse, error) {
	count := rand.Int() % 100
	users := make([]model.UserInSubscriptionsResult, 0, count)
	for i := 0; i != count; i++ {
		recipesCount := rand.Int() % 100
		recipes := make([]model.RecipeInUsersSubscriptions, 0, recipesCount)
		add := recipesCount % int(request.RecipesLimit)
		for c := 0; c != add; c++ {
			recipes = append(recipes, model.RecipeInUsersSubscriptions{
				ID:          uuid.NewString(),
				Name:        random.String(uint8(rand.Uint32()%20), random.Alphabetic),
				Image:       random.String(uint8(rand.Uint32()%20), random.Alphanumeric),
				CookingTime: uint32(rand.Int()%200) + 1,
			})
		}
		users = append(users, model.UserInSubscriptionsResult{
			ID:           uuid.NewString(),
			Email:        uuid.NewString(),
			Username:     uuid.NewString(),
			FirstName:    uuid.NewString(),
			LastName:     uuid.NewString(),
			IsSubscribed: rand.Int()%2 == 0,
			Recipes:      recipes,
			RecipesCount: uint32(count),
		})
	}
	return &model.RecipesGetSubscribedResponse{
		Count:    uint32(count),
		Next:     "",
		Previous: "",
		Results:  users,
	}, nil

}

func (u *usersService) Subscribe(_ context.Context, request model.UsersSubscribeRequest) (*model.UserSubscribedResponse, error) {
	count := rand.Int() % 100
	recipes := make([]model.RecipeInUsersSubscriptions, 0, count)
	add := count % int(request.RecipesLimit)
	for i := 0; i != add; i++ {
		recipes = append(recipes, model.RecipeInUsersSubscriptions{
			ID:          uuid.NewString(),
			Name:        random.String(uint8(rand.Uint32()%20), random.Alphabetic),
			Image:       random.String(uint8(rand.Uint32()%20), random.Alphanumeric),
			CookingTime: uint32(rand.Int())%200 + 1,
		})
	}
	return &model.UserSubscribedResponse{
		ID:           uuid.NewString(),
		Email:        uuid.NewString(),
		Username:     uuid.NewString(),
		FirstName:    uuid.NewString(),
		LastName:     uuid.NewString(),
		IsSubscribed: true,
		Recipes:      nil,
		RecipesCount: uint32(count),
	}, nil
}

func (u *usersService) Unsubscribe(context.Context, model.UsersUnsubscribeRequest) error {
	return nil
}
