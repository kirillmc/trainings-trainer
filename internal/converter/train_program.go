package converter

import (
	"github.com/kirillmc/platform_common/pkg/nillable"
	"github.com/kirillmc/trainings-server/internal/model"
	descModer "github.com/kirillmc/trainings-server/pkg/moderator_v1"
	desc "github.com/kirillmc/trainings-server/pkg/training_v1"
)

func ToTrainingProgramFromDesc(trainingProgram *desc.CreateTrainingProgramRequest) *model.TrainProgramToCreate {
	return &model.TrainProgramToCreate{
		ProgramName: trainingProgram.Info.ProgramName,
		UserId:      trainingProgram.Info.UserId,
		Description: trainingProgram.Info.Description,
		Status:      model.Status(trainingProgram.Info.Status),
	}
}

func ToGetTrainingProgramResponseFromService(trainingProgram *model.TrainProgram) *desc.TrainingProgram {
	return &desc.TrainingProgram{
		Id: trainingProgram.Id,
		Info: &desc.TrainingProgramInfo{
			ProgramName: trainingProgram.ProgramName,
			UserId:      trainingProgram.UserId,
			Description: trainingProgram.Description,
			Status:      desc.Status(trainingProgram.Status),
		},
	}
}

func ToGetPublicTrainingProgramResponseFromService(trainingProgram *model.TrainProgram) *desc.PublicTrainingProgram {
	return &desc.PublicTrainingProgram{
		Id: trainingProgram.Id,
		Info: &desc.PublicTrainingProgramInfo{
			ProgramName: trainingProgram.ProgramName,
			AuthorId:    trainingProgram.UserId,
			Description: trainingProgram.Description,
		},
	}
}
func ToGetProgramToModerationResponseFromService(trainingProgram *model.TrainProgram) *descModer.TrainingProgramToModeration {
	return &descModer.TrainingProgramToModeration{
		Id: trainingProgram.Id,
		Info: &descModer.TrainingProgramToModerationInfo{
			ProgramName: trainingProgram.ProgramName,
			AuthorId:    trainingProgram.UserId,
			Description: trainingProgram.Description,
			Status:      descModer.Status(trainingProgram.Status),
		},
	}
}

func ToGetTrainingProgramsResponseFromService(trainingPrograms []*model.TrainProgram) *desc.GetTrainingProgramsResponse {
	var trainingProgramsList []*desc.TrainingProgram
	for _, elem := range trainingPrograms {
		trainingProgramsList = append(trainingProgramsList, ToGetTrainingProgramResponseFromService(elem))
	}
	return &desc.GetTrainingProgramsResponse{TrainPrograms: trainingProgramsList}
}

func ToGetPublicTrainingProgramsResponseFromService(trainingPrograms []*model.TrainProgram) *desc.GetPublicTrainingProgramsResponse {
	var trainingProgramsList []*desc.PublicTrainingProgram
	for _, elem := range trainingPrograms {
		trainingProgramsList = append(trainingProgramsList, ToGetPublicTrainingProgramResponseFromService(elem))
	}
	return &desc.GetPublicTrainingProgramsResponse{TrainPrograms: trainingProgramsList}
}

func ToGetProgramsToModerationResponseFromService(trainingPrograms []*model.TrainProgram) *descModer.GetProgramsToModerationResponse {
	var trainingProgramsList []*descModer.TrainingProgramToModeration
	for _, elem := range trainingPrograms {
		trainingProgramsList = append(trainingProgramsList, ToGetProgramToModerationResponseFromService(elem))
	}
	return &descModer.GetProgramsToModerationResponse{ProgramsToModeration: trainingProgramsList}
}

func ToTrainingProgramUpdateFromDesc(trainingProgram *desc.UpdateTrainingProgramRequest) *model.TrainProgramToUpdate {
	return &model.TrainProgramToUpdate{
		Id:          trainingProgram.GetId(),
		ProgramName: nillable.Create(trainingProgram.Info.ProgramName.Value),
		Description: nillable.Create(trainingProgram.Info.Description.Value),
		Status:      model.Status(trainingProgram.Info.Status),
	}
}
