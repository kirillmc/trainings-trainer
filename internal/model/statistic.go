package model

import (
	"time"

	"github.com/kirillmc/platform_common/pkg/nillable"
)

type Statistic struct {
	Id         int64
	Quantity   int64
	Weight     float64
	ProgramId  int64
	TrainDayId int64
	ExerciseId int64
	SetId      int64
	Time       time.Time
}

type StatisticToCreate struct {
	Quantity   int64
	Weight     float64
	ProgramId  int64
	TrainDayId int64
	ExerciseId int64
	SetId      int64
}

type StatisticToUpdate struct {
	Id       int64
	Quantity nillable.NilInt
	Weight   nillable.NilDouble
}
