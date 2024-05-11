package training

import (
	"context"

	sq "github.com/Masterminds/squirrel"
	"github.com/kirillmc/platform_common/pkg/db"
	"github.com/kirillmc/trainings-server/internal/model"
)

func (r *repo) CreateProgram(ctx context.Context, req *model.TrainProgramToCreate) (int64, error) {
	builder := sq.Insert(trainProgramsTable).PlaceholderFormat(sq.Dollar).
		Columns(programName, description, status).
		Values(req.ProgramName, req.Description, req.Status).
		Suffix(returnId)

	query, args, err := builder.ToSql()
	if err != nil {
		return 0, err
	}

	q := db.Query{
		Name:     "training_repository.CreateTrainingProgram",
		QueryRaw: query,
	}

	var id int64

	err = r.db.DB().QueryRowContext(ctx, q, args...).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (r *repo) CreateTrainDay(ctx context.Context, req *model.TrainDayToCreate) (int64, error) {
	builder := sq.Insert(trainDaysTable).PlaceholderFormat(sq.Dollar).
		Columns(dayName, description).
		Values(req.DayName, req.Description).
		Suffix(returnId)

	query, args, err := builder.ToSql()
	if err != nil {
		return 0, err
	}

	q := db.Query{
		Name:     "training_repository.CreateTrainDay",
		QueryRaw: query,
	}

	var id int64

	err = r.db.DB().QueryRowContext(ctx, q, args...).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil

}

func (r *repo) CreateExercise(ctx context.Context, req *model.ExerciseToCreate) (int64, error) {
	builder := sq.Insert(exercisesTable).PlaceholderFormat(sq.Dollar).
		Columns(exerciseName, description, pictures).
		Values(req.ExerciseName, req.Description, req.Pictures).
		Suffix(returnId)

	query, args, err := builder.ToSql()
	if err != nil {
		return 0, err
	}

	q := db.Query{
		Name:     "training_repository.CreateExercise",
		QueryRaw: query,
	}

	var id int64

	err = r.db.DB().QueryRowContext(ctx, q, args...).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (r *repo) CreateSet(ctx context.Context, req *model.SetToCreate) (int64, error) {
	builder := sq.Insert(setsTable).PlaceholderFormat(sq.Dollar).
		Columns(quantity, weight).
		Values(req.Quantity, req.Weight).
		Suffix(returnId)

	query, args, err := builder.ToSql()
	if err != nil {
		return 0, err
	}

	q := db.Query{
		Name:     "training_repository.CreateSet",
		QueryRaw: query,
	}

	var id int64

	err = r.db.DB().QueryRowContext(ctx, q, args...).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (r *repo) CreateStatistic(ctx context.Context, req *model.StatisticToCreate) (int64, error) {
	builder := sq.Insert(statisticsTable).PlaceholderFormat(sq.Dollar).
		Columns(quantity, weight, programIdColumn, trainsDayIdColumn, exerciseIdColumn, setIdColumn).
		Values(req.Quantity, req.Weight, req.ProgramId, req.TrainDayId, req.ExerciseId, req.SetId).
		Suffix(returnId)

	query, args, err := builder.ToSql()
	if err != nil {
		return 0, err
	}

	q := db.Query{
		Name:     "training_repository.CreateStatistic",
		QueryRaw: query,
	}

	var id int64

	err = r.db.DB().QueryRowContext(ctx, q, args...).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}
