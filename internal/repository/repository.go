package repository

import (
	"context"

	"github.com/kirillmc/trainings-server/internal/model"
)

type TrainingRepository interface {
	CreateProgram(ctx context.Context, req *model.TrainProgramToCreate) (int64, error)
	CreateTrainDay(ctx context.Context, req *model.TrainDayToCreate) (int64, error)
	CreateExercise(ctx context.Context, req *model.ExerciseToCreate) (int64, error)
	CreateSet(ctx context.Context, req *model.SetToCreate) (int64, error)
	CreateStatistic(ctx context.Context, req *model.StatisticToCreate) (int64, error)

	CreateUsersTrainersPrograms(ctx context.Context, userId, trainerId, programId int64) error
	CreateUsersPrograms(ctx context.Context, userId, programId int64) error
	CreateTrainersPrograms(ctx context.Context, trainerId, programId int64) error
	CreateDaysPrograms(ctx context.Context, programId, dayId int64) error
	CreateExercisesDays(ctx context.Context, dayId, exerciseId int64) error
	CreateSetsExercises(ctx context.Context, exerciseId, setId int64) error
	CreateStatisticsSets(ctx context.Context, setId, statisticId int64) error

	DeleteProgram(ctx context.Context, id int64) error
	DeleteTrainDay(ctx context.Context, id int64) error
	DeleteExercise(ctx context.Context, id int64) error
	DeleteSet(ctx context.Context, id int64) error
	DeleteStatistic(ctx context.Context, id int64) error

	GetPublicPrograms(ctx context.Context) ([]*model.TrainProgram, error)
	GetProgram(ctx context.Context, id int64) (*model.TrainProgram, error)
	GetTrainDay(ctx context.Context, id int64) (*model.TrainDay, error)
	GetExercise(ctx context.Context, id int64) (*model.Exercise, error)
	GetSet(ctx context.Context, id int64) (*model.Set, error)
	GetStatistic(ctx context.Context, id int64) (*model.Statistic, error)

	GetProgramsIdsByUserId(ctx context.Context, userId int64) ([]int64, error)
	GetTrainDaysIdsByProgramId(ctx context.Context, programId int64) ([]int64, error)
	GetExercisesIdsByTrainDayId(ctx context.Context, trainDayId int64) ([]int64, error)
	GetSetsIdsByExerciseId(ctx context.Context, exerciseId int64) ([]int64, error)

	GetUserIdByProgramId(ctx context.Context, programId int64) (int64, error)
	GetTrainerIdByProgramId(ctx context.Context, programId int64) (int64, error)
	GetProgramIdByTrainDayId(ctx context.Context, trainDayId int64) (int64, error)
	GetTrainDayIdByExerciseId(ctx context.Context, exerciseId int64) (int64, error)
	GetExerciseIdBySetId(ctx context.Context, setId int64) (int64, error)

	UpdateProgram(ctx context.Context, req *model.TrainProgramToUpdate) error
	UpdateTrainDay(ctx context.Context, req *model.TrainDayToUpdate) error
	UpdateExercise(ctx context.Context, req *model.ExerciseToUpdate) error
	UpdateSet(ctx context.Context, req *model.SetToUpdate) error
	UpdateStatistic(ctx context.Context, req *model.StatisticToUpdate) error
}

type ModerRepository interface {
	GetProgramsToModeration(ctx context.Context) ([]*model.TrainProgram, error)
	EnableProgramsPublic(ctx context.Context, programId int64) error
	DisableProgramsPublic(ctx context.Context, programId int64) error

	GetTrainerIdByProgramId(ctx context.Context, programId int64) (int64, error)
	GetProgramsStatus(ctx context.Context, programId int64) (model.Status, error)
}
