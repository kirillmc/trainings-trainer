package converter

import (
	"github.com/kirillmc/platform_common/pkg/nillable"
	"github.com/kirillmc/trainings-server/internal/model"
	desc "github.com/kirillmc/trainings-server/pkg/training_v1"
)

func ToSetFromDesc(set *desc.CreateSetRequest) *model.SetToCreate {
	return &model.SetToCreate{
		Quantity:   set.Info.Quantity,
		Weight:     set.Info.Weight,
		ExerciseId: set.Info.ExerciseId,
	}
}

func ToGetSetResponseFromService(set *model.Set) *desc.Set {
	return &desc.Set{
		Id: set.Id,
		Info: &desc.SetInfo{
			Quantity:   set.Quantity,
			Weight:     set.Weight,
			ExerciseId: set.ExerciseId,
		},
	}
}

func ToGetSetsResponseFromService(sets []*model.Set) *desc.GetSetsResponse {
	var setList []*desc.Set
	for _, elem := range sets {
		setList = append(setList, ToGetSetResponseFromService(elem))
	}
	return &desc.GetSetsResponse{Sets: setList}
}

func ToSetUpdateFromDesc(set *desc.UpdateSetRequest) *model.SetToUpdate {
	return &model.SetToUpdate{
		Id:       set.GetId(),
		Quantity: nillable.CreateNillableInt(set.Info.Quantity.Value),
		Weight:   nillable.CreateNillableDouble(set.Info.Weight.Value),
	}
}
