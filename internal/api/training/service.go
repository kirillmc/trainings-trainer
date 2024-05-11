package training

import (
	"github.com/kirillmc/trainings-server/internal/service"
	desc "github.com/kirillmc/trainings-server/pkg/training_v1"
)

type Implementation struct {
	desc.UnimplementedTrainingV1Server
	trainingService service.TrainingService
}

func NewImplementation(trainigService service.TrainingService) *Implementation {
	return &Implementation{
		trainingService: trainigService,
	}
}
