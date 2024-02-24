package tags

import (
	"context"

	"github.com/hell-kitchen/api-gateway/internal/convertor"
	"github.com/hell-kitchen/api-gateway/internal/model"
	pb "github.com/hell-kitchen/api-gateway/pkg/api/proto/tags"
)

func (cli *Service) GetAll(ctx context.Context) (*model.TagsGetManyResponse, error) {
	protoResponse, err := cli.cli.GetAll(ctx, &pb.GetAllRequest{})
	if err != nil {
		return nil, err
	}

	tags := convertor.Tag().ProtoAllResponseToDTOs(protoResponse)
	res := model.TagsGetManyResponse(tags)

	return &res, nil
}

func (cli *Service) Get(ctx context.Context, request model.TagsGetByIDRequest) (*model.TagsGetOneResponse, error) {
	protoResult, err := cli.cli.Get(ctx, &pb.GetRequest{Id: request.ID})
	if err != nil {
		return nil, err
	}
	tag := convertor.Tag().ProtoToDTO(protoResult.GetTag())
	return (*model.TagsGetOneResponse)(&tag), nil
}
