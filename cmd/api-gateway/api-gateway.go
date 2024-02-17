package main

import (
	"github.com/hell-kitchen/api-gateway/internal/config"
	"github.com/hell-kitchen/api-gateway/internal/controller/http"
	"github.com/hell-kitchen/api-gateway/internal/service/fake"
	"github.com/hell-kitchen/pkg/logger"
	"go.uber.org/fx"
)

func main() {
	fx.New(
		NewOptions(),
	).Run()
}

func NewOptions() fx.Option {
	return fx.Options(
		fx.Provide(
			config.NewServer,
			config.NewTags,
			http.NewServer,
			logger.NewProduction,
			fake.New,
		),
		fx.Invoke(
			addServerStartup,
		),
		fx.NopLogger,
	)
}

func addServerStartup(lc fx.Lifecycle, server *http.Server) {
	lc.Append(fx.Hook{
		OnStart: server.OnStart,
		OnStop:  server.OnStop,
	})
}
