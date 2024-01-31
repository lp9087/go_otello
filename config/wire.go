//go:build wireinject
// +build wireinject

package config

import (
	"github.com/google/wire"
	"github.com/lp9087/go_otello_dashboard_api/internal/usecase"
	"github.com/lp9087/go_otello_dashboard_api/internal/usecase/repository"
	"github.com/lp9087/go_otello_dashboard_api/pkg/postgres"
)

func InitializeMostLoyalHotelsUseCase(pg *postgres.Postgres) *usecase.MostLoyalUseCase {
	wire.Build(wire.Bind(new(usecase.MostLoyalHotelsRepo), new(*repository.PGMostLoyalHotelsRepo)), repository.NewPGMostLoyalHotelsRepo, usecase.NewMostLoyalHotelsUseCase)
	return &usecase.MostLoyalUseCase{}
}

func InitializeHotelsStatisticUseCase(pg *postgres.Postgres) *usecase.StatisticUseCase {
	wire.Build(wire.Bind(new(usecase.HotelStatisticRepo), new(*repository.PGHotelStatisticRepo)), repository.NewPGHotelStatisticRepo, usecase.NewHotelStatisticUseCase)
	return &usecase.StatisticUseCase{}
}
