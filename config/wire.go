//go:build wireinject
// +build wireinject

package config

import (
	"github.com/google/wire"
	"github.com/jmoiron/sqlx"
	"github.com/lp9087/go_otello_dashboard_api/internal/usecase"
	"github.com/lp9087/go_otello_dashboard_api/internal/usecase/repository"
)

func InitializeFirstDashboardUseCase(db *sqlx.DB) *usecase.FirstDashboardUseCase {
	wire.Build(wire.Bind(new(usecase.FirstDashboardRepo), new(*repository.PGFirstDashboardRepo)), repository.NewPGFirstDashboardRepo, usecase.NewFirstDashboardUseCase)
	return &usecase.FirstDashboardUseCase{}
}
