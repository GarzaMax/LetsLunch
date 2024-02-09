package usecases

import (
	"cmd/app/entities/gatheringPlace"
	gathering_place_repository "cmd/app/entities/gatheringPlace/repository"
	"cmd/app/models"
	"context"
	"fmt"
)

type CreateGatheringPlaceUseCase struct {
	gathering_places gathering_place_repository.PlacesRepository
}

type CreateGatheringPlaceCommand struct {
	Address     models.Address
	AvgPrice    int
	CusineType  gatheringPlace.CuisineType
	Rating      int
	PhoneNumber string
	Description string
	Title       string
	PhotoLink   string
}

func NewCreateGatheringPlaceUseCase(gathering_places gathering_place_repository.PlacesRepository) *CreateGatheringPlaceUseCase {
	return &CreateGatheringPlaceUseCase{gathering_places: gathering_places}
}

func (useCase *CreateGatheringPlaceUseCase) Handle(
	ctx context.Context,
	command CreateGatheringPlaceCommand,
) (*gatheringPlace.GatheringPlace, error) {
	gatheringPlace := gatheringPlace.NewGatheringPlace(
		command.Address,
		command.AvgPrice,
		command.CusineType,
		command.Rating,
		command.PhoneNumber,
		command.Description,
		command.Title,
		command.PhotoLink)

	_, err := useCase.gathering_places.Create(ctx, gatheringPlace)
	if err != nil {
		return nil, fmt.Errorf("gathering_place : create gathering place %w", err)
	}

	return gatheringPlace, nil
}
