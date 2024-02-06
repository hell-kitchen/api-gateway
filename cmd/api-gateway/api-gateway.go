package main

import (
	"github.com/hell-kitchen/api-gateway/internal/config"
	"github.com/hell-kitchen/api-gateway/internal/controller/http"
	"go.uber.org/fx"
	"go.uber.org/zap"
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
			http.NewServer,
			zap.NewProduction,
		),
		fx.Invoke(
			addServerStartup,
		),
	)
}

func addServerStartup(lc fx.Lifecycle, server *http.Server) {
	lc.Append(fx.Hook{
		OnStart: server.OnStart,
		OnStop:  server.OnStop,
	})
}
