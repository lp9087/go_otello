package repository

import (
	"github.com/lp9087/go_otello_lk/internal/database/models"
	"gorm.io/gorm"
	"log"
)

type HotelRepository struct {
	db *gorm.DB
}

func CreateHotelRepository(db *gorm.DB) *HotelRepository {
	return &HotelRepository{
		db: db,
	}
}

func (r *HotelRepository) GetList() []models.Hotel {
	var hotels []models.Hotel
	err := r.db.Preload("RoomType").Find(&hotels).Error
	if err != nil {
		log.Fatal("Can't load data from SQL")
	}
	return hotels
}

func (r *HotelRepository) GetById(id string) (models.Hotel, error) {
	var product models.Hotel
	resp := r.db.First(&product, id)
	return product, resp.Error
}
