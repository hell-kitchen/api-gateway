package fake

import "github.com/hell-kitchen/api-gateway/internal/service"

var _ service.Interface = (*Service)(nil)

type Service struct {
	recipes     service.RecipesService
	ingredients service.IngredientsService
	users       service.UsersService
	tokens      service.TokensService
	tags        service.TagsService
}

func New() (service.Interface, error) {
	s := &Service{}
	NewRecipes(s)
	NewIngredients(s)
	NewUsers(s)
	NewTokens(s)
	NewTags(s)
	return s, nil
}

func (s *Service) Recipes() service.RecipesService {
	return s.recipes
}

func (s *Service) Ingredients() service.IngredientsService {
	return s.ingredients
}

func (s *Service) Users() service.UsersService {
	return s.users
}

func (s *Service) Tokens() service.TokensService {
	return s.tokens
}

func (s *Service) Tags() service.TagsService {
	return s.tags
}

func (s *Service) ApplyRecipes(srv service.RecipesService) {
	s.recipes = srv
}

func (s *Service) ApplyIngredients(srv service.IngredientsService) {
	s.ingredients = srv
}

func (s *Service) ApplyUsers(srv service.UsersService) {
	s.users = srv
}

func (s *Service) ApplyTokens(srv service.TokensService) {
	s.tokens = srv
}

func (s *Service) ApplyTags(srv service.TagsService) {
	s.tags = srv
}
