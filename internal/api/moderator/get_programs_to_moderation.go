package moderator

import (
	"context"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/kirillmc/trainings-server/internal/converter"
	desc "github.com/kirillmc/trainings-server/pkg/moderator_v1"
)

func (i *Implementation) GetProgramsToModeration(ctx context.Context, _ *empty.Empty) (*desc.GetProgramsToModerationResponse, error) {
	programs, err := i.moderatorService.GetProgramsToModeration(ctx)
	if err != nil {
		return nil, err
	}

	return converter.ToGetProgramsToModerationResponseFromService(programs), nil

}
