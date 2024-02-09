package usecases

import (
	domain "cmd/app/entities/gatheringPlace"
	gathrering_places_repository "cmd/app/entities/gatheringPlace/repository"
	"cmd/app/models"
	"context"
	"github.com/gofrs/uuid/v5"
)

type FindGatheringPlaceByIdResponse struct {
	ID          uuid.UUID          `json:"id"`
	Address     models.Address     `json:"address"`
	AvgPrice    int                `json:"averagePrice"`
	CusineType  domain.CuisineType `json:"cusineType"`
	Rating      int                `json:"rating"`
	PhoneNumber string             `json:"phoneNumber"`
}

type FindGatheringPlaceByIdUseCase struct {
	gathering_place gathrering_places_repository.PlacesRepository
}

func NewFindGatheringPlaceByIdUseCase(gathering_places gathrering_places_repository.PlacesRepository) *FindGatheringPlaceByIdUseCase {
	return &FindGatheringPlaceByIdUseCase{gathering_place: gathering_places}
}

func (useCase *FindGatheringPlaceByIdUseCase) Handle(ctx context.Context, id uuid.UUID) (*FindGatheringPlaceByIdResponse, error) {
	gathering_places, err := useCase.gathering_place.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return &FindGatheringPlaceByIdResponse{
		ID:          gathering_places.ID,
		Address:     gathering_places.Address,
		AvgPrice:    gathering_places.AveragePrice,
		CusineType:  gathering_places.CuisineType,
		Rating:      gathering_places.Rating,
		PhoneNumber: gathering_places.PhoneNumber,
	}, nil
}
