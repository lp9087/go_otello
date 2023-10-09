package services

import (
	"github.com/lp9087/go_otello_lk/internal/database"
	"github.com/lp9087/go_otello_lk/internal/repository"
)

func CreateHotelService() HotelService {
	hotelRepository := repository.CreateHotelRepository(database.PostgresDB)
	return HotelService{repository: hotelRepository}
}

type HotelService struct {
	repository *repository.HotelRepository
}

func (s *HotelService) GetList() []database.Hotel {
	return s.repository.GetList()
}
func (s *HotelService) GetById(id string) (database.Hotel, error) {
	return s.repository.GetById(id)
}
