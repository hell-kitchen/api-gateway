package convertor

import (
	"github.com/hell-kitchen/api-gateway/internal/model"
	pb "github.com/hell-kitchen/api-gateway/pkg/api/proto/tags"
)

func FromProtobufTag(proto *pb.Tag) model.TagDTO {
	return model.TagDTO{
		ID:    proto.GetId(),
		Name:  proto.GetName(),
		Color: proto.GetColor(),
		Slug:  proto.GetSlug(),
	}
}

func FromProtobufTagsGetAllResponse(proto *pb.GetAllResponse) []model.TagDTO {
	var res = make([]model.TagDTO, 0, len(proto.GetTag()))
	for _, tag := range proto.GetTag() {
		res = append(res, FromProtobufTag(tag))
	}
	return res
}
