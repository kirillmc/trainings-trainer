package moderator

import (
	"context"

	"google.golang.org/protobuf/types/known/emptypb"

	desc "github.com/kirillmc/trainings-server/pkg/moderator_v1"
)

func (i *Implementation) DisableProgramsPublic(ctx context.Context, req *desc.DisableProgramsPublicRequest) (*emptypb.Empty, error) {
	err := i.moderatorService.DisableProgramsPublic(ctx, req.GetProgramId())
	if err != nil {
		return nil, err
	}

	return nil, nil
}
