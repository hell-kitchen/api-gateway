package tags

import (
	"github.com/hell-kitchen/api-gateway/internal/service"
	pb "github.com/hell-kitchen/api-gateway/pkg/api/proto/tags"
	"github.com/hell-kitchen/pkg/logger"
	"go.uber.org/zap"
)

var _ service.TagsService = (*Service)(nil)

type Service struct {
	cli pb.TagsServiceClient
	log *zap.Logger
}

func New(client pb.TagsServiceClient, log *zap.Logger) (*Service, error) {
	log = log.With(logger.WithLayer("tags-service"))
	log.Info("initialized tag service")
	cli := &Service{
		cli: client,
		log: log,
	}
	return cli, nil
}
