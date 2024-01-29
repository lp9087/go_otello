package usecase

import (
	"context"
	"github.com/lp9087/go_otello_dashboard_api/internal/entity"
)

type FirstDashboardRepo interface {
	Store(ctx context.Context) ([]entity.FirstDashboard, error)
}

type FirstDashboard interface {
	Get(ctx context.Context) ([]entity.FirstDashboard, error)
}
