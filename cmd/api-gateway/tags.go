package main

import (
	"context"
	"github.com/hell-kitchen/api-gateway/internal/config"
	tagsPB "github.com/hell-kitchen/api-gateway/pkg/api/proto/tags"
	"github.com/hell-kitchen/api-gateway/pkg/client/tags"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

func createTagsClient(lc fx.Lifecycle, log *zap.Logger, cfg *config.Tags) (res tagsPB.TagsServiceClient, err error) {
	var cli *tags.Client

	cli, err = tags.New(cfg)
	if err != nil {
		return
	}
	resp, err := cli.GetAll(context.Background(), &tagsPB.GetAllRequest{})
	log.Info("got request", zap.Any("cfg", cfg), zap.Any("response-tags", resp.GetTags()), zap.Error(err))

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			zap.L().Info("tags client on start just called")
			return nil
		},
		OnStop: cli.Stop,
	})
	return cli, nil
}
