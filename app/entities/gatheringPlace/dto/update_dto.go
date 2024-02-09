package dto

import (
	"cmd/app/entities/gatheringPlace"
	"cmd/app/models"
)

type UpdateGatheringPlaceDto struct {
	Address     models.Address             `json:"address"`
	AvgPrice    int                        `json:"avgPrice"`
	CusineType  gatheringPlace.CuisineType `json:"cusineType"`
	Rating      int                        `json:"rating"`
	PhoneNumber string                     `json:"phoneNumber"`
}
