package training

import (
	"context"

	"github.com/kirillmc/trainings-server/internal/model"
)

func (s *serv) GetPrograms(ctx context.Context, userId int64) ([]*model.TrainProgram, error) {
	programsIds, err := s.trainingRepository.GetProgramsIdsByUserId(ctx, userId)
	if err != nil {
		return nil, err
	}

	var programs []*model.TrainProgram

	for _, program_id := range programsIds {
		program, errInternal := s.trainingRepository.GetProgram(ctx, program_id)
		if errInternal != nil {
			return nil, err
		}
		programs = append(programs, program)
	}

	return programs, nil
}

func (s *serv) GetPublicPrograms(ctx context.Context) ([]*model.TrainProgram, error) {
	var programs []*model.TrainProgram

	programs, err := s.trainingRepository.GetPublicPrograms(ctx)
	if err != nil {
		return nil, err
	}

	for i := range programs {
		temp, err := s.trainingRepository.GetTrainerIdByProgramId(ctx, programs[i].Id)
		if err != nil {
			return nil, err
		}

		programs[i].UserId = temp
	}

	return programs, nil
}

func (s *serv) GetTrainDays(ctx context.Context, programId int64) ([]*model.TrainDay, error) {
	tainDaysIds, err := s.trainingRepository.GetTrainDaysIdsByProgramId(ctx, programId)
	if err != nil {
		return nil, err
	}

	var trainDays []*model.TrainDay

	for _, trainDayId := range tainDaysIds {
		trainDay, errInternal := s.trainingRepository.GetTrainDay(ctx, trainDayId)
		if errInternal != nil {
			return nil, err
		}
		trainDays = append(trainDays, trainDay)
	}

	return trainDays, nil
}

func (s *serv) GetExercises(ctx context.Context, trainDayId int64) ([]*model.Exercise, error) {
	exercisesIds, err := s.trainingRepository.GetExercisesIdsByTrainDayId(ctx, trainDayId)
	if err != nil {
		return nil, err
	}

	var exercises []*model.Exercise

	for _, exerciseId := range exercisesIds {
		exercise, errInternal := s.trainingRepository.GetExercise(ctx, exerciseId)
		if errInternal != nil {
			return nil, err
		}
		exercises = append(exercises, exercise)
	}

	return exercises, nil
}

func (s *serv) GetSets(ctx context.Context, exerciseId int64) ([]*model.Set, error) {
	setsIds, err := s.trainingRepository.GetSetsIdsByExerciseId(ctx, exerciseId)
	if err != nil {
		return nil, err
	}

	var sets []*model.Set

	for _, setId := range setsIds {
		set, errInternal := s.trainingRepository.GetSet(ctx, setId)
		if errInternal != nil {
			return nil, err
		}
		sets = append(sets, set)
	}

	return sets, nil
}
