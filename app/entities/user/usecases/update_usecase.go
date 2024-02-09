package usecases

import (
	"github.com/gofrs/uuid/v5"

	"cmd/app/entities/user"
	"cmd/app/entities/user/repository"
	"context"
	"time"
)

type UpdateUserResponse struct {
	ID uuid.UUID `json:"id"`
}

type UpdateUserUseCase struct {
	user repository.UsersRepository
}

type UpdateUserCommand struct {
	Username       string
	DisplayName    string
	Birthday       time.Time
	PhoneNumber    string
	Gender         user.Gender
	HashedPassword []byte
}

func NewUpdateUserUseCase(user repository.UsersRepository) *UpdateUserUseCase {
	return &UpdateUserUseCase{user: user}
}

func (useCase *UpdateUserUseCase) Handle(
	ctx context.Context,
	command UpdateUserCommand,
	id uuid.UUID,
) (*UpdateUserResponse, error) {
	user, err := useCase.user.FindUserByID(ctx, id)
	if err != nil {
		return nil, err
	}

	user.Username = command.Username
	user.DisplayName = command.DisplayName
	user.Birthday = command.Birthday
	user.PhoneNumber = command.PhoneNumber
	user.Gender = command.Gender
	user.HashedPassword = command.HashedPassword

	_, err = useCase.user.Update(ctx, user)
	if err != nil {
		return nil, err
	}

	return &UpdateUserResponse{ID: id}, nil
}
