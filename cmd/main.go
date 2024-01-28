package main

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/lp9087/go_otello_dashboard_api/configs"
	rest_ "github.com/lp9087/go_otello_dashboard_api/internal/controller/rest"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func GetEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func getRoutes() *gin.Engine {
	router := rest_.SetupRouter()
	rest_.AddDashboardRoutes(router)
	return router
}

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()
	// Get ENV
	GetEnv()
	// Connect Database
	dbConfig := configs.NewDBConfig(
		os.Getenv("POSTGRES_DB"),
		os.Getenv("POSTGRES_HOST"),
		os.Getenv("POSTGRES_PORT"),
		os.Getenv("POSTGRES_SSL"),
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
	)
	db, err := configs.GetDBConnect(dbConfig)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	// Start Router
	router := getRoutes()

	srv := &http.Server{
		Addr:    ":8080",
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
