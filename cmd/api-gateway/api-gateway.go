package main

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/hell-kitchen/api-gateway/internal/config"
	"github.com/hell-kitchen/api-gateway/internal/controller/http"
	"github.com/hell-kitchen/api-gateway/internal/service"
	"github.com/hell-kitchen/api-gateway/internal/service/fake"
	ingredientsService "github.com/hell-kitchen/api-gateway/internal/service/production/ingredients"
	tagsService "github.com/hell-kitchen/api-gateway/internal/service/production/tags"
	"github.com/hell-kitchen/pkg/logger"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

var (
	buildVersion = "N/A"
	buildDate    = "N/A"
	buildCommit  = "N/A"
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
			getLogger,
			tagsService.New,
			ingredientsService.New,
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

func getLogger() (*zap.Logger, error) {
	log, err := logger.NewProduction(
		zap.Fields(
			logger.WithInstanceID(uuid.NewString()),
			logger.WithService("api-gateway"),
			zap.String("build-commit", buildCommit),
			zap.String("build-date", buildDate),
			zap.String("build-version", buildVersion),
		),
	)
	if err != nil {
		return nil, fmt.Errorf("got non nil error: %w", err)
	}
	return log, nil
}

// addServerStartup starts http server and adds necessary calls on start and stop.
func addServerStartup(lc fx.Lifecycle, server *http.Server) {
	lc.Append(fx.Hook{
		OnStart: server.OnStart,
		OnStop:  server.OnStop,
	})
}
