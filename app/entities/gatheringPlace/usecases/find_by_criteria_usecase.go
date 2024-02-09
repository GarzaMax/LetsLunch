package usecases

import (
	"cmd/app/entities/gatheringPlace"
	"cmd/app/entities/gatheringPlace/query"
	gathering_place_repository "cmd/app/entities/gatheringPlace/repository"
	"cmd/app/models"
	"context"
)

type FindGatheringPlacesByCriteriaUseCase struct {
	gathering_place gathering_place_repository.PlacesRepository
}

func NewFindGatheringPlacesByCriteriaUseCase(gathering_places gathering_place_repository.PlacesRepository) *FindGatheringPlacesByCriteriaUseCase {
	return &FindGatheringPlacesByCriteriaUseCase{gathering_place: gathering_places}
}

func (useCase *FindGatheringPlacesByCriteriaUseCase) Handle(ctx context.Context, criteria models.FindGatheringPlaceCriteria) ([]gatheringPlace.GatheringPlace, error) {
	address := models.Address{
		Country:        criteria.Country,
		City:           criteria.City,
		StreetName:     criteria.StreetName,
		HouseNumber:    criteria.HouseNumber,
		BuildingNumber: criteria.BuildingNumber,
	}

	queryCriteria := query.FindCriteria{}
	queryCriteria.Address = &address
	queryCriteria.InitiatorID.UUID = criteria.InitiatorID
	queryCriteria.InitiatorID.Valid = false
	queryCriteria.Rating = criteria.Rating

	if address.City == "" || address.Country == "" || address.StreetName == "" || address.BuildingNumber == 0 {
		queryCriteria.Address = nil
	}
	if criteria.CuisineType == nil {
		queryCriteria.CuisineType.Valid = false
	} else {
		queryCriteria.CuisineType.Int16 = int16(*criteria.CuisineType)
		queryCriteria.CuisineType.Valid = true
	}

	queryResult, err := useCase.gathering_place.FindByCriteria(ctx, queryCriteria)

	if err != nil {
		return nil, err
	}
	if queryResult == nil {
		return []gatheringPlace.GatheringPlace{}, nil
	}
	return queryResult, nil
}
