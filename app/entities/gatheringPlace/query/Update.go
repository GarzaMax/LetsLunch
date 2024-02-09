package query

import (
	"cmd/app/entities/gatheringPlace"
	"context"
	"database/sql"
	"fmt"
)

func Update(ctx context.Context, place *gatheringPlace.GatheringPlace, db *sql.DB) error {
	const query = `UPDATE gathering_places
    SET id = $1, country = $2, city = $3, street_name = $4, house_number = $5, building_number = $6,
                             average_price = $7, cuisine_type = $8, rating = $9, phone_number = $10,
                             description = $11, photo_link = $12, title = $13
    WHERE id = $1`
	_, err := db.Exec(query, place.ID, place.Address.Country, place.Address.City, place.Address.StreetName,
		place.Address.HouseNumber, place.Address.BuildingNumber, place.AveragePrice, place.CuisineType, place.Rating,
		place.PhoneNumber, place.Description, place.PhotoLink, place.Title)
	if err != nil {
		return fmt.Errorf("database query execution error: %w", err)
	}
	return nil
}

/*func Update(ctx context.Context, place *gatheringPlace.GatheringPlace, db *sql.DB) error {
	const query = `UPDATE gatheringPlaces SET rating = $1 WHERE id = $2`
	_, err := db.Exec(query, place.Rating, place.ID)
	if err != nil {
		return fmt.Errorf("database query execution error: %w", err)
	}
	return nil
}*/
