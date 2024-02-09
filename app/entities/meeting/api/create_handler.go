package api

import (
	"cmd/app/entities/meeting/dto"
	"cmd/app/entities/meeting/usecases"
	"encoding/json"
	"github.com/gofrs/uuid/v5"
	"net/http"
)

type JsonCreateMeetingResponse struct {
	MeetingID uuid.UUID `json:"id"`
}

type CreateMeetingHandler struct {
	useCase *usecases.CreateMeetingUseCase
}

func NewCreateMeetingHandler(useCase *usecases.CreateMeetingUseCase) *CreateMeetingHandler {
	return &CreateMeetingHandler{useCase: useCase}
}

func (handler *CreateMeetingHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	var createMeetingDto dto.CreateMeetingDto
	if err := json.NewDecoder(request.Body).Decode(&createMeetingDto); err != nil {
		marshaledError, _ := json.Marshal(err.Error())

		writer.WriteHeader(http.StatusBadRequest)
		writer.Write(marshaledError)
		return
	}

	command := usecases.CreateMeetingCommand{}
	command.GatheringPlaceID = createMeetingDto.GatheringPlaceID
	command.InitiatorsID = createMeetingDto.InitiatorsID
	command.StartTime = createMeetingDto.StartTime
	command.EndTime = createMeetingDto.EndTime
	command.UsersQuantity = createMeetingDto.UsersQuantity
	command.State = createMeetingDto.State

	meeting, err := handler.useCase.Handle(request.Context(), command)

	if err != nil {
		marshaledError, _ := json.Marshal(err.Error())

		writer.WriteHeader(http.StatusInternalServerError)
		writer.Write(marshaledError)
		return
	}

	response := JsonCreateMeetingResponse{MeetingID: meeting.ID}

	marshaledResponse, err := json.Marshal(response)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	writer.WriteHeader(http.StatusOK)
	writer.Write(marshaledResponse)
}
