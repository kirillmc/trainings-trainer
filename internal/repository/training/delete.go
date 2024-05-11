package training

import (
	"context"

	sq "github.com/Masterminds/squirrel"
	"github.com/kirillmc/platform_common/pkg/db"
)

func (r *repo) DeleteProgram(ctx context.Context, id int64) error {
	builder := sq.Delete(trainProgramsTable).PlaceholderFormat(sq.Dollar).Where(sq.Eq{idColumm: id})

	query, args, err := builder.ToSql()
	if err != nil {
		return nil
	}

	q := db.Query{
		Name:     "training_repository.DeleteProgram",
		QueryRaw: query,
	}

	_, err = r.db.DB().ExecContext(ctx, q, args...)
	if err != nil {
		return nil
	}

	return nil
}

func (r *repo) DeleteTrainDay(ctx context.Context, id int64) error {
	builder := sq.Delete(trainDaysTable).PlaceholderFormat(sq.Dollar).Where(sq.Eq{idColumm: id})

	query, args, err := builder.ToSql()
	if err != nil {
		return nil
	}

	q := db.Query{
		Name:     "training_repository.DeleteTrainDay",
		QueryRaw: query,
	}

	_, err = r.db.DB().ExecContext(ctx, q, args...)
	if err != nil {
		return nil
	}

	return nil
}

func (r *repo) DeleteExercise(ctx context.Context, id int64) error {
	builder := sq.Delete(exercisesTable).PlaceholderFormat(sq.Dollar).Where(sq.Eq{idColumm: id})

	query, args, err := builder.ToSql()
	if err != nil {
		return nil
	}

	q := db.Query{
		Name:     "training_repository.DeleteExercise",
		QueryRaw: query,
	}

	_, err = r.db.DB().ExecContext(ctx, q, args...)
	if err != nil {
		return nil
	}

	return nil
}

func (r *repo) DeleteSet(ctx context.Context, id int64) error {
	builder := sq.Delete(setsTable).PlaceholderFormat(sq.Dollar).Where(sq.Eq{idColumm: id})

	query, args, err := builder.ToSql()
	if err != nil {
		return nil
	}

	q := db.Query{
		Name:     "training_repository.DeleteSet",
		QueryRaw: query,
	}

	_, err = r.db.DB().ExecContext(ctx, q, args...)
	if err != nil {
		return nil
	}

	return nil
}

func (r *repo) DeleteStatistic(ctx context.Context, id int64) error {
	builder := sq.Delete(statisticsTable).PlaceholderFormat(sq.Dollar).Where(sq.Eq{idColumm: id})

	query, args, err := builder.ToSql()
	if err != nil {
		return nil
	}

	q := db.Query{
		Name:     "training_repository.DeleteStatistic",
		QueryRaw: query,
	}

	_, err = r.db.DB().ExecContext(ctx, q, args...)
	if err != nil {
		return nil
	}

	return nil
}
