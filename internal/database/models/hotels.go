package models

type Hotel struct {
	ID        int             `json:"id"`
	HotelName string          `json:"hotel_name"`
	Address   string          `json:"address"`
	OtelloId  string          `json:"otello_id"`
	RoomType  []HotelRoomType `json:"room_types"`
}

func (Hotel) TableName() string {
	return "hotel_hotel"
}

type HotelRoomType struct {
	ID      int    `json:"id"`
	HotelId string `json:"-"`
	Hotel   Hotel  `json:"-"`
}

func (HotelRoomType) TableName() string {
	return "room_type_hotelroomtype"
}
