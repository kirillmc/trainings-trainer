package training

import (
	"context"

	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/kirillmc/trainings-server/internal/converter"
	desc "github.com/kirillmc/trainings-server/pkg/training_v1"
)

func (i *Implementation) UpdateTrainingProgram(ctx context.Context, req *desc.UpdateTrainingProgramRequest) (*emptypb.Empty, error) {
	err := i.trainingService.UpdateProgram(ctx, converter.ToTrainingProgramUpdateFromDesc(req))
	if err != nil {
		return nil, err
	}

	return nil, nil
}

func (i *Implementation) UpdateTrainDay(ctx context.Context, req *desc.UpdateTrainDayRequest) (*emptypb.Empty, error) {
	err := i.trainingService.UpdateTrainDay(ctx, converter.ToTrainDayUpdateFromDesc(req))
	if err != nil {
		return nil, err
	}

	return nil, nil
}

func (i *Implementation) UpdateExercise(ctx context.Context, req *desc.UpdateExerciseRequest) (*emptypb.Empty, error) {
	err := i.trainingService.UpdateExercise(ctx, converter.ToExerciseUpdateFromDesc(req))
	if err != nil {
		return nil, err
	}

	return nil, nil
}

func (i *Implementation) UpdateSet(ctx context.Context, req *desc.UpdateSetRequest) (*emptypb.Empty, error) {
	err := i.trainingService.UpdateSet(ctx, converter.ToSetUpdateFromDesc(req))
	if err != nil {
		return nil, err
	}

	return nil, nil
}

func (i *Implementation) UpdateStatistic(ctx context.Context, req *desc.UpdateStatisticRequest) (*emptypb.Empty, error) {
	err := i.trainingService.UpdateStatistic(ctx, converter.ToStatitsticUpdateFromDesc(req))
	if err != nil {
		return nil, err
	}

	return nil, nil
}
