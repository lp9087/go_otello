package entity

type MostLoyalHotels struct {
	ID          string `db:"id"`
	HotelName   string `db:"hotel_name"`
	ExternalID  string `db:"external_id"`
	RoomTypes   int    `db:"ROOM_TYPES_AMOUNT"`
	Rates       int    `db:"RATES_AMOUNT"`
	TotalAmount int    `db:"TOTAL_AMOUNT"`
}

type HotelStatistic struct {
	Registered int `db:"registered"`
	Published  int `db:"published"`
}
