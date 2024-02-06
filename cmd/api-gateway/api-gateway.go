package main

import (
	"github.com/hell-kitchen/api-gateway/internal/config"
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
		),
	)
}
