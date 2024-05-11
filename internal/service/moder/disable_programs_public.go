package moder

import (
	"context"
	"errors"

	"github.com/kirillmc/trainings-server/internal/model"
)

func (s *serv) DisableProgramsPublic(ctx context.Context, programId int64) error {
	status, err := s.moderRepository.GetProgramsStatus(ctx, programId)
	if err != nil {
		return err
	}

	if status != model.Modering {
		return errors.New("program isn't in modering status")
	}

	err = s.moderRepository.DisableProgramsPublic(ctx, programId)
	if err != nil {
		return err
	}

	return nil
}
