package api

import (
	"cmd/app/entities/meeting/query"
	"cmd/app/entities/meeting/usecases"
	"encoding/json"
	"github.com/gofrs/uuid/v5"
	"net/http"
)

type FindMeetingsByCriteriaHandler struct {
	useCase *usecases.FindMeetingsByCriteriaUseCase
}

func NewFindMeetingsByCriteriaHandler(useCase *usecases.FindMeetingsByCriteriaUseCase) *FindMeetingsByCriteriaHandler {
	return &FindMeetingsByCriteriaHandler{useCase: useCase}
}

func (handler *FindMeetingsByCriteriaHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	gatheringPlaceId := request.URL.Query().Get("gathering_place_id")
	initiatorsId := request.URL.Query().Get("initiators_id")

	criteria := query.FindCriteria{}

	if gatheringPlaceId != "" {
		uuidID, err := uuid.FromString(gatheringPlaceId)
		if err != nil {
			marshaledError, _ := json.Marshal(err)

			writer.WriteHeader(http.StatusBadRequest)
			writer.Write(marshaledError)
			return
		}
		criteria.GatheringPlaceId.UUID = uuidID
	}

	if initiatorsId != "" {
		uuidID, err := uuid.FromString(initiatorsId)
		if err != nil {
			marshaledError, _ := json.Marshal(err)

			writer.WriteHeader(http.StatusBadRequest)
			writer.Write(marshaledError)
			return
		}
		criteria.InitiatorID.UUID = uuidID
	}

	meetings, err := handler.useCase.Handle(request.Context(), criteria)

	if err != nil {
		marshaledError, _ := json.Marshal(err)

		writer.WriteHeader(http.StatusInternalServerError)
		writer.Write(marshaledError)
		return
	}

	marshaledResponse, err := json.Marshal(meetings)
	if err != nil {
		marshaledError, _ := json.Marshal(err)

		writer.WriteHeader(http.StatusInternalServerError)
		writer.Write(marshaledError)
		return
	}

	writer.WriteHeader(http.StatusOK)
	writer.Write(marshaledResponse)
}
