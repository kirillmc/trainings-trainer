package moder

import (
	"github.com/kirillmc/platform_common/pkg/db"
	"github.com/kirillmc/trainings-server/internal/repository"
)

const (
	usersProgramsTable         = "users_programs"
	trainersProgramsTable      = "trainers_programs"
	usersTrainersProgramsTable = "users_trainers_programs"
	trainProgramsTable         = "train_programs"
	trainDaysTable             = "train_days"
	daysProgramsTable          = "days_programs"
	exercisesTable             = "exercises"
	exercisesDaysTable         = "exercises_days"
	setsTable                  = "sets"
	setsExercisesTable         = "sets_exercises"
	statisticsTable            = "statistics"
	statisticsSetsTable        = "statistics_sets"

	idColumm          = "id"
	userIdColumn      = "user_id"
	trainerIdColumn   = "trainer_id"
	programIdColumn   = "program_id"
	trainsDayIdColumn = "trains_day_id"
	exerciseIdColumn  = "exercise_id"
	setIdColumn       = "set_id"
	statisticIdColumn = "statistic_id"

	description  = "description"
	programName  = "program_name"
	dayName      = "day_name"
	exerciseName = "exercise_name"
	status       = "status"

	pictures = "pictures"
	quantity = "quantity"
	weight   = "weight"
	time     = "time"

	publicStatusValue = 3
	moderStatusValue  = 2

	returnId = "RETURNING id"
)

type repo struct {
	db db.Client
}

func NewRepository(db db.Client) repository.ModerRepository {
	return &repo{db: db}
}
