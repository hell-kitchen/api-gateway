package tags

import (
	"github.com/hell-kitchen/api-gateway/internal/service"
	pb "github.com/hell-kitchen/api-gateway/pkg/api/proto/tags"
)

var _ service.TagsService = (*Service)(nil)

type Service struct {
	cli pb.TagsServiceClient
}

func New(client pb.TagsServiceClient) (*Service, error) {
	cli := &Service{
		cli: client,
	}
	return cli, nil
}
