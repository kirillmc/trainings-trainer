package training

import (
	"context"

	"github.com/kirillmc/trainings-server/internal/model"
)

func (s *serv) CreateProgram(ctx context.Context, req *model.TrainProgramToCreate) (int64, error) {
	id, err := s.trainingRepository.CreateProgram(ctx, req)
	if err != nil {
		return 0, err
	}

	err = s.trainingRepository.CreateUsersPrograms(ctx, req.UserId, id)
	if err != nil {
		return 0, err
	}

	err = s.trainingRepository.CreateTrainersPrograms(ctx, req.UserId, id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (s *serv) CreateTrainDay(ctx context.Context, req *model.TrainDayToCreate) (int64, error) {
	id, err := s.trainingRepository.CreateTrainDay(ctx, req)
	if err != nil {
		return 0, err
	}

	err = s.trainingRepository.CreateDaysPrograms(ctx, req.ProgramId, id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (s *serv) CreateExercise(ctx context.Context, req *model.ExerciseToCreate) (int64, error) {
	id, err := s.trainingRepository.CreateExercise(ctx, req)
	if err != nil {
		return 0, err
	}

	err = s.trainingRepository.CreateExercisesDays(ctx, req.DayId, id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (s *serv) CreateSet(ctx context.Context, req *model.SetToCreate) (int64, error) {
	id, err := s.trainingRepository.CreateSet(ctx, req)
	if err != nil {
		return 0, err
	}

	err = s.trainingRepository.CreateSetsExercises(ctx, req.ExerciseId, id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (s *serv) CreateStatistic(ctx context.Context, req *model.StatisticToCreate) (int64, error) {
	id, err := s.trainingRepository.CreateStatistic(ctx, req)
	if err != nil {
		return 0, err
	}

	err = s.trainingRepository.CreateStatisticsSets(ctx, req.SetId, id)
	if err != nil {
		return 0, err
	}

	return id, nil
}
