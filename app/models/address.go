package models

type Address struct {
	Country        string `json:"country"`
	City           string `json:"city"`
	StreetName     string `json:"streetName"`
	HouseNumber    string `json:"houseNumber"`
	BuildingNumber int    `json:"buildingNumber"`
}
