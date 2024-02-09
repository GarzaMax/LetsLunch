package usecases

import (
	"cmd/app/entities/user"
	domain "cmd/app/entities/user"
	"cmd/app/entities/user/repository"
	"cmd/app/entities/user/validators"
	"context"
	"fmt"
	"time"
)

type CreateUserUseCase struct {
	User repository.UsersRepository
}

type CreateUserCommand struct {
	Username       string
	DisplayName    string
	Birthday       time.Time
	PhoneNumber    string
	Gender         user.Gender
	HashedPassword []byte
}

func NewCreateUserUseCase(user repository.UsersRepository) *CreateUserUseCase {
	return &CreateUserUseCase{User: user}
}

func (useCase *CreateUserUseCase) Handle(
	ctx context.Context,
	command CreateUserCommand,
) (*domain.User, error) {
	user := user.NewUser(
		command.Username,
		command.DisplayName,
		command.Birthday,
		command.PhoneNumber,
		command.Gender,
		command.HashedPassword,
	)
	isUsernameUnique, err := validators.IsUsernameUnique(ctx, user, useCase.User)
	isPhoneUnique, err := validators.IsPhoneNubmerUnique(ctx, user, useCase.User)
	if (isUsernameUnique && isPhoneUnique) == true {
		_, err = useCase.User.Create(ctx, user)
		if err != nil {
			return nil, fmt.Errorf("user: create user %w", err)
		}
		return user, nil
	} else {
		return nil, nil
	}

}
