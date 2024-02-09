package repository

import (
	"cmd/app/entities/gatheringPlace"
	"cmd/app/entities/gatheringPlace/query"
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/gofrs/uuid/v5"
)

//go:generate mockery --name=PlacesRepository
type PlacesRepository interface {
	FindByCriteria(ctx context.Context, criteria query.FindCriteria) ([]gatheringPlace.GatheringPlace, error)
	FindByID(ctx context.Context, id uuid.UUID) (*gatheringPlace.GatheringPlace, error)
	Create(ctx context.Context, place *gatheringPlace.GatheringPlace) (*gatheringPlace.GatheringPlace, error)
	Update(ctx context.Context, place *gatheringPlace.GatheringPlace) (*gatheringPlace.GatheringPlace, error)
	Delete(ctx context.Context, place *gatheringPlace.GatheringPlace) error
}

type PlacesDatabaseRepository struct {
	db *sql.DB
}

func NewPlacesDatabaseRepository(providedConnection *sql.DB) *PlacesDatabaseRepository {
	return &PlacesDatabaseRepository{db: providedConnection}
}
func (repository *PlacesDatabaseRepository) FindByID(ctx context.Context, id uuid.UUID) (*gatheringPlace.GatheringPlace, error) {
	var currentPlace gatheringPlace.GatheringPlace
	row := query.FindByID(ctx, id, repository.db)
	if err := row.Scan(&currentPlace.ID, &currentPlace.Address.Country, &currentPlace.Address.City, &currentPlace.Address.StreetName, &currentPlace.Address.HouseNumber, &currentPlace.Address.BuildingNumber,
		&currentPlace.AveragePrice, &currentPlace.CuisineType, &currentPlace.Rating, &currentPlace.PhoneNumber,
		&currentPlace.Description, &currentPlace.PhotoLink, &currentPlace.Title); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("no such gathering place: %w", err)
		}
		return nil, fmt.Errorf("cannot query the database: %w", err)
	}
	return &currentPlace, nil
}
func (repository *PlacesDatabaseRepository) FindByCriteria(ctx context.Context, criteria query.FindCriteria) ([]gatheringPlace.GatheringPlace, error) {
	var places []gatheringPlace.GatheringPlace
	rows, err := query.FindByCriteria(ctx, criteria, repository.db)
	if err != nil {
		return nil, fmt.Errorf("cannot query the database: %w", err)
	}
	var currentPlace gatheringPlace.GatheringPlace
	for rows.Next() {
		if err = rows.Scan(&currentPlace.ID, &currentPlace.Address.Country, &currentPlace.Address.City, &currentPlace.Address.StreetName, &currentPlace.Address.HouseNumber, &currentPlace.Address.BuildingNumber,
			&currentPlace.AveragePrice, &currentPlace.CuisineType, &currentPlace.Rating, &currentPlace.PhoneNumber,
			&currentPlace.Description, &currentPlace.PhotoLink, &currentPlace.Title); err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				return nil, fmt.Errorf("no such gathering place: %w", err)
			}
			return nil, fmt.Errorf("cannot query the database: %w", err)
		}
		places = append(places, currentPlace)
	}
	return places, nil
}
func (repository *PlacesDatabaseRepository) Create(ctx context.Context, place *gatheringPlace.GatheringPlace) (*gatheringPlace.GatheringPlace, error) {
	var err = query.Create(ctx, place, repository.db)
	if err != nil {
		return place, fmt.Errorf("place cannot be created: %w", err)
	}
	return place, nil
}

func (repository *PlacesDatabaseRepository) Update(ctx context.Context, place *gatheringPlace.GatheringPlace) (*gatheringPlace.GatheringPlace, error) {
	var err = query.Update(ctx, place, repository.db)
	if err != nil {
		return place, fmt.Errorf("place cannot be updated: %w", err)
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
