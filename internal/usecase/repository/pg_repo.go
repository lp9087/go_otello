package repository

import (
	"context"
	"fmt"
	"github.com/lp9087/go_otello_dashboard_api/internal/entity"
	"github.com/lp9087/go_otello_dashboard_api/pkg/postgres"
)

type PGMostLoyalHotelsRepo struct {
	*postgres.Postgres
}

func NewPGMostLoyalHotelsRepo(pg *postgres.Postgres) *PGMostLoyalHotelsRepo {
	return &PGMostLoyalHotelsRepo{
		pg,
	}
}

func (pg *PGMostLoyalHotelsRepo) Store(ctx context.Context) ([]entity.MostLoyalHotels, error) {
	sql, _, err := pg.Builder.
		Select("h.id, " +
			"h.hotel_name, " +
			"h.external_id, " +
			"(SELECT COUNT(*) FROM room_type_hotelroomtype rth WHERE rth.hotel_id = h.id AND rth.deprecated = False) AS ROOM_TYPES_AMOUNT," +
			"(SELECT COUNT(*) FROM rates_plan rp WHERE rp.hotel_id = h.id AND rp.deprecated = False) AS RATES_AMOUNT," +
			"(SELECT COUNT(*) FROM checkerboard_roomavailability cr " +
			"JOIN room_type_hotelroomtype rth ON rth.id = cr.hotel_room_type_id " +
			"WHERE rth.hotel_id = h.id AND cr.date >= CURRENT_DATE) + (SELECT COUNT(*) FROM checkerboard_price cp " +
			"JOIN rates_roomplan rr ON rr.id = cp.room_plan_id " +
			"JOIN rates_plan rp ON rp.id = rr.plan_id WHERE rp.hotel_id = h.id AND cp.price_date >= CURRENT_DATE) AS TOTAL_AMOUNT").
		From("hotel_hotel h").
		Where("h.otello_id IS NOT NULL").
		OrderBy("total_amount DESC").
		Limit(5).
		ToSql()
	if err != nil {
		return nil, fmt.Errorf("MostLoyalHotels - Store - pg.Builder: %w", err)
	}

	rows, err := pg.Connect.QueryxContext(ctx, sql)
	if err != nil {
		return nil, fmt.Errorf("MostLoyalHotels - Store - pg.Connect.QueryxContext: %w", err)
	}
	defer rows.Close()

	entities := make([]entity.MostLoyalHotels, 0)

	for rows.Next() {
		e := entity.MostLoyalHotels{}

		err = rows.Scan(&e.ID, &e.HotelName, &e.ExternalID, &e.RoomTypes, &e.Rates, &e.TotalAmount)
		if err != nil {
			return nil, fmt.Errorf("MostLoyalHotels - Store - rows.Scan: %w", err)
		}

		entities = append(entities, e)
	}

	return entities, nil
}
