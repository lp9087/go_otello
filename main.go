package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/lp9087/go_otello_lk/internal/database"
	"github.com/lp9087/go_otello_lk/internal/handlers"
	"log"
	"net/http"
)

func main() {
	// Load Configurations from config.json using Viper
	LoadAppConfig()
	// Initialize Database
	database.Connect(AppConfig.ConnectionString)
	//database.Migrate()

	// Initialize the router
	router := mux.NewRouter().StrictSlash(true)
	// Register Routes
	handlers.RegisterHotelRoutes(router)
	// Start the server
	log.Println(fmt.Sprintf("Starting Server on port %s", AppConfig.Port))
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", AppConfig.Port), router))
}
