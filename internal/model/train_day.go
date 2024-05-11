package model

import (
	"github.com/kirillmc/platform_common/pkg/nillable"
)

type TrainDay struct {
	Id          int64
	DayName     string
	Description string
	ProgramId   int64
}

type TrainDayToCreate struct {
	DayName     string
	Description string
	ProgramId   int64
}

type TrainDayToUpdate struct {
	Id          int64
	DayName     nillable.NilString
	Description nillable.NilString
}
