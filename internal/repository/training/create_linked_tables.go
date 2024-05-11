package training

import (
	"context"

	sq "github.com/Masterminds/squirrel"
	"github.com/kirillmc/platform_common/pkg/db"
)

func (r *repo) CreateUsersTrainersPrograms(ctx context.Context, userId, trainerId, programId int64) error {
	builder := sq.Insert(usersTrainersProgramsTable).
		PlaceholderFormat(sq.Dollar).
		Columns(userIdColumn, trainerIdColumn, programIdColumn).
		Values(userId, trainerId, programId).
		Suffix(returnId)

	query, args, err := builder.ToSql()
	if err != nil {
		return err
	}

	q := db.Query{
		Name:     "training_repository.CreateUsersTrainersPrograms",
		QueryRaw: query,
	}

	var id int64

	err = r.db.DB().QueryRowContext(ctx, q, args...).Scan(&id)
	if err != nil {
		return err
	}

	return nil
}
func (r *repo) CreateUsersPrograms(ctx context.Context, userId, programId int64) error {
	builder := sq.Insert(usersProgramsTable).PlaceholderFormat(sq.Dollar).
		Columns(userIdColumn, programIdColumn).
		Values(userId, programId).
		Suffix(returnId)

	query, args, err := builder.ToSql()
	if err != nil {
		return err
	}

	q := db.Query{
		Name:     "training_repository.CreateUsersPrograms",
		QueryRaw: query,
	}

	var id int64

	err = r.db.DB().QueryRowContext(ctx, q, args...).Scan(&id)
	if err != nil {
		return err
	}

	return nil
}
func (r *repo) CreateTrainersPrograms(ctx context.Context, trainerId, programId int64) error {
	builder := sq.Insert(trainersProgramsTable).PlaceholderFormat(sq.Dollar).
		Columns(trainerIdColumn, programIdColumn).
		Values(trainerId, programId).
		Suffix(returnId)

	query, args, err := builder.ToSql()
	if err != nil {
		return err
	}

	q := db.Query{
		Name:     "training_repository.CreateTrainersPrograms",
		QueryRaw: query,
	}

	var id int64

	err = r.db.DB().QueryRowContext(ctx, q, args...).Scan(&id)
	if err != nil {
		return err
	}

	return nil
}

func (r *repo) CreateDaysPrograms(ctx context.Context, programId, dayId int64) error {
	builder := sq.Insert(daysProgramsTable).PlaceholderFormat(sq.Dollar).
		Columns(programIdColumn, trainsDayIdColumn).
		Values(programId, dayId).
		Suffix(returnId)

	query, args, err := builder.ToSql()
	if err != nil {
		return err
	}

	q := db.Query{
		Name:     "training_repository.CreateDaysPrograms",
		QueryRaw: query,
	}

	var id int64

	err = r.db.DB().QueryRowContext(ctx, q, args...).Scan(&id)
	if err != nil {
		return err
	}

	return nil
}

func (r *repo) CreateExercisesDays(ctx context.Context, dayId, exerciseId int64) error {
	builder := sq.Insert(exercisesDaysTable).PlaceholderFormat(sq.Dollar).
		Columns(trainsDayIdColumn, exerciseIdColumn).
		Values(dayId, exerciseId).
		Suffix(returnId)

	query, args, err := builder.ToSql()
	if err != nil {
		return err
	}

	q := db.Query{
		Name:     "training_repository.CreateExercisesDays",
		QueryRaw: query,
	}

	var id int64

	err = r.db.DB().QueryRowContext(ctx, q, args...).Scan(&id)
	if err != nil {
		return err
	}

	return nil
}

func (r *repo) CreateSetsExercises(ctx context.Context, exerciseId, setId int64) error {
	builder := sq.Insert(setsExercisesTable).PlaceholderFormat(sq.Dollar).
		Columns(exerciseIdColumn, setIdColumn).
		Values(exerciseId, setId).
		Suffix(returnId)

	query, args, err := builder.ToSql()
	if err != nil {
		return err
	}

	q := db.Query{
		Name:     "training_repository.CreateSetsExercises",
		QueryRaw: query,
	}

	var id int64

	err = r.db.DB().QueryRowContext(ctx, q, args...).Scan(&id)
	if err != nil {
		return err
	}

	return nil
}

func (r *repo) CreateStatisticsSets(ctx context.Context, setId, statisticId int64) error {
	builder := sq.Insert(statisticsSetsTable).PlaceholderFormat(sq.Dollar).
		Columns(setIdColumn, statisticIdColumn).
		Values(setId, statisticId).
		Suffix(returnId)

	query, args, err := builder.ToSql()
	if err != nil {
		return err
	}

	q := db.Query{
		Name:     "training_repository.CreateStatisticsSets",
		QueryRaw: query,
	}

	var id int64

	err = r.db.DB().QueryRowContext(ctx, q, args...).Scan(&id)
	if err != nil {
		return err
	}

	return nil
}
