package usecases

import (
	"cmd/app/entities/user/repository"
	"context"
	"github.com/gofrs/uuid/v5"
)

type DeleteUserUseCase struct {
	user repository.UsersRepository
}

func NewDeleteUserUseCase(user repository.UsersRepository) *DeleteUserUseCase {
	return &DeleteUserUseCase{user: user}
}

func (useCase *DeleteUserUseCase) Handle(
	ctx context.Context,
	id uuid.UUID,
) error {
	user, err := useCase.user.FindUserByID(ctx, id)
	if err != nil {
		return err
	}

	err = useCase.user.Delete(ctx, user)
	return err
}
