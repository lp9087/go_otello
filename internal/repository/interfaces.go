package repository

import (
	"context"
	"github.com/lp9087/go_otello_dashboard_api/internal/entity"
)

type FirstDashboardRepo interface {
	Get(ctx context.Context) ([]entity.FirstDashboard, error)
}
