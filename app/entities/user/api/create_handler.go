package api

import (
	"cmd/app/entities/user/dto"
	"cmd/app/entities/user/usecases"
	"encoding/json"
	"github.com/gofrs/uuid/v5"
	"net/http"
)

type JsonCreateUserResponse struct {
	UserUUID uuid.UUID `json:"id"`
}

type CreateUserHandler struct {
	useCase *usecases.CreateUserUseCase
}

func NewCreateUserHandler(useCase *usecases.CreateUserUseCase) *CreateUserHandler {
	return &CreateUserHandler{useCase: useCase}
}

func (handler *CreateUserHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	var createUserDto dto.CreateUserDto
	if err := json.NewDecoder(request.Body).Decode(&createUserDto); err != nil {
		marshaledError, _ := json.Marshal(err)

		writer.WriteHeader(http.StatusBadRequest)
		writer.Write(marshaledError)
		return
	}

	command := usecases.CreateUserCommand{}
	command.Username = createUserDto.Username
	command.DisplayName = createUserDto.DisplayName
	command.Birthday = createUserDto.Birthday
	command.PhoneNumber = createUserDto.PhoneNumber
	command.Gender = createUserDto.Gender

	user, err := handler.useCase.Handle(request.Context(), command)

	if err != nil {
		marshaledError, _ := json.Marshal(err)

		writer.WriteHeader(http.StatusInternalServerError)
		writer.Write(marshaledError)
		return
	}

	response := JsonCreateUserResponse{UserUUID: user.ID}

	marshaledResponse, err := json.Marshal(response)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	writer.WriteHeader(http.StatusOK)
	writer.Write(marshaledResponse)
}