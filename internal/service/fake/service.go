package fake

import "github.com/hell-kitchen/api-gateway/internal/service"

var _ service.Interface = (*Service)(nil)

type Service struct {
	recipes     *recipesService
	ingredients *ingredientsService
	users       *usersService
	tokens      *tokensService
	tags        *tagsService
}

func New() (service.Interface, error) {
	s := &Service{}
	newRecipes(s)
	newIngredients(s)
	newUsers(s)
	newTokens(s)
	newTags(s)
	newTags(s)
	newTags(s)
	newTags(s)
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
