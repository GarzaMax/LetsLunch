package query

import (
	"cmd/app/entities/gatheringPlace"
	"context"
	"database/sql"
	"fmt"
)

func Create(ctx context.Context, place *gatheringPlace.GatheringPlace, db *sql.DB) error {
	const query = `INSERT INTO gathering_places(id, country, city, street_name, house_number, building_number,
                             average_price, cuisine_type, rating, phone_number, description, photo_link, title)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13)`
	_, err := db.Exec(query, place.ID, place.Address.Country, place.Address.City, place.Address.StreetName,
		place.Address.HouseNumber, place.Address.BuildingNumber, place.AveragePrice, place.CuisineType, place.Rating,
		place.PhoneNumber, place.Description, place.PhotoLink, place.Title)
	if err != nil {
		return fmt.Errorf("database query execution error: %w", err)
	}
	return nil
}
