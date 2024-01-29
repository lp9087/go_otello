package app

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/lp9087/go_otello_dashboard_api/config"
	"github.com/lp9087/go_otello_dashboard_api/internal/controller/rest/v1"
	"github.com/lp9087/go_otello_dashboard_api/internal/usecase"
	"github.com/lp9087/go_otello_dashboard_api/internal/usecase/repository"
	"github.com/lp9087/go_otello_dashboard_api/pkg/postgres"
	"log"
	"net/http"
	"os/signal"
	"syscall"
	"time"
)

func Run(cfg *config.Config) {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()
	// Connect Database

	db, err := postgres.New(&cfg.DB)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	// Use case
	mostLoyalHotelsUseCase := usecase.NewFirstDashboardUseCase(
		repository.NewPGFirstDashboardRepo(db.Connect),
	)

	// Start Router
	router := gin.New()
	groupRouter := v1.NewRouter(router)
	v1.NewDashboardRoutes(groupRouter, mostLoyalHotelsUseCase)

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", cfg.HTTP.Port),
		Handler: router,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	<-ctx.Done()

	stop()
	log.Println("shutting down gracefully, press Ctrl+C again to force")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown: ", err)
	}

	log.Println("Server exiting")
}
