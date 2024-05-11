package converter

import (
	"github.com/kirillmc/platform_common/pkg/nillable"
	"github.com/kirillmc/trainings-server/internal/model"
	desc "github.com/kirillmc/trainings-server/pkg/training_v1"
)

func ToExerciseFromDesc(exercise *desc.CreateExerciseRequest) *model.ExerciseToCreate {
	return &model.ExerciseToCreate{
		ExerciseName: exercise.Info.ExerciseName,
		Pictures:     exercise.Info.Pictures,
		Description:  exercise.Info.Description,
		DayId:        exercise.Info.DayId,
	}
}

func ToGetExerciseResponseFromService(exercise *model.Exercise) *desc.Exercise {
	return &desc.Exercise{
		Id: exercise.Id,
		Info: &desc.ExerciseInfo{
			ExerciseName: exercise.ExerciseName,
			Pictures:     exercise.Pictures,
			Description:  exercise.Description,
			DayId:        exercise.DayId,
		},
	}
}

func ToGetExercisesResponseFromService(exercises []*model.Exercise) *desc.GetExercisesResponse {
	var exerciseList []*desc.Exercise
	for _, elem := range exercises {
		exerciseList = append(exerciseList, ToGetExerciseResponseFromService(elem))
	}
	return &desc.GetExercisesResponse{Exercises: exerciseList}
}

func ToExerciseUpdateFromDesc(exercise *desc.UpdateExerciseRequest) *model.ExerciseToUpdate {
	return &model.ExerciseToUpdate{
		Id:           exercise.GetId(),
		ExerciseName: nillable.Create(exercise.Info.ExerciseName.Value),
		Pictures:     nillable.Create(exercise.Info.Pictures.Value),
		Description:  nillable.Create(exercise.Info.Description.Value),
	}
}
