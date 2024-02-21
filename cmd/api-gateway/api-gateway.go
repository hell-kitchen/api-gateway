package main

import (
	"context"
	"github.com/hell-kitchen/api-gateway/internal/config"
	"github.com/hell-kitchen/api-gateway/internal/controller/http"
	"github.com/hell-kitchen/api-gateway/internal/service"
	"github.com/hell-kitchen/api-gateway/internal/service/fake"
	tagsService "github.com/hell-kitchen/api-gateway/internal/service/production/tags"
	pb "github.com/hell-kitchen/api-gateway/pkg/api/proto/tags"
	"github.com/hell-kitchen/api-gateway/pkg/client/tags"
	"github.com/hell-kitchen/pkg/logger"
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
			config.NewTags,
			http.NewServer,
			logger.NewProduction,
			tagsService.New,
			createTagsClient,
			fx.Annotate(service.New, fx.As(new(service.Interface))),
		),
		fx.Invoke(
			fake.NewRecipes,
			fake.NewTokens,
			fake.NewIngredients,
			fake.NewUsers,
			applyTagsService,

			addServerStartup,
		),
		//fx.WithLogger(func(log *zap.Logger) fxevent.Logger {
		//	return &fxevent.ZapLogger{Logger: log}
		//}),
	)
}

func applyTagsService(srv service.Interface, tag *tagsService.Service) {
	srv.ApplyTags(tag)
}

func createTagsClient(lc fx.Lifecycle, cfg *config.Tags) (res pb.TagsServiceClient, err error) {
	var cli *tags.Client

	cli, err = tags.New(cfg)
	if err != nil {
		return
	}

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			zap.L().Info("tags client on start just called")
			return nil
		},
		OnStop: cli.Stop,
	})
	return cli, nil
}

func addServerStartup(lc fx.Lifecycle, server *http.Server) {
	lc.Append(fx.Hook{
		OnStart: server.OnStart,
		OnStop:  server.OnStop,
	})
}
