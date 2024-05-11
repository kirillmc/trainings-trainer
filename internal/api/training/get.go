package training

import (
	"context"

	"github.com/kirillmc/trainings-server/internal/converter"
	desc "github.com/kirillmc/trainings-server/pkg/training_v1"
)

func (i *Implementation) GetTrainingProgram(ctx context.Context, req *desc.GetRequest) (*desc.GetTrainingProgramResponse, error) {
	program, err := i.trainingService.GetProgram(ctx, req.GetId())
	if err != nil {
		return nil, err
	}

	return &desc.GetTrainingProgramResponse{
		TrainProgram: converter.ToGetTrainingProgramResponseFromService(program),
	}, nil

}

func (i *Implementation) GetTrainDay(ctx context.Context, req *desc.GetRequest) (*desc.GetTrainDayResponse, error) {
	trainDay, err := i.trainingService.GetTrainDay(ctx, req.GetId())
	if err != nil {
		return nil, err
	}

	return &desc.GetTrainDayResponse{
		TrainDay: converter.ToGetTrainDayResponseFromService(trainDay),
	}, nil

}

func (i *Implementation) GetExercise(ctx context.Context, req *desc.GetRequest) (*desc.GetExerciseResponse, error) {
	exercise, err := i.trainingService.GetExercise(ctx, req.GetId())
	if err != nil {
		return nil, err
	}

	return &desc.GetExerciseResponse{
		Exercise: converter.ToGetExerciseResponseFromService(exercise),
	}, nil

}

func (i *Implementation) GetSet(ctx context.Context, req *desc.GetRequest) (*desc.GetSetResponse, error) {
	set, err := i.trainingService.GetSet(ctx, req.GetId())
	if err != nil {
		return nil, err
	}

	return &desc.GetSetResponse{
		Set: converter.ToGetSetResponseFromService(set),
	}, nil

}

func (i *Implementation) GetStatistic(ctx context.Context, req *desc.GetRequest) (*desc.GetStatisticResponse, error) {
	statistic, err := i.trainingService.GetStatistic(ctx, req.GetId())
	if err != nil {
		return nil, err
	}

	return &desc.GetStatisticResponse{
		Statistic: converter.ToGetStatisticResponseFromService(statistic),
	}, nil

}
