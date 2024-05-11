package training

import (
	"context"
)

func (s *serv) DeleteProgram(ctx context.Context, id int64) error {
	err := s.trainingRepository.DeleteProgram(ctx, id)
	if err != nil {
		return err
	}

	return nil
}

func (s *serv) DeleteTrainDay(ctx context.Context, id int64) error {
	err := s.trainingRepository.DeleteTrainDay(ctx, id)
	if err != nil {
		return err
	}

	return nil
}

func (s *serv) DeleteExercise(ctx context.Context, id int64) error {
	err := s.trainingRepository.DeleteExercise(ctx, id)
	if err != nil {
		return err
	}

	return nil
}

func (s *serv) DeleteSet(ctx context.Context, id int64) error {
	err := s.trainingRepository.DeleteSet(ctx, id)
	if err != nil {
		return err
	}

	return nil
}

func (s *serv) DeleteStatistic(ctx context.Context, id int64) error {
	err := s.trainingRepository.DeleteStatistic(ctx, id)
	if err != nil {
		return err
	}

	return nil
}
