package usecase

import (
	"context"
	"github.com/lp9087/go_otello_dashboard_api/internal/entity"
)

//go:generate mockgen -source=interfaces.go -destination=./mocks_test.go -package=usecase_test

type MostLoyalHotelsRepo interface {
	Store(ctx context.Context) ([]entity.MostLoyalHotels, error)
}

type MostLoyalHotelsUseCase interface {
	Get(ctx context.Context) ([]entity.MostLoyalHotels, error)
}
