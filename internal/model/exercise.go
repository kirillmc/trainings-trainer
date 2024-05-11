package model

import (
	"github.com/kirillmc/platform_common/pkg/nillable"
)

type Exercise struct {
	Id           int64
	ExerciseName string
	Pictures     string
	Description  string
	DayId        int64
}

type ExerciseToCreate struct {
	ExerciseName string
	Pictures     string
	Description  string
	DayId        int64
}

type ExerciseToUpdate struct {
	Id           int64
	Pictures     nillable.NilString
	ExerciseName nillable.NilString
	Description  nillable.NilString
}
