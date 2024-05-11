package training

import (
	"context"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/kirillmc/trainings-server/internal/converter"
	desc "github.com/kirillmc/trainings-server/pkg/training_v1"
)

func (i *Implementation) GetPublicTrainingPrograms(ctx context.Context, _ *empty.Empty) (*desc.GetPublicTrainingProgramsResponse, error) {
	programs, err := i.trainingService.GetPublicPrograms(ctx)
	if err != nil {
		return nil, err
	}

	return converter.ToGetPublicTrainingProgramsResponseFromService(programs), nil

}

func (i *Implementation) GetTrainingPrograms(ctx context.Context, req *desc.GetRequest) (*desc.GetTrainingProgramsResponse, error) {
	programs, err := i.trainingService.GetPrograms(ctx, req.GetId())
	if err != nil {
		return nil, err
	}

	return converter.ToGetTrainingProgramsResponseFromService(programs), nil

}

func (i *Implementation) GetTrainDays(ctx context.Context, req *desc.GetRequest) (*desc.GetTrainDaysResponse, error) {
	traindDays, err := i.trainingService.GetTrainDays(ctx, req.GetId())
	if err != nil {
		return nil, err
	}

	return converter.ToGetTrainDaysResponseFromService(traindDays), nil

}

func (i *Implementation) GetExercises(ctx context.Context, req *desc.GetRequest) (*desc.GetExercisesResponse, error) {
	exercises, err := i.trainingService.GetExercises(ctx, req.GetId())
	if err != nil {
		return nil, err
	}

	return converter.ToGetExercisesResponseFromService(exercises), nil

}

func (i *Implementation) GetSets(ctx context.Context, req *desc.GetRequest) (*desc.GetSetsResponse, error) {
	sets, err := i.trainingService.GetSets(ctx, req.GetId())
	if err != nil {
		return nil, err
	}

	return converter.ToGetSetsResponseFromService(sets), nil

}
