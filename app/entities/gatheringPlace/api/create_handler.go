package api

import (
	"cmd/app/entities/gatheringPlace/dto"
	"cmd/app/entities/gatheringPlace/usecases"
	"encoding/json"
	"github.com/gofrs/uuid/v5"
	"net/http"
)

type JsonCreateGatheringPlaceResponse struct {
	GatheringPlaceID uuid.UUID `json:"id"`
}

type CreateGatheringPlaceHandler struct {
	UseCase *usecases.CreateGatheringPlaceUseCase
}

func NewCreateGatheringPlaceHandler(useCase *usecases.CreateGatheringPlaceUseCase) *CreateGatheringPlaceHandler {
	return &CreateGatheringPlaceHandler{UseCase: useCase}
}

func (handler *CreateGatheringPlaceHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	var createGatheringPlaceDto dto.CreateGatheringPlaceDto
	if err := json.NewDecoder(request.Body).Decode(&createGatheringPlaceDto); err != nil {
		marshaledError, _ := json.Marshal(err.Error())

		writer.WriteHeader(http.StatusBadRequest)
		writer.Write(marshaledError)
		return
	}

	command := usecases.CreateGatheringPlaceCommand{}
	command.Address = createGatheringPlaceDto.Address
	command.AvgPrice = createGatheringPlaceDto.AvgPrice
	command.CusineType = createGatheringPlaceDto.CusineType
	command.Rating = createGatheringPlaceDto.Rating
	command.PhoneNumber = createGatheringPlaceDto.PhoneNumber
	command.Description = createGatheringPlaceDto.Description
	command.Title = createGatheringPlaceDto.Title

	gathering_place, err := handler.UseCase.Handle(request.Context(), command)

	if err != nil {
		marshaledError, _ := json.Marshal(err.Error())

		writer.WriteHeader(http.StatusInternalServerError)
		writer.Write(marshaledError)
		return
	}

	response := JsonCreateGatheringPlaceResponse{GatheringPlaceID: gathering_place.ID}

	marshaledResponse, err := json.Marshal(response)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	writer.WriteHeader(http.StatusOK)
	writer.Write(marshaledResponse)
}
