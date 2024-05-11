package training

import (
	"context"

	"github.com/kirillmc/trainings-server/internal/converter"
	desc "github.com/kirillmc/trainings-server/pkg/training_v1"
)

func (i *Implementation) CreateTrainingProgram(ctx context.Context, req *desc.CreateTrainingProgramRequest) (*desc.CreateResponse, error) {
	id, err := i.trainingService.CreateProgram(ctx, converter.ToTrainingProgramFromDesc(req))
	if err != nil {
		return nil, err
	}

	return &desc.CreateResponse{
		Id: id,
	}, nil

}

func (i *Implementation) CreateTrainDay(ctx context.Context, req *desc.CreateTrainDayRequest) (*desc.CreateResponse, error) {
	id, err := i.trainingService.CreateTrainDay(ctx, converter.ToTrainDayFromDesc(req))
	if err != nil {
		return nil, err
	}

	return &desc.CreateResponse{
		Id: id,
	}, nil

}

func (i *Implementation) CreateExercise(ctx context.Context, req *desc.CreateExerciseRequest) (*desc.CreateResponse, error) {
	id, err := i.trainingService.CreateExercise(ctx, converter.ToExerciseFromDesc(req))
	if err != nil {
		return nil, err
	}

	return &desc.CreateResponse{
		Id: id,
	}, nil

}

func (i *Implementation) CreateSet(ctx context.Context, req *desc.CreateSetRequest) (*desc.CreateResponse, error) {
	id, err := i.trainingService.CreateSet(ctx, converter.ToSetFromDesc(req))
	if err != nil {
		return nil, err
	}

	return &desc.CreateResponse{
		Id: id,
	}, nil

}

func (i *Implementation) CreateStatistic(ctx context.Context, req *desc.CreateStatisticRequest) (*desc.CreateResponse, error) {
	id, err := i.trainingService.CreateStatistic(ctx, converter.ToStatisticFromDesc(req))
	if err != nil {
		return nil, err
	}

	return &desc.CreateResponse{
		Id: id,
	}, nil

}
