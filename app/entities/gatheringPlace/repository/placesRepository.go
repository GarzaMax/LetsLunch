package repository

import (
	"cmd/app/entities/gatheringPlace"
	"cmd/app/entities/gatheringPlace/query"
	"context"
	"database/sql"
	"fmt"
)

//go:generate mockery --name=PlacesRepository
type PlacesRepository interface {
	FindByCriteria(ctx context.Context, criteria query.FindCriteria) ([]gatheringPlace.GatheringPlace, error)
	Create(ctx context.Context, place gatheringPlace.GatheringPlace) (*gatheringPlace.GatheringPlace, error)
	Update(ctx context.Context, place gatheringPlace.GatheringPlace) (*gatheringPlace.GatheringPlace, error)
	Delete(ctx context.Context, place gatheringPlace.GatheringPlace) error
}

type PlacesDatabaseRepository struct {
	db *sql.DB
}

func NewPlacesDatabaseRepository(providedConnection *sql.DB) *PlacesDatabaseRepository {
	return &PlacesDatabaseRepository{db: providedConnection}
}
func (repository *PlacesDatabaseRepository) FindByCriteria(ctx context.Context, criteria query.FindCriteria) ([]gatheringPlace.GatheringPlace, error) {
	var places []gatheringPlace.GatheringPlace
	rows, err := query.FindByCriteria(ctx, criteria, repository.db)
	if err != nil {
		return nil, fmt.Errorf("cannot query the database %w", err)
	}
	var currentPlace gatheringPlace.GatheringPlace
	for rows.Next() {
		if err = rows.Scan(&currentPlace.ID, &currentPlace.Address.Country, &currentPlace.Address.City, &currentPlace.Address.StreetName, &currentPlace.Address.HouseNumber, &currentPlace.Address.BuildingNumber,
			&currentPlace.AveragePrice, &currentPlace.CuisineType, &currentPlace.Rating, &currentPlace.PhoneNumber); err != nil {
			return nil, fmt.Errorf("cannot query the database %w", err)
		}
		places = append(places, currentPlace)
	}
	return places, nil
}

func (repository *PlacesDatabaseRepository) Create(ctx context.Context, place *gatheringPlace.GatheringPlace) (*gatheringPlace.GatheringPlace, error) {
	var err = query.Create(ctx, place, repository.db)
	if err != nil {
		return place, fmt.Errorf("meeting cannot be created: %v", err)
	}
	return place, nil
}

func (repository *PlacesDatabaseRepository) Update(ctx context.Context, place *gatheringPlace.GatheringPlace) (*gatheringPlace.GatheringPlace, error) {
	var err = query.Update(ctx, place, repository.db)
	if err != nil {
		return place, fmt.Errorf("place cannot be updated: %v", err)
	}
	return place, nil
}

func (repository *PlacesDatabaseRepository) Delete(ctx context.Context, place *gatheringPlace.GatheringPlace) error {
	var err = query.Delete(ctx, place, repository.db)
	if err != nil {
		return fmt.Errorf("place cannot be deleted: %v", err)
	}
	return nil
}