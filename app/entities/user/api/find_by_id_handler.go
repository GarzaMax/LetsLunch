package api

import (
	domain "cmd/app/entities/user"
	usecases "cmd/app/entities/user/usecases"
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"github.com/gofrs/uuid/v5"
	"net/http"
	"time"
)

type JsonFindUserByIdResponse struct {
	ID               uuid.UUID     `json:"id"`
	Username         string        `json:"username"`
	DisplayName      string        `json:"displayName"`
	CurrentMeetingID uuid.NullUUID `json:"currentMeetingId"`
	MeetingHistory   []uuid.UUID   `json:"meetingHistory"`
	Rating           int           `json:"rating"`
	Birthday         time.Time     `json:"birthday"`
	Gender           domain.Gender `json:"gender"`
}

type FindUserByIdHandler struct {
	UseCase *usecases.FindUserByIdUseCase
}

func NewFindUserByIdHandler(useCase *usecases.FindUserByIdUseCase) *FindUserByIdHandler {
	return &FindUserByIdHandler{UseCase: useCase}
}

func (handler *FindUserByIdHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	id := chi.URLParam(request, "userID")

	uuidID, err := uuid.FromString(id)
	if err != nil {
		marshaledError, _ := json.Marshal(err.Error())

		writer.WriteHeader(http.StatusBadRequest)
		writer.Write(marshaledError)
		return
	}

	users, err := handler.UseCase.Handle(request.Context(), uuidID)
	if err != nil {
		marshaledError, _ := json.Marshal(err.Error())

		writer.WriteHeader(http.StatusInternalServerError)
		writer.Write(marshaledError)
		return
	}

	response := JsonFindUserByIdResponse{
		ID:               users.ID,
		Username:         users.Username,
		DisplayName:      users.DisplayName,
		CurrentMeetingID: users.CurrentMeetingID,
		MeetingHistory:   users.MeetingHistory,
		Rating:           users.Rating,
		Birthday:         users.Birthday,
		Gender:           users.Gender,
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
