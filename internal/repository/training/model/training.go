package model

import (
	"time"

	"github.com/kirillmc/trainings-server/internal/model"
)

type TrainProgram struct {
	Id          int64        `db:"id"`
	ProgramName string       `db:"program_name"`
	Description string       `db:"description"`
	Status      model.Status `db:"status"`
}

type TrainDay struct {
	Id          int64  `db:"id"`
	DayName     string `db:"day_name"`
	Description string `db:"description"`
}

type Exercise struct {
	Id           int64  `db:"id"`
	ExerciseName string `db:"exercise_name"`
	Pictures     string `db:"pictures"`
	Description  string `db:"description"`
}

type Set struct {
	Id       int64   `db:"id"`
	Quantity int64   `db:"quantity"`
	Weight   float64 `db:"weight"`
}

type Statistic struct {
	Id         int64     `db:"id"`
	Time       time.Time `db:"time"`
	Quantity   int64     `db:"quantity"`
	Weight     float64   `db:"weight"`
	ProgramId  int64     `db:"set_id"`
	TrainDayId int64     `db:"exercise_id"`
	ExerciseId int64     `db:"trains_day_id"`
	SetId      int64     `db:"program_id"`
}
