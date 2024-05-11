package training

import (
	"context"

	sq "github.com/Masterminds/squirrel"
	"github.com/kirillmc/platform_common/pkg/db"
)

func (r *repo) GetUserIdByProgramId(ctx context.Context, programId int64) (int64, error) {
	builder := sq.Select(userIdColumn).
		PlaceholderFormat(sq.Dollar).
		From(usersProgramsTable).
		Where(sq.Eq{programIdColumn: programId}).
		Limit(1)

	query, args, err := builder.ToSql()
	if err != nil {
		return 0, err
	}

	q := db.Query{
		Name:     "training_repository.GetUserIdByProgramId",
		QueryRaw: query,
	}

	var id int64

	err = r.db.DB().ScanOneContext(ctx, &id, q, args...)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (r *repo) GetTrainerIdByProgramId(ctx context.Context, programId int64) (int64, error) {
	builder := sq.Select(trainerIdColumn).
		PlaceholderFormat(sq.Dollar).
		From(trainersProgramsTable).
		Where(sq.Eq{programIdColumn: programId}).
		Limit(1)

	query, args, err := builder.ToSql()
	if err != nil {
		return 0, err
	}

	q := db.Query{
		Name:     "training_repository.GetTrainerIdByProgramId",
		QueryRaw: query,
	}

	var id int64

	err = r.db.DB().ScanOneContext(ctx, &id, q, args...)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (r *repo) GetProgramIdByTrainDayId(ctx context.Context, trainDayId int64) (int64, error) {
	builder := sq.Select(programIdColumn).
		PlaceholderFormat(sq.Dollar).
		From(daysProgramsTable).
		Where(sq.Eq{trainsDayIdColumn: trainDayId}).
		Limit(1)

	query, args, err := builder.ToSql()
	if err != nil {
		return 0, err
	}

	q := db.Query{
		Name:     "training_repository.GetProgramIdByTrainDayId",
		QueryRaw: query,
	}

	var id int64

	err = r.db.DB().ScanOneContext(ctx, &id, q, args...)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (r *repo) GetTrainDayIdByExerciseId(ctx context.Context, exerciseId int64) (int64, error) {
	builder := sq.Select(trainsDayIdColumn).
		PlaceholderFormat(sq.Dollar).
		From(exercisesDaysTable).
		Where(sq.Eq{exerciseIdColumn: exerciseId}).
		Limit(1)

	query, args, err := builder.ToSql()
	if err != nil {
		return 0, err
	}

	q := db.Query{
		Name:     "training_repository.GetTrainDayIdByExerciseId",
		QueryRaw: query,
	}

	var id int64

	err = r.db.DB().ScanOneContext(ctx, &id, q, args...)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (r *repo) GetExerciseIdBySetId(ctx context.Context, setId int64) (int64, error) {
	builder := sq.Select(exerciseIdColumn).
		PlaceholderFormat(sq.Dollar).
		From(setsExercisesTable).
		Where(sq.Eq{setIdColumn: setId}).
		Limit(1)

	query, args, err := builder.ToSql()
	if err != nil {
		return 0, err
	}

	q := db.Query{
		Name:     "training_repository.GetExerciseIdBySetId",
		QueryRaw: query,
	}

	var id int64

	err = r.db.DB().ScanOneContext(ctx, &id, q, args...)
	if err != nil {
		return 0, err
	}

	return id, nil
}
