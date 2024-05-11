package converter

import (
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/kirillmc/platform_common/pkg/nillable"
	"github.com/kirillmc/trainings-server/internal/model"
	desc "github.com/kirillmc/trainings-server/pkg/training_v1"
)

func ToStatisticFromDesc(set *desc.CreateStatisticRequest) *model.StatisticToCreate {
	return &model.StatisticToCreate{
		Quantity:   set.Info.Quantity,
		Weight:     set.Info.Weight,
		ProgramId:  set.Info.ProgramId,
		TrainDayId: set.Info.TrainDayId,
		ExerciseId: set.Info.ExerciseId,
		SetId:      set.Info.SetId,
	}
}

func ToGetStatisticResponseFromService(statistic *model.Statistic) *desc.Statistic {
	return &desc.Statistic{
		Id: statistic.Id,
		Info: &desc.StatisticInfo{
			Quantity:   statistic.Quantity,
			Weight:     statistic.Weight,
			ProgramId:  statistic.ProgramId,
			TrainDayId: statistic.TrainDayId,
			ExerciseId: statistic.ExerciseId,
			SetId:      statistic.SetId,
		},
		Time: timestamppb.New(statistic.Time),
	}
}

func ToStatitsticUpdateFromDesc(statistic *desc.UpdateStatisticRequest) *model.StatisticToUpdate {
	return &model.StatisticToUpdate{
		Id:       statistic.GetId(),
		Quantity: nillable.CreateNillableInt(statistic.Info.Quantity.Value),
		Weight:   nillable.CreateNillableDouble(statistic.Info.Weight.Value),
	}
}
