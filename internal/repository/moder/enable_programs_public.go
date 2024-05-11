package moder

import (
	"context"

	sq "github.com/Masterminds/squirrel"
	"github.com/kirillmc/platform_common/pkg/db"
)

func (r *repo) EnableProgramsPublic(ctx context.Context, programId int64) error {
	builder := sq.Update(trainProgramsTable).Set(status, 3).
		PlaceholderFormat(sq.Dollar).Where(sq.Eq{idColumm: programId})

	query, args, err := builder.ToSql()
	if err != nil {
		return err
	}

	q := db.Query{
		Name:     "moder_repository.EnableProgramsPublic",
		QueryRaw: query,
	}

	_, err = r.db.DB().ExecContext(ctx, q, args...)
	if err != nil {
		return err
	}

	return nil
}
