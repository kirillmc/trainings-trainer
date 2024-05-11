package training

import (
	"context"

	"google.golang.org/protobuf/types/known/emptypb"

	desc "github.com/kirillmc/trainings-server/pkg/training_v1"
)

func (i *Implementation) DeleteTrainingProgram(ctx context.Context, req *desc.DeleteRequest) (*emptypb.Empty, error) {
	err := i.trainingService.DeleteProgram(ctx, req.GetId())
	if err != nil {
		return nil, err
	}

	return nil, nil
}

func (i *Implementation) DeleteTrainDay(ctx context.Context, req *desc.DeleteRequest) (*emptypb.Empty, error) {
	err := i.trainingService.DeleteTrainDay(ctx, req.GetId())
	if err != nil {
		return nil, err
	}

	return nil, nil
}

func (i *Implementation) DeleteExercise(ctx context.Context, req *desc.DeleteRequest) (*emptypb.Empty, error) {
	err := i.trainingService.DeleteExercise(ctx, req.GetId())
	if err != nil {
		return nil, err
	}

	return nil, nil
}

func (i *Implementation) DeleteSet(ctx context.Context, req *desc.DeleteRequest) (*emptypb.Empty, error) {
	err := i.trainingService.DeleteSet(ctx, req.GetId())
	if err != nil {
		return nil, err
	}

	return nil, nil
}

func (i *Implementation) DeleteStatistic(ctx context.Context, req *desc.DeleteRequest) (*emptypb.Empty, error) {
	err := i.trainingService.DeleteStatistic(ctx, req.GetId())
	if err != nil {
		return nil, err
	}

	return nil, nil
}
