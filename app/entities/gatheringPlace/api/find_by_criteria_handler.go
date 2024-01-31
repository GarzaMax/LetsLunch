package api

import (
	"cmd/app/entities/gatheringPlace/usecases"
	"cmd/app/models"
	"encoding/json"
	"net/http"
	"strconv"
)

type FindGatheringPlacesByCriteriaHandler struct {
	useCase *usecases.FindGatheringPlacesByCriteriaUseCase
}

func NewFindGatheringPlacesByCriteriaHandler(useCase *usecases.FindGatheringPlacesByCriteriaUseCase) *FindGatheringPlacesByCriteriaHandler {
	return &FindGatheringPlacesByCriteriaHandler{useCase: useCase}
}

func (handler *FindGatheringPlacesByCriteriaHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	//initiator := request.URL.Query().Get("initiator")
	rating := request.URL.Query().Get("rating")
	country := request.URL.Query().Get("country")
	city := request.URL.Query().Get("city")
	streetName := request.URL.Query().Get("street_name")
	buildingNubmer := request.URL.Query().Get("building_number")
	houseNumber := request.URL.Query().Get("house_number")

	findCriteria := models.FindGatheringPlaceCriteria{}
	findCriteria.Country = country
	findCriteria.City = city
	findCriteria.StreetName = streetName
	if buildingNubmer != "" {
		convertedBuildingNumber, err := strconv.Atoi(buildingNubmer)
		if err != nil {
			marshaledError, _ := json.Marshal(err)

			writer.WriteHeader(http.StatusBadRequest)
			writer.Write(marshaledError)
			return
		}
		findCriteria.BuildingNumber = convertedBuildingNumber
	}

	if rating != "" {
		convertedRating, err := strconv.Atoi(rating)
		if err != nil {
			marshaledError, _ := json.Marshal(err)

			writer.WriteHeader(http.StatusBadRequest)
			writer.Write(marshaledError)
			return
		}
		findCriteria.BuildingNumber = convertedRating
	}

	if houseNumber != "" {
		findCriteria.HouseNumber = houseNumber
	}

	gatheringPlaces, err := handler.useCase.Handle(request.Context(), findCriteria)
	if err != nil {
		marshaledError, _ := json.Marshal(err)

		writer.WriteHeader(http.StatusInternalServerError)
		writer.Write(marshaledError)
		return
	}

	response := gatheringPlaces

	marshaledResponse, err := json.Marshal(response)
	if err != nil {
		marshaledError, _ := json.Marshal(err)

		writer.WriteHeader(http.StatusInternalServerError)
		writer.Write(marshaledError)
		return
	}

	writer.WriteHeader(http.StatusOK)
	writer.Write(marshaledResponse)
}