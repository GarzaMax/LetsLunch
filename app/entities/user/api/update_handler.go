package api

import (
	"cmd/app/entities/user/dto"
	"cmd/app/entities/user/usecases"
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"github.com/gofrs/uuid/v5"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

type JsonUpdateUserResponse struct {
	ID uuid.UUID `json:"id"`
}

type UpdateUserHandler struct {
	UseCase *usecases.UpdateUserUseCase
}

func NewUpdateUserHandler(useCase *usecases.UpdateUserUseCase) *UpdateUserHandler {
	return &UpdateUserHandler{UseCase: useCase}
}

func (handler *UpdateUserHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	var updateUserDto dto.UpdateUserDto
	if err := json.NewDecoder(request.Body).Decode(&updateUserDto); err != nil {
		marshaledError, _ := json.Marshal(err.Error())

		writer.WriteHeader(http.StatusBadRequest)
		writer.Write(marshaledError)
		return
	}

	command := usecases.UpdateUserCommand{}
	command.Username = updateUserDto.Username
	command.DisplayName = updateUserDto.DisplayName
	command.Birthday = updateUserDto.Birthday
	command.PhoneNumber = updateUserDto.PhoneNumber
	command.Gender = updateUserDto.Gender
	var err error
	command.HashedPassword, err = bcrypt.GenerateFromPassword([]byte(updateUserDto.Password), 8)
	if err != nil {
		marshaledError, _ := json.Marshal(err.Error())

		writer.WriteHeader(http.StatusInternalServerError)
		writer.Write(marshaledError)
		return
	}
	id := chi.URLParam(request, "userID")

	uuidID, err := uuid.FromString(id)
	if err != nil {
		marshaledError, _ := json.Marshal(err.Error())

		writer.WriteHeader(http.StatusBadRequest)
		writer.Write(marshaledError)
		return
	}

	response, err := handler.UseCase.Handle(request.Context(), command, uuidID)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}
	marshaledResponse, err := json.Marshal(response)

	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	writer.WriteHeader(http.StatusOK)
	writer.Write(marshaledResponse)
}
