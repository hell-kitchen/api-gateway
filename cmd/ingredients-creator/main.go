package main

import (
	"context"
	"encoding/json"
	"fmt"
	"go.uber.org/zap"

	"github.com/hell-kitchen/api-gateway/internal/config"
	pb "github.com/hell-kitchen/api-gateway/pkg/api/proto/ingredients"
	"github.com/hell-kitchen/api-gateway/pkg/client/ingredients"
)

type CreateRequest struct {
	Name            string `json:"name"`
	MeasurementUnit string `json:"measurement_unit"`
}

func main() {
	var (
		data []CreateRequest
		err  error
		cfg  *config.Ingredients
		cli  *ingredients.Client
		ctx  = context.Background()
	)

	if err = json.Unmarshal(ingredientsRawData, &data); err != nil {
		panic(err.Error())
	}
	cfg, err = config.NewIngredients()
	if err != nil {
		panic(err.Error())
	}
	cli, err = ingredients.New(cfg)
	if err != nil {
		panic(err.Error())
	}
	resp, err := cli.CreateMany(ctx, CreateManyToProto(data))
	if err != nil {
		fmt.Println(err)
	}
	zap.L().Debug("created", zap.Int("created_count", len(resp.GetIngredients())))
}

// CreateManyToProto converts slice of CreateRequest objects
// to protobuf CreateManyRequest.
func CreateManyToProto(ingredients []CreateRequest) *pb.CreateManyRequest {
	var req = make([]*pb.CreateRequest, 0, len(ingredients))
	for _, ingredient := range ingredients {
		req = append(req, convert(ingredient))
	}

	return &pb.CreateManyRequest{
		Ingredients: req,
	}
}

// convert converts CreateRequest to protobuf CreateRequest.
func convert(request CreateRequest) *pb.CreateRequest {
	return &pb.CreateRequest{
		Name:            request.Name,
		MeasurementUnit: request.MeasurementUnit,
	}
}
