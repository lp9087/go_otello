package usecase

import (
	"context"
	"fmt"
	"github.com/lp9087/go_otello_dashboard_api/internal/entity"
	"time"
)

type StatisticUseCase struct {
	repo HotelStatisticRepo
}

func NewHotelStatisticUseCase(repo HotelStatisticRepo) *StatisticUseCase {
	return &StatisticUseCase{
		repo: repo,
	}
}

func (pr *StatisticUseCase) Get(ctx context.Context, dateFrom time.Time, dateTo time.Time) (*entity.HotelStatistic, error) {
	entities, err := pr.repo.Store(ctx, dateFrom, dateTo)
	if err != nil {
		return &entity.HotelStatistic{}, fmt.Errorf("HotelStatisticUseCase - Get - repo.Store: %w", err)
	}

	return entities, nil
}
