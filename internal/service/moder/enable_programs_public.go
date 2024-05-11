package moder

import (
	"context"

	"github.com/kirillmc/trainings-server/internal/model"
	"github.com/pkg/errors"
)

func (s *serv) EnableProgramsPublic(ctx context.Context, programId int64) error {
	status, err := s.moderRepository.GetProgramsStatus(ctx, programId)
	if err != nil {
		return err
	}

	if status != model.Modering {
		return errors.New("program isn't in modering status")
	}

	err = s.moderRepository.EnableProgramsPublic(ctx, programId)
	if err != nil {
		return err
	}

	return nil
}
