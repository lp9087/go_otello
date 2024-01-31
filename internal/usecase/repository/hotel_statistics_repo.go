package repository

import (
	"context"
	"fmt"
	"github.com/Masterminds/squirrel"
	"github.com/lp9087/go_otello_dashboard_api/internal/entity"
	"github.com/lp9087/go_otello_dashboard_api/pkg/postgres"
	"time"
)

type PGHotelStatisticRepo struct {
	*postgres.Postgres
}

func NewPGHotelStatisticRepo(pg *postgres.Postgres) *PGHotelStatisticRepo {
	return &PGHotelStatisticRepo{
		pg,
	}
}

func (pg *PGHotelStatisticRepo) Store(ctx context.Context, dateFrom time.Time, dateTo time.Time) (*entity.HotelStatistic, error) {
	sql, args, err := pg.Builder.
		Select("COUNT('id') as registered," +
			"COUNT(case when otello_id is not null then 1 else null end) AS published").
		From("hotel_hotel").
		Where(squirrel.And{
			squirrel.GtOrEq{"created": dateFrom},
			squirrel.LtOrEq{"created": dateTo},
		}).
		ToSql()
	if err != nil {
		return nil, fmt.Errorf("HotelStatistic - Store - pg.Builder: %w", err)
	}
	var res entity.HotelStatistic
	row := pg.Connect.QueryRow(sql, args...)
	err = row.Scan(&res.Registered, &res.Published)
	if err != nil {
		return nil, fmt.Errorf("HotelStatistic - Store - pg.Connect.Select: %w", err)
	}
	return &res, nil
}
