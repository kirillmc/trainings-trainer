package moderator

import (
	"context"

	"google.golang.org/protobuf/types/known/emptypb"

	desc "github.com/kirillmc/trainings-server/pkg/moderator_v1"
)

func (i *Implementation) EnableProgramsPublic(ctx context.Context, req *desc.EnableProgramsPublicRequest) (*emptypb.Empty, error) {
	err := i.moderatorService.EnableProgramsPublic(ctx, req.GetProgramId())
	if err != nil {
		return nil, err
	}

	return nil, nil
}
