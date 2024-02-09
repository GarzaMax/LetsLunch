package validators

import (
	"cmd/app/entities/user"
	"cmd/app/entities/user/query"
	"cmd/app/entities/user/repository"
	"context"
	"fmt"
)

func IsPhoneNubmerUnique(ctx context.Context, user *user.User, repository repository.UsersRepository) (isValid bool, err error) {
	criteria := query.FindCriteria{}
	criteria.PhoneNumber.String = user.PhoneNumber
	criteria.PhoneNumber.Valid = true

	users, err := repository.FindUsersByCriteria(ctx, criteria)
	if err != nil {
		return false, fmt.Errorf("user: find user %w", err)
	}

	if len(users) == 0 {
		return true, nil
	} else {
		return false, nil
	}
}
