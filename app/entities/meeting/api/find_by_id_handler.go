package api

import (
	domain "cmd/app/entities/meeting"
	usecase "cmd/app/entities/meeting/usecases"
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"github.com/gofrs/uuid/v5"
	"net/http"
	"time"
)

type JsonFindMeetingByIdResponse struct {
	ID               uuid.UUID           `json:"id"`
	GatheringPlaceID uuid.UUID           `json:"gatheringPlaceId"`
	InitiatorsID     uuid.UUID           `json:"initiatorsId"`
	StartTime        time.Time           `json:"startTime"`
	EndTime          time.Time           `json:"endTime"`
	UsersQuantity    int                 `json:"usersQuantity"`
	State            domain.MeetingState `json:"state"`
}

type FindMeetingByIdHandler struct {
	useCase *usecase.FindMeetingByIdUseCase
}

func NewFindMeetingByIdHandler(useCase *usecase.FindMeetingByIdUseCase) *FindMeetingByIdHandler {
	return &FindMeetingByIdHandler{useCase: useCase}
}

func (handler *FindMeetingByIdHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	id := chi.URLParam(request, "meetingID")

	uuidID, err := uuid.FromString(id)
	if err != nil {
		marshaledError, _ := json.Marshal(err.Error())

		writer.WriteHeader(http.StatusBadRequest)
		writer.Write(marshaledError)
		return
	}

	meeting, err := handler.useCase.Handle(request.Context(), uuidID)
	if err != nil {
		marshaledError, _ := json.Marshal(err.Error())

		writer.WriteHeader(http.StatusInternalServerError)
		writer.Write(marshaledError)
		return
	}

	response := JsonFindMeetingByIdResponse{
		ID:               meeting.ID,
		GatheringPlaceID: meeting.GatheringPlaceID,
		InitiatorsID:     meeting.InitiatorsID,
		StartTime:        meeting.StartTime,
		EndTime:          meeting.EndTime,
		UsersQuantity:    meeting.UsersQuantity,
		State:            meeting.State,
	}

	marshaledResponse, err := json.Marshal(response)
	if err != nil {
		marshaledError, _ := json.Marshal(err.Error())

		writer.WriteHeader(http.StatusInternalServerError)
		writer.Write(marshaledError)
		return
	}

	writer.WriteHeader(http.StatusOK)
	writer.Write(marshaledResponse)
}
