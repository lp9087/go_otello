package usecase

import (
	"context"
	"fmt"
	"github.com/lp9087/go_otello_dashboard_api/internal/entity"
)

type MostLoyalUseCase struct {
	repo MostLoyalHotelsRepo
}

func NewMostLoyalHotelsUseCase(fr MostLoyalHotelsRepo) *MostLoyalUseCase {
	return &MostLoyalUseCase{
		repo: fr,
	}
}

func (pr *MostLoyalUseCase) Get(ctx context.Context) ([]entity.MostLoyalHotels, error) {
	entities, err := pr.repo.Store(ctx)
	if err != nil {
		return []entity.MostLoyalHotels{}, fmt.Errorf("MostLoyalUseCase - Get - repo.Store: %w", err)
	}

	return entities, nil
}
