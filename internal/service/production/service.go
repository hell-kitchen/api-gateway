package production

import "github.com/hell-kitchen/api-gateway/internal/service"

var _ service.Interface = (*Service)(nil)

type Service struct {
	ingredients service.IngredientsService
	recipes     service.RecipesService
	tags        service.TagsService
	tokens      service.TokensService
	users       service.UsersService
}

// New creates new Service object and return a pointer to it.
func New() (*Service, error) {
	srv := new(Service)
	return srv, nil
}

func (srv *Service) Recipes() service.RecipesService {
	return srv.recipes
}

func (srv *Service) Ingredients() service.IngredientsService {
	return srv.ingredients
}

func (srv *Service) Users() service.UsersService {
	return srv.users
}

func (srv *Service) Tokens() service.TokensService {
	return srv.tokens
}

func (srv *Service) Tags() service.TagsService {
	return srv.tags
}

func (srv *Service) ApplyIngredients(ingredients service.IngredientsService) {
	srv.ingredients = ingredients
}

func (srv *Service) ApplyRecipes(recipes service.RecipesService) {
	srv.recipes = recipes
}

func (srv *Service) ApplyTokens(tokens service.TokensService) {
	srv.tokens = tokens
}

func (srv *Service) ApplyUsers(users service.UsersService) {
	srv.users = users
}

func (srv *Service) ApplyTags(tags service.TagsService) {
	srv.tags = tags
}
