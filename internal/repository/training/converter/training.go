package converter

import (
	"github.com/kirillmc/trainings-server/internal/model"
	modelRepo "github.com/kirillmc/trainings-server/internal/repository/training/model"
)

func ToTrainingProgramFromRepo(trainingProgram *modelRepo.TrainProgram) *model.TrainProgram {
	return &model.TrainProgram{
		Id:          trainingProgram.Id,
		ProgramName: trainingProgram.ProgramName,
		Description: trainingProgram.Description,
		Status:      trainingProgram.Status,
	}
}
func ToTrainingProgramsFromRepo(trainingPrograms []*modelRepo.TrainProgram) []*model.TrainProgram {
	var programs []*model.TrainProgram
	for _, elem := range trainingPrograms {
		programs = append(programs, ToTrainingProgramFromRepo(elem))
	}

	return programs
}

func ToTrainDayFromRepo(trainDay *modelRepo.TrainDay) *model.TrainDay {
	return &model.TrainDay{
		Id:          trainDay.Id,
		DayName:     trainDay.DayName,
		Description: trainDay.Description,
	}
}

func ToExerciseFromRepo(exercise *modelRepo.Exercise) *model.Exercise {
	return &model.Exercise{
		Id:           exercise.Id,
		ExerciseName: exercise.ExerciseName,
		Pictures:     exercise.Pictures,
		Description:  exercise.Description,
	}
}

func ToSetFromRepo(set *modelRepo.Set) *model.Set {
	return &model.Set{
		Id:       set.Id,
		Quantity: set.Quantity,
		Weight:   set.Weight,
	}
}

func ToStatisticFromRepo(set *modelRepo.Statistic) *model.Statistic {
	return &model.Statistic{
		Id:         set.Id,
		Quantity:   set.Quantity,
		Weight:     set.Weight,
		ProgramId:  set.ProgramId,
		TrainDayId: set.TrainDayId,
		ExerciseId: set.ExerciseId,
		SetId:      set.SetId,
		Time:       set.Time,
	}
}
