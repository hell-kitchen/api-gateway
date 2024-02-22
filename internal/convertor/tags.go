package convertor

import (
	"github.com/hell-kitchen/api-gateway/internal/model"
	pb "github.com/hell-kitchen/api-gateway/pkg/api/proto/tags"
)

type TagConvertor struct{}

func Tag() TagConvertor {
	return TagConvertor{}
}

func (TagConvertor) ProtoToDTO(proto *pb.Tag) model.TagDTO {
	return model.TagDTO{
		ID:    proto.GetId(),
		Name:  proto.GetName(),
		Color: proto.GetColor(),
		Slug:  proto.GetSlug(),
	}
}

func (conv TagConvertor) ProtoAllResponseToDTOs(proto *pb.GetAllResponse) []model.TagDTO {
	tags := proto.GetTags()

	var res = make([]model.TagDTO, 0, len(tags))
	for _, tag := range tags {
		res = append(res, conv.ProtoToDTO(tag))
	}

	return res
}
