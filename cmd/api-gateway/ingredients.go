package main

import (
	"context"
	"github.com/hell-kitchen/api-gateway/internal/config"
	pb "github.com/hell-kitchen/api-gateway/pkg/api/proto/ingredients"
	"github.com/hell-kitchen/api-gateway/pkg/client/ingredients"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

func createIngredientsClient(lc fx.Lifecycle, log *zap.Logger, cfg *config.Ingredients) (res pb.IngredientServiceClient, err error) {
	var cli *ingredients.Client
	var resp *pb.GetAllResponse

	cli, err = ingredients.New(cfg)
	if err != nil {
		return
	}
	resp, err = cli.GetAll(context.Background(), &pb.GetAllRequest{})
	log.Info("got response from ingredient service", zap.Any("cfg", cfg), zap.Any("response-ingredients", resp.GetIngredients()), zap.Error(err))

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			zap.L().Info("tags client on start just called")
			return nil
		},
		OnStop: cli.Stop,
	})
	return cli, nil
}
