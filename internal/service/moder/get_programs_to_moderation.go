package moder

import (
	"context"

	"github.com/kirillmc/trainings-server/internal/model"
)

func (s *serv) GetProgramsToModeration(ctx context.Context) ([]*model.TrainProgram, error) {
	var programs []*model.TrainProgram

	programs, err := s.moderRepository.GetProgramsToModeration(ctx)
	if err != nil {
		return nil, err
	}

	for i := range programs {
		temp, err := s.moderRepository.GetTrainerIdByProgramId(ctx, programs[i].Id)
		if err != nil {
			return nil, err
		}

		programs[i].UserId = temp
	}

	return programs, nil
}
