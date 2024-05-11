package moderator

import (
	"github.com/kirillmc/trainings-server/internal/service"
	desc "github.com/kirillmc/trainings-server/pkg/moderator_v1"
)

type Implementation struct {
	desc.UnimplementedModeratorV1Server
	moderatorService service.ModerService
}

func NewImplementation(moderatorService service.ModerService) *Implementation {
	return &Implementation{
		moderatorService: moderatorService,
	}
}
