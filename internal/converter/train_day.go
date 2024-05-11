package converter

import (
	"github.com/kirillmc/platform_common/pkg/nillable"
	"github.com/kirillmc/trainings-server/internal/model"
	desc "github.com/kirillmc/trainings-server/pkg/training_v1"
)

func ToTrainDayFromDesc(trainDay *desc.CreateTrainDayRequest) *model.TrainDayToCreate {
	return &model.TrainDayToCreate{
		DayName:     trainDay.Info.DayName,
		Description: trainDay.Info.Description,
		ProgramId:   trainDay.Info.ProgramId,
	}
}

func ToGetTrainDayResponseFromService(trainDay *model.TrainDay) *desc.TrainDay {
	return &desc.TrainDay{
		Id: trainDay.Id,
		Info: &desc.TrainDayInfo{
			DayName:     trainDay.DayName,
			Description: trainDay.Description,
			ProgramId:   trainDay.ProgramId,
		},
	}
}

func ToGetTrainDaysResponseFromService(trainDays []*model.TrainDay) *desc.GetTrainDaysResponse {
	var trainDayList []*desc.TrainDay
	for _, elem := range trainDays {
		trainDayList = append(trainDayList, ToGetTrainDayResponseFromService(elem))
	}
	return &desc.GetTrainDaysResponse{TrainDays: trainDayList}
}

func ToTrainDayUpdateFromDesc(trainDay *desc.UpdateTrainDayRequest) *model.TrainDayToUpdate {
	return &model.TrainDayToUpdate{
		Id:          trainDay.GetId(),
		DayName:     nillable.Create(trainDay.Info.DayName.Value),
		Description: nillable.Create(trainDay.Info.Description.Value),
	}
}
