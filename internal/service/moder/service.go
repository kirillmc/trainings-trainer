package moder

import (
	"github.com/kirillmc/trainings-server/internal/repository"
	def "github.com/kirillmc/trainings-server/internal/service"
)

var _ def.ModerService = (*serv)(nil)

type serv struct {
	moderRepository repository.ModerRepository
}

func NewService(moderRepository repository.ModerRepository) *serv {
	return &serv{moderRepository: moderRepository}
}
