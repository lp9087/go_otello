package repository

import (
	"context"
	"github.com/jmoiron/sqlx"
	"github.com/lp9087/go_otello_dashboard_api/internal/entity"
)

type PGFirstDashboardRepo struct {
	db *sqlx.DB
}

func NewPGFirstDashboardRepo(db *sqlx.DB) *PGFirstDashboardRepo {
	return &PGFirstDashboardRepo{
		db: db,
	}
}

func (pr *PGFirstDashboardRepo) Store(ctx context.Context) ([]entity.FirstDashboard, error) {
	return nil, nil
}
