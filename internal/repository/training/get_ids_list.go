package training

import (
	"context"

	sq "github.com/Masterminds/squirrel"
	"github.com/kirillmc/platform_common/pkg/db"
)

func (r *repo) GetProgramsIdsByUserId(ctx context.Context, userId int64) ([]int64, error) {
	builder := sq.Select(programIdColumn).
		PlaceholderFormat(sq.Dollar).
		From(usersProgramsTable).
		Where(sq.Eq{userIdColumn: userId})

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, err
	}

	q := db.Query{
		Name:     "training_repository.GetProgramsIdsByUserId",
		QueryRaw: query,
	}

	var ids []int64

	err = r.db.DB().ScanAllContext(ctx, &ids, q, args...)
	if err != nil {
		return nil, err
	}

	return ids, nil
}

func (r *repo) GetTrainDaysIdsByProgramId(ctx context.Context, programId int64) ([]int64, error) {
	builder := sq.Select(trainsDayIdColumn).
		PlaceholderFormat(sq.Dollar).
		From(daysProgramsTable).
		Where(sq.Eq{programIdColumn: programId})

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, err
	}

	q := db.Query{
		Name:     "training_repository.GetTrainDaysIdsByProgramId",
		QueryRaw: query,
	}

	var ids []int64

	err = r.db.DB().ScanAllContext(ctx, &ids, q, args...)
	if err != nil {
		return nil, err
	}

	return ids, nil
}

func (r *repo) GetExercisesIdsByTrainDayId(ctx context.Context, trainDayId int64) ([]int64, error) {
	builder := sq.Select(exerciseIdColumn).
		PlaceholderFormat(sq.Dollar).
		From(exercisesDaysTable).
		Where(sq.Eq{trainsDayIdColumn: trainDayId})

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, err
	}

	q := db.Query{
		Name:     "training_repository.GetExercisesIdsByTrainDayId",
		QueryRaw: query,
	}

	var ids []int64

	err = r.db.DB().ScanAllContext(ctx, &ids, q, args...)
	if err != nil {
		return nil, err
	}

	return ids, nil
}

func (r *repo) GetSetsIdsByExerciseId(ctx context.Context, exerciseId int64) ([]int64, error) {
	builder := sq.Select(setIdColumn).
		PlaceholderFormat(sq.Dollar).
		From(setsExercisesTable).
		Where(sq.Eq{exerciseIdColumn: exerciseId})

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, err
	}

	q := db.Query{
		Name:     "training_repository.GetSetsIdsByExerciseId",
		QueryRaw: query,
	}

	var ids []int64

	err = r.db.DB().ScanAllContext(ctx, &ids, q, args...)
	if err != nil {
		return nil, err
	}

	return ids, nil
}
