package dto

import (
	"cmd/app/entities/gatheringPlace"
	"cmd/app/models"
)

type CreateGatheringPlaceDto struct {
	Address     models.Address             `json:"address"`
	AvgPrice    int                        `json:"avgPrice"`
	CusineType  gatheringPlace.CuisineType `json:"cusine_Type"`
	Rating      int                        `json:"rating"`
	PhoneNumber string                     `json:"phoneNumber"`
	Description string                     `json:"description"`
	Title       string                     `json:"title"`
	PhotoLink   string                     `json:"photoLink"`
}
