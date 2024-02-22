package convertor

import (
	"github.com/hell-kitchen/api-gateway/internal/model"
	pb "github.com/hell-kitchen/api-gateway/pkg/api/proto/ingredients"
)

type IngredientConvertor struct{}

func Ingredient() IngredientConvertor {
	return IngredientConvertor{}
}

func (conv IngredientConvertor) AllResponseToGetManyResponse(proto *pb.GetAllResponse) model.IngredientsGetManyResponse {
	ingredients := proto.GetIngredients()
	var resp = make([]model.IngredientsGetByIDResponse, 0, len(ingredients))
	for _, ingredient := range ingredients {
		resp = append(resp, conv.ProtoToDTO(ingredient))
	}
	return resp
}
func (IngredientConvertor) ProtoToDTO(proto *pb.Ingredient) model.IngredientsGetByIDResponse {
	return model.IngredientsGetByIDResponse{
		ID:              proto.GetId(),
		Name:            proto.GetName(),
		MeasurementUnit: proto.GetMeasurementUnit(),
	}
}
