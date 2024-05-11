package moder

import (
	"context"

	sq "github.com/Masterminds/squirrel"
	"github.com/kirillmc/platform_common/pkg/db"
	"github.com/kirillmc/trainings-server/internal/model"
	"github.com/kirillmc/trainings-server/internal/repository/training/converter"
	modelRepo "github.com/kirillmc/trainings-server/internal/repository/training/model"
)

func (r *repo) GetProgramsToModeration(ctx context.Context) ([]*model.TrainProgram, error) {
	builder := sq.Select(idColumm, programName, description, status).
		PlaceholderFormat(sq.Dollar).
		From(trainProgramsTable).
		Where(sq.Eq{status: moderStatusValue})

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, err
	}

	q := db.Query{
		Name:     "moder_repository.GetProgramsToModeration",
		QueryRaw: query,
	}

	var programs []*modelRepo.TrainProgram

	err = r.db.DB().ScanAllContext(ctx, &programs, q, args...)
	if err != nil {
		return nil, err
	}

	return converter.ToTrainingProgramsFromRepo(programs), nil
}

func (r *repo) GetProgramsStatus(ctx context.Context, programId int64) (model.Status, error) {
	builder := sq.Select(status).
		PlaceholderFormat(sq.Dollar).
		From(trainProgramsTable).
		Where(sq.Eq{idColumm: programId})

	query, args, err := builder.ToSql()
	if err != nil {
		return model.Unknown, err
	}

	q := db.Query{
		Name:     "moder_repository.GetProgramsStatus",
		QueryRaw: query,
	}

	var programStatus model.Status

	err = r.db.DB().ScanOneContext(ctx, &programStatus, q, args...)
	if err != nil {
		return model.Unknown, err
	}

	return programStatus, nil
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
		Name:     "moder_repository.GetTrainerIdByProgramId",
		QueryRaw: query,
	}

	var id int64

	err = r.db.DB().ScanOneContext(ctx, &id, q, args...)
	if err != nil {
		return 0, err
	}

	return id, nil
}
