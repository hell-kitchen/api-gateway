package main

import (
	"github.com/hell-kitchen/api-gateway/internal/config"
	"github.com/hell-kitchen/api-gateway/internal/controller/http"
	"github.com/hell-kitchen/api-gateway/internal/service"
	"github.com/hell-kitchen/api-gateway/internal/service/fake"
	ingredientsService "github.com/hell-kitchen/api-gateway/internal/service/production/ingredients"
	tagsService "github.com/hell-kitchen/api-gateway/internal/service/production/tags"
	"github.com/hell-kitchen/pkg/logger"
	"go.uber.org/fx"
)

func main() {
	fx.New(
		NewOptions(),
	).Run()
}

// NewOptions creates options to build and run main application.
func NewOptions() fx.Option {
	return fx.Options(
		fx.Provide(
			config.NewServer,
			config.NewTags,
			config.NewIngredients,
			http.NewServer,
			logger.NewProduction,
			tagsService.New,
			ingredientsService.New,
			createTagsClient,
			createIngredientsClient,
			fx.Annotate(service.New, fx.As(new(service.Interface))),
		),
		fx.Invoke(
			fake.NewRecipes,
			fake.NewTokens,
			fake.NewUsers,
			fake.NewIngredients,
			fake.NewTags,

			addServerStartup,
		),
		fx.NopLogger,
	)
}

// addServerStartup starts http server and adds necessary calls on start and stop.
func addServerStartup(lc fx.Lifecycle, server *http.Server) {
	lc.Append(fx.Hook{
		OnStart: server.OnStart,
		OnStop:  server.OnStop,
	})
}
