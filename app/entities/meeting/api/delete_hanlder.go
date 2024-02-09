package api

import (
	"cmd/app/entities/meeting/usecases"
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"github.com/gofrs/uuid/v5"
	"net/http"
)

type DeleteMeetingHandler struct {
	useCase *usecases.DeleteMeetingUseCase
}

func NewDeleteMeetingHandler(useCase *usecases.DeleteMeetingUseCase) *DeleteMeetingHandler {
	return &DeleteMeetingHandler{useCase: useCase}
}

func (handler *DeleteMeetingHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	id := chi.URLParam(request, "meetingID")

	uuidID, err := uuid.FromString(id)
	if err != nil {
		marshaledError, _ := json.Marshal(err.Error())

		writer.WriteHeader(http.StatusBadRequest)
		writer.Write(marshaledError)
		return
	}

	err = handler.useCase.Handle(request.Context(), uuidID)
	if err != nil {
		marshaledError, _ := json.Marshal(err.Error())

		writer.WriteHeader(http.StatusInternalServerError)
		writer.Write(marshaledError)
		return
	}

	writer.WriteHeader(http.StatusOK)
}
