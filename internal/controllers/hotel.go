package controllers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/lp9087/go_otello_lk/internal/services"
	"log"
	"net/http"
)

func GetHotels(w http.ResponseWriter, _ *http.Request) {
	service := services.CreateHotelService()
	hotels := service.GetList()
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err := json.NewEncoder(w).Encode(hotels)
	if err != nil {
		log.Fatal("Cannot parse responses")
	}
}

func GetHotelById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	service := services.CreateHotelService()
	hotels, err := service.GetById(vars["id"])
	w.Header().Set("Content-Type", "application/json")
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		err = json.NewEncoder(w).Encode(err)
		if err != nil {
			log.Fatal("Cannot parse responses")
		}
	} else {
		w.WriteHeader(http.StatusOK)
		err = json.NewEncoder(w).Encode(hotels)
		if err != nil {
			log.Fatal("Cannot parse responses")
		}
	}

}
