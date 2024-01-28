package usecase

import "github.com/lp9087/go_otello_dashboard_api/internal/repository"

type FirstDashboardService struct {
	repo repository.FirstDashboardRepo
}

func NewFirstDashboardService(fr repository.FirstDashboardRepo) *FirstDashboardService {
	return &FirstDashboardService{
		repo: fr,
	}
}
