package service

import (
	"context"

	"github.com/kirillmc/trainings-server/internal/model"
)

type TrainingService interface {
	CreateProgram(ctx context.Context, req *model.TrainProgramToCreate) (int64, error)
	CreateTrainDay(ctx context.Context, req *model.TrainDayToCreate) (int64, error)
	CreateExercise(ctx context.Context, req *model.ExerciseToCreate) (int64, error)
	CreateSet(ctx context.Context, req *model.SetToCreate) (int64, error)
	CreateStatistic(ctx context.Context, req *model.StatisticToCreate) (int64, error)

	DeleteProgram(ctx context.Context, id int64) error
	DeleteTrainDay(ctx context.Context, id int64) error
	DeleteExercise(ctx context.Context, id int64) error
	DeleteSet(ctx context.Context, id int64) error
	DeleteStatistic(ctx context.Context, id int64) error

	GetProgram(ctx context.Context, id int64) (*model.TrainProgram, error)
	GetTrainDay(ctx context.Context, id int64) (*model.TrainDay, error)
	GetExercise(ctx context.Context, id int64) (*model.Exercise, error)
	GetSet(ctx context.Context, id int64) (*model.Set, error)
	GetStatistic(ctx context.Context, id int64) (*model.Statistic, error)

	GetPrograms(ctx context.Context, userId int64) ([]*model.TrainProgram, error)
	GetPublicPrograms(ctx context.Context) ([]*model.TrainProgram, error)
	GetTrainDays(ctx context.Context, programId int64) ([]*model.TrainDay, error)
	GetExercises(ctx context.Context, trainDayId int64) ([]*model.Exercise, error)
	GetSets(ctx context.Context, exerciseId int64) ([]*model.Set, error)

	UpdateProgram(ctx context.Context, req *model.TrainProgramToUpdate) error
	UpdateTrainDay(ctx context.Context, req *model.TrainDayToUpdate) error
	UpdateExercise(ctx context.Context, req *model.ExerciseToUpdate) error
	UpdateSet(ctx context.Context, req *model.SetToUpdate) error
	UpdateStatistic(ctx context.Context, req *model.StatisticToUpdate) error
}

type ModerService interface {
	GetProgramsToModeration(ctx context.Context) ([]*model.TrainProgram, error)
	EnableProgramsPublic(ctx context.Context, programId int64) error
	DisableProgramsPublic(ctx context.Context, programId int64) error
}
