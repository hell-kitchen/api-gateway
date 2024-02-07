package fake

import (
	"context"
	"github.com/hell-kitchen/api-gateway/internal/model"
	"github.com/hell-kitchen/api-gateway/internal/service"
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

func (u *usersService) Create(ctx context.Context, request model.UsersCreateRequest) (*model.UsersCreateResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (u *usersService) GetByID(ctx context.Context, request model.UsersGetByIDRequest) (*model.UsersGetByIDResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (u *usersService) GetAll(ctx context.Context, request model.UsersGetAllRequest) (*model.UsersGetAllResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (u *usersService) GetMe(ctx context.Context) (*model.UsersGetMeResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (u *usersService) SetPassword(ctx context.Context, request model.UsersSetPasswordRequest) error {
	//TODO implement me
	panic("implement me")
}

func (u *usersService) GetSubscriptions(ctx context.Context, request model.UsersGetSubscriptionsRequest) (*model.UsersGetSubscriptionsRequest, error) {
	//TODO implement me
	panic("implement me")
}

func (u *usersService) Subscribe(ctx context.Context, request model.UsersSubscribeRequest) (*model.UserSubscribedResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (u *usersService) Unsubscribe(ctx context.Context, request model.UsersUnsubscribeRequest) error {
	//TODO implement me
	panic("implement me")
}
