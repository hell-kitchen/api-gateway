package ingredients

import (
	"github.com/hell-kitchen/api-gateway/internal/service"
	pb "github.com/hell-kitchen/api-gateway/pkg/api/proto/ingredients"
)

var _ service.IngredientsService = (*Service)(nil)

type Service struct {
	cli pb.IngredientServiceClient
}

func New(client pb.IngredientServiceClient) (*Service, error) {
	cli := &Service{
		cli: client,
	}
	return cli, nil
}
