package main

import (
	"context"
	"encoding/json"
	"fmt"
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
		panic(err.Error())
	}
	fmt.Println(resp)
}

func CreateManyToProto(ingreds []CreateRequest) *pb.CreateManyRequest {
	var req = make([]*pb.CreateRequest, 0, len(ingreds))
	for _, i := range ingreds {
		req = append(req, dto2pb(i))
	}
	return &pb.CreateManyRequest{
		Ingredients: req,
	}
}

func dto2pb(request CreateRequest) *pb.CreateRequest {
	return &pb.CreateRequest{
		Name:            request.Name,
		MeasurementUnit: request.MeasurementUnit,
	}
}
