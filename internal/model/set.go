package model

import (
	"github.com/kirillmc/platform_common/pkg/nillable"
)

type Set struct {
	Id         int64
	Quantity   int64
	Weight     float64
	ExerciseId int64
}

type SetToCreate struct {
	Quantity   int64
	Weight     float64
	ExerciseId int64
}

type SetToUpdate struct {
	Id       int64
	Quantity nillable.NilInt
	Weight   nillable.NilDouble
}
