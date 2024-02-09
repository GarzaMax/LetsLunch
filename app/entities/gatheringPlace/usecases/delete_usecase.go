package usecases

import (
	gathrering_places_repository "cmd/app/entities/gatheringPlace/repository"
	"context"
	"github.com/gofrs/uuid/v5"
)

type DeleteGatheringPlaceUseCase struct {
	gathering_place gathrering_places_repository.PlacesRepository
}

func NewDeleteGatheringPlaceUsecase(gathering_place gathrering_places_repository.PlacesRepository) *DeleteGatheringPlaceUseCase {
	return &DeleteGatheringPlaceUseCase{gathering_place: gathering_place}
}

func (useCase *DeleteGatheringPlaceUseCase) Handle(
	ctx context.Context,
	id uuid.UUID,
) error {
	gathering_place, err := useCase.gathering_place.FindByID(ctx, id)
	if err != nil {
		return err
	}

	err = useCase.gathering_place.Delete(ctx, gathering_place)
	return err
}
