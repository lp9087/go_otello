package usecase

import "github.com/lp9087/go_otello_dashboard_api/internal/entity"
import "context"

type FirstDashboardUseCase struct {
	repo FirstDashboardRepo
}

func NewFirstDashboardUseCase(fr FirstDashboardRepo) *FirstDashboardUseCase {
	return &FirstDashboardUseCase{
		repo: fr,
	}
}

func (pr *FirstDashboardUseCase) Get(ctx context.Context) ([]entity.FirstDashboard, error) {
	return nil, nil
}
