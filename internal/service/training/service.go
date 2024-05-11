package training

import (
	"github.com/kirillmc/trainings-server/internal/repository"
	def "github.com/kirillmc/trainings-server/internal/service"
)

var _ def.TrainingService = (*serv)(nil)

type serv struct {
	trainingRepository repository.TrainingRepository
}

func NewService(trainingRepository repository.TrainingRepository) *serv {
	return &serv{trainingRepository: trainingRepository}
}
