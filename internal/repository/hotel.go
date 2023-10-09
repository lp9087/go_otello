package repository

import (
	"github.com/lp9087/go_otello_lk/internal/database"
	"gorm.io/gorm"
)

type HotelRepository struct {
	db *gorm.DB
}

func CreateHotelRepository(db *gorm.DB) *HotelRepository {
	return &HotelRepository{
		db: db,
	}
}

func (r *HotelRepository) GetList() []database.Hotel {
	var products []database.Hotel
	r.db.Table("hotel_hotel").Find(&products)
	return products
}

func (r *HotelRepository) GetById(id string) (database.Hotel, error) {
	var product database.Hotel
	resp := r.db.Table("hotel_hotel").First(&product, id)
	return product, resp.Error
}
