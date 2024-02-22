package service

import (
	"sync"
)

var (
	_global Interface = (*Service)(nil)
	_mu     sync.Mutex
)

type Service struct {
	ingredients IngredientsService
	recipes     RecipesService
	tags        TagsService
	tokens      TokensService
	users       UsersService
	mu          sync.Mutex
}

// New creates new Service object and return a pointer to it.
func New() (*Service, error) {
	srv := new(Service)
	_mu.Lock()
	_global = srv
	_mu.Unlock()
	return srv, nil
}

func Get() Interface {
	_mu.Lock()
	srv := _global
	_mu.Unlock()
	return srv
}

func (srv *Service) Recipes() RecipesService {
	return srv.recipes
}

func (srv *Service) Ingredients() IngredientsService {
	return srv.ingredients
}

func (srv *Service) Users() UsersService {
	return srv.users
}

func (srv *Service) Tokens() TokensService {
	return srv.tokens
}

func (srv *Service) Tags() TagsService {
	return srv.tags
}

func (srv *Service) ApplyIngredients(ingredients IngredientsService) {
	srv.ingredients = ingredients
}

func (srv *Service) ApplyRecipes(recipes RecipesService) {
	srv.recipes = recipes
}

func (srv *Service) ApplyTokens(tokens TokensService) {
	srv.tokens = tokens
}

func (srv *Service) ApplyUsers(users UsersService) {
	srv.users = users
}

func (srv *Service) ApplyTags(tags TagsService) {
	srv.tags = tags
}
