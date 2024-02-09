package usecases

import (
	"cmd/app/entities/gatheringPlace"
	gathrering_places_repository "cmd/app/entities/gatheringPlace/repository"
	"cmd/app/models"
	"context"
	"github.com/gofrs/uuid/v5"
)

type UpdateGatheringPlaceResponse struct {
	ID uuid.UUID `json:"id"`
}

type UpdateGatheringPlaceUseCase struct {
	gathering_place gathrering_places_repository.PlacesRepository
}

type UpdateGatheringPlaceCommand struct {
	Address     models.Address
	AvgPrice    int
	CusineType  gatheringPlace.CuisineType
	Rating      int
	PhoneNumber string
}

func NewUpdateGatheringPlaceUseCase(gathering_place gathrering_places_repository.PlacesRepository) *UpdateGatheringPlaceUseCase {
	return &UpdateGatheringPlaceUseCase{gathering_place: gathering_place}
}

func (useCase *UpdateGatheringPlaceUseCase) Handle(
	ctx context.Context,
	command UpdateGatheringPlaceCommand,
	id uuid.UUID,
) (*UpdateGatheringPlaceResponse, error) {
	gathering_place, err := useCase.gathering_place.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}

	gathering_place.Address = command.Address
	gathering_place.AveragePrice = command.AvgPrice
	gathering_place.CuisineType = command.CusineType
	gathering_place.Rating = command.Rating
	gathering_place.PhoneNumber = command.PhoneNumber

	_, err = useCase.gathering_place.Update(ctx, gathering_place)
	if err != nil {
		return nil, err
	}

	return &UpdateGatheringPlaceResponse{ID: id}, nil
}
