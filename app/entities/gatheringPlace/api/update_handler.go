package api

import (
	"cmd/app/entities/gatheringPlace/dto"
	"cmd/app/entities/gatheringPlace/usecases"
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"github.com/gofrs/uuid/v5"
	"net/http"
)

type JsonUpdateGatheringPlaceResponse struct {
	ID uuid.UUID `json:"id"`
}

type UpdateGatheringPlaceHandler struct {
	useCase *usecases.UpdateGatheringPlaceUseCase
}

func NewUpdateGatheringPlaceHandler(useCase *usecases.UpdateGatheringPlaceUseCase) *UpdateGatheringPlaceHandler {
	return &UpdateGatheringPlaceHandler{useCase: useCase}
}

func (handler *UpdateGatheringPlaceHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	var updateGatheringPlaceDto dto.UpdateGatheringPlaceDto
	if err := json.NewDecoder(request.Body).Decode(&updateGatheringPlaceDto); err != nil {
		marshaledError, _ := json.Marshal(err.Error())

		writer.WriteHeader(http.StatusBadRequest)
		writer.Write(marshaledError)
		return
	}

	command := usecases.UpdateGatheringPlaceCommand{}
	command.Address = updateGatheringPlaceDto.Address
	command.AvgPrice = updateGatheringPlaceDto.AvgPrice
	command.CusineType = updateGatheringPlaceDto.CusineType
	command.Rating = updateGatheringPlaceDto.Rating
	command.PhoneNumber = updateGatheringPlaceDto.PhoneNumber
	id := chi.URLParam(request, "placeID")

	uuidID, err := uuid.FromString(id)
	if err != nil {
		marshaledError, _ := json.Marshal(err.Error())

		writer.WriteHeader(http.StatusBadRequest)
		writer.Write(marshaledError)
		return
	}

	response, err := handler.useCase.Handle(request.Context(), command, uuidID)

	marshaledResponse, err := json.Marshal(response)

	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	writer.WriteHeader(http.StatusOK)
	writer.Write(marshaledResponse)
}
