package training

import (
	"context"

	sq "github.com/Masterminds/squirrel"
	"github.com/kirillmc/platform_common/pkg/db"
	"github.com/kirillmc/trainings-server/internal/model"
	"github.com/kirillmc/trainings-server/internal/repository/training/converter"
	modelRepo "github.com/kirillmc/trainings-server/internal/repository/training/model"
)

func (r *repo) GetPublicPrograms(ctx context.Context) ([]*model.TrainProgram, error) {
	builder := sq.Select(idColumm, programName, description, status).
		PlaceholderFormat(sq.Dollar).
		From(trainProgramsTable).
		Where(sq.Eq{status: publicStatusValue})

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, err
	}

	q := db.Query{
		Name:     "training_repository.GetPublicPrograms",
		QueryRaw: query,
	}

	var programs []*modelRepo.TrainProgram

	err = r.db.DB().ScanAllContext(ctx, &programs, q, args...)
	if err != nil {
		return nil, err
	}

	return converter.ToTrainingProgramsFromRepo(programs), nil
}

func (r *repo) GetProgram(ctx context.Context, id int64) (*model.TrainProgram, error) {
	builder := sq.Select(idColumm, programName, description, status).
		PlaceholderFormat(sq.Dollar).
		From(trainProgramsTable).
		Where(sq.Eq{idColumm: id}).
		Limit(1)

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, err
	}

	q := db.Query{
		Name:     "training_repository.GetProgram",
		QueryRaw: query,
	}

	var program modelRepo.TrainProgram

	err = r.db.DB().ScanOneContext(ctx, &program, q, args...)
	if err != nil {
		return nil, err
	}

	return converter.ToTrainingProgramFromRepo(&program), nil
}

func (r *repo) GetTrainDay(ctx context.Context, id int64) (*model.TrainDay, error) {
	builder := sq.Select(idColumm, dayName, description).
		PlaceholderFormat(sq.Dollar).
		From(trainDaysTable).
		Where(sq.Eq{idColumm: id}).
		Limit(1)

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, err
	}

	q := db.Query{
		Name:     "training_repository.GetTrainDay",
		QueryRaw: query,
	}

	var trainDay modelRepo.TrainDay

	err = r.db.DB().ScanOneContext(ctx, &trainDay, q, args...)
	if err != nil {
		return nil, err
	}

	return converter.ToTrainDayFromRepo(&trainDay), nil
}

func (r *repo) GetExercise(ctx context.Context, id int64) (*model.Exercise, error) {
	builder := sq.Select(idColumm, exerciseName, pictures, description).
		PlaceholderFormat(sq.Dollar).
		From(exercisesTable).
		Where(sq.Eq{idColumm: id}).
		Limit(1)

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, err
	}

	q := db.Query{
		Name:     "training_repository.GetExercise",
		QueryRaw: query,
	}

	var exercise modelRepo.Exercise

	err = r.db.DB().ScanOneContext(ctx, &exercise, q, args...)
	if err != nil {
		return nil, err
	}

	return converter.ToExerciseFromRepo(&exercise), nil

}

func (r *repo) GetSet(ctx context.Context, id int64) (*model.Set, error) {
	builder := sq.Select(idColumm, quantity, weight).
		PlaceholderFormat(sq.Dollar).
		From(setsTable).
		Where(sq.Eq{idColumm: id}).
		Limit(1)

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, err
	}

	q := db.Query{
		Name:     "training_repository.GetSet",
		QueryRaw: query,
	}

	var set modelRepo.Set

	err = r.db.DB().ScanOneContext(ctx, &set, q, args...)
	if err != nil {
		return nil, err
	}

	return converter.ToSetFromRepo(&set), nil
}

func (r *repo) GetStatistic(ctx context.Context, id int64) (*model.Statistic, error) {
	builder := sq.Select(idColumm, quantity, weight, setIdColumn, exerciseIdColumn, trainsDayIdColumn, programIdColumn).
		PlaceholderFormat(sq.Dollar).
		From(statisticsTable).
		Where(sq.Eq{idColumm: id}).
		Limit(1)

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, err
	}

	q := db.Query{
		Name:     "training_repository.GetStatistic",
		QueryRaw: query,
	}

	var statistic modelRepo.Statistic

	err = r.db.DB().ScanOneContext(ctx, &statistic, q, args...)
	if err != nil {
		return nil, err
	}

	return converter.ToStatisticFromRepo(&statistic), nil
}
