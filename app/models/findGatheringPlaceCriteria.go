package models

import (
	"github.com/gofrs/uuid/v5"
)

type FindGatheringPlaceCriteria struct {
	Country        string
	City           string
	StreetName     string
	HouseNumber    string
	BuildingNumber int
	InitiatorID    uuid.UUID
	CuisineType    *int
	Rating         int
}
