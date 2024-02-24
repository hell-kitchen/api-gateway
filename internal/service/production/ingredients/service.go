package ingredients

import (
	"github.com/hell-kitchen/api-gateway/internal/service"
	pb "github.com/hell-kitchen/api-gateway/pkg/api/proto/ingredients"
	"github.com/hell-kitchen/pkg/logger"
	"go.uber.org/zap"
)

var _ service.IngredientsService = (*Service)(nil)

type Service struct {
	cli pb.IngredientServiceClient
	log *zap.Logger
}

func New(client pb.IngredientServiceClient, log *zap.Logger) (*Service, error) {
	log = log.With(logger.WithLayer("ingredients-service"))
	log.Info("initialized ingredients service")
	cli := &Service{
		cli: client,
		log: log,
	}
	return cli, nil
}
