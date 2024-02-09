package api

import (
	"cmd/app/entities/meeting/dto"
	"cmd/app/entities/meeting/usecases"
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"github.com/gofrs/uuid/v5"
	"net/http"
)

type JsonUpdateMeetingResponse struct {
	ID uuid.UUID `json:"id"`
}

type UpdateMeetingHandler struct {
	useCase *usecases.UpdateMeetingUseCase
}

func NewUpdateMeetingHandler(useCase *usecases.UpdateMeetingUseCase) *UpdateMeetingHandler {
	return &UpdateMeetingHandler{useCase: useCase}
}

func (handler *UpdateMeetingHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	var updateMeetingDto dto.UpdateMeetingDto
	if err := json.NewDecoder(request.Body).Decode(&updateMeetingDto); err != nil {
		marshaledError, _ := json.Marshal(err.Error())

		writer.WriteHeader(http.StatusBadRequest)
		writer.Write(marshaledError)
		return
	}

	command := usecases.UpdateMeetingCommand{}
	command.GatheringPlaceID = updateMeetingDto.GatheringPlaceID
	command.InitiatorsID = updateMeetingDto.InitiatorsID
	command.StartTime = updateMeetingDto.StartTime
	command.EndTime = updateMeetingDto.EndTime
	command.UsersQuantity = updateMeetingDto.UsersQuantity
	command.State = updateMeetingDto.State

	id := chi.URLParam(request, "meetingID")

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
