package api

import (
	"cmd/app/entities/user"
	"cmd/app/entities/user/usecases"
	"cmd/app/models"
	"encoding/json"
	"github.com/gofrs/uuid/v5"
	"net/http"
	"strconv"
)

type FindUsersByCriteriaHandler struct {
	UseCase *usecases.FindUsersByCriteriaUseCase
}

func NewFindUsersByCriteriaHandler(useCase *usecases.FindUsersByCriteriaUseCase) *FindUsersByCriteriaHandler {
	return &FindUsersByCriteriaHandler{UseCase: useCase}
}

func (handler *FindUsersByCriteriaHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	username := request.URL.Query().Get("user_name")
	displayName := request.URL.Query().Get("dislay_name")
	currentMeetingID := request.URL.Query().Get("current_meeting_id")
	age := request.URL.Query().Get("age")
	gender := request.URL.Query().Get("gender")
	phoneNumber := request.URL.Query().Get("phone_number")

	findCriteria := models.FindUserCriteria{}
	findCriteria.Username = username
	findCriteria.DisplayName = displayName
	findCriteria.PhoneNumber = phoneNumber

	if currentMeetingID != "" {
		uuidID, err := uuid.FromString(currentMeetingID)
		if err != nil {
			marshaledError, _ := json.Marshal(err)

			writer.WriteHeader(http.StatusBadRequest)
			writer.Write(marshaledError)
			return
		}
		findCriteria.CurrentMeetingId = uuidID
	}

	if age != "" {
		convertedAge, err := strconv.Atoi(age)
		if err != nil {
			marshaledError, _ := json.Marshal(err)

			writer.WriteHeader(http.StatusBadRequest)
			writer.Write(marshaledError)
			return
		}
		findCriteria.Age = convertedAge
	}

	if gender != "" {
		convertedGender, err := strconv.Atoi(gender)
		if err != nil {
			marshaledError, _ := json.Marshal(err)

			writer.WriteHeader(http.StatusBadRequest)
			writer.Write(marshaledError)
			return
		}

		actualGender := user.Gender(convertedGender)

		findCriteria.Gender = &actualGender
	}

	users, err := handler.UseCase.Handle(request.Context(), findCriteria)
	if err != nil {
		marshaledError, _ := json.Marshal(err)

		writer.WriteHeader(http.StatusInternalServerError)
		writer.Write(marshaledError)
		return
	}

	marshaledResponse, err := json.Marshal(users)
	if err != nil {
		marshaledError, _ := json.Marshal(err)

		writer.WriteHeader(http.StatusInternalServerError)
		writer.Write(marshaledError)
		return
	}

	writer.WriteHeader(http.StatusOK)
	writer.Write(marshaledResponse)
}
