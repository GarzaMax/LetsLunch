package api

import (
	"cmd/app/entities/user/dto"
	"cmd/app/entities/user/usecases"
	"encoding/json"
	"github.com/gofrs/uuid/v5"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

type JsonCreateUserResponse struct {
	UserUUID uuid.UUID `json:"id"`
}

type CreateUserHandler struct {
	UseCase *usecases.CreateUserUseCase
}

func NewCreateUserHandler(useCase *usecases.CreateUserUseCase) *CreateUserHandler {
	return &CreateUserHandler{UseCase: useCase}
}

func (handler *CreateUserHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	var createUserDto dto.CreateUserDto
	if err := json.NewDecoder(request.Body).Decode(&createUserDto); err != nil {
		marshaledError, _ := json.Marshal(err.Error())

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
	var err error
	command.HashedPassword, err = bcrypt.GenerateFromPassword([]byte(createUserDto.Password), 8)
	if err != nil {
		marshaledError, _ := json.Marshal(err.Error())

		writer.WriteHeader(http.StatusInternalServerError)
		writer.Write(marshaledError)
		return
	}

	user, err := handler.UseCase.Handle(request.Context(), command)

	if err != nil {
		marshaledError, _ := json.Marshal(err.Error())

		writer.WriteHeader(http.StatusInternalServerError)
		writer.Write(marshaledError)
		return
	}

	if user == nil {
		writer.WriteHeader(http.StatusUnprocessableEntity)
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
