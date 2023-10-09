package handlers

import (
	"github.com/gorilla/mux"
	"github.com/lp9087/go_otello_lk/internal/controllers"
)

func RegisterHotelRoutes(router *mux.Router) {
	router.HandleFunc("/api/v1/hotels", controllers.GetHotels).Methods("GET")
	router.HandleFunc("/api/v1/hotels/{id}", controllers.GetHotelById).Methods("GET")
}
