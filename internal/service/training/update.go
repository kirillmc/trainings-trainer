package training

import (
	"context"

	"github.com/kirillmc/trainings-server/internal/model"
)

func (s *serv) UpdateProgram(ctx context.Context, req *model.TrainProgramToUpdate) error {
	err := s.trainingRepository.UpdateProgram(ctx, req)
	if err != nil {
		return err
	}

	return nil
}

func (s *serv) UpdateTrainDay(ctx context.Context, req *model.TrainDayToUpdate) error {
	err := s.trainingRepository.UpdateTrainDay(ctx, req)
	if err != nil {
		return err
	}

	return nil
}

func (s *serv) UpdateExercise(ctx context.Context, req *model.ExerciseToUpdate) error {
	err := s.trainingRepository.UpdateExercise(ctx, req)
	if err != nil {
		return err
	}

	return nil
}

func (s *serv) UpdateSet(ctx context.Context, req *model.SetToUpdate) error {
	err := s.trainingRepository.UpdateSet(ctx, req)
	if err != nil {
		return err
	}

	return nil
}

func (s *serv) UpdateStatistic(ctx context.Context, req *model.StatisticToUpdate) error {
	err := s.trainingRepository.UpdateStatistic(ctx, req)
	if err != nil {
		return err
	}

	return nil
}
