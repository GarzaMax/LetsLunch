package usecases

import (
	"cmd/app/entities/user"
	"cmd/app/entities/user/query"
	"cmd/app/entities/user/repository"
	"cmd/app/models"
	"context"
	"github.com/gofrs/uuid/v5"
)

type FindUsersByCriteriaUseCase struct {
	user repository.UsersRepository
}

func NewFindUsersByCriteriaUseCase(users repository.UsersRepository) *FindUsersByCriteriaUseCase {
	return &FindUsersByCriteriaUseCase{user: users}
}

func (useCase *FindUsersByCriteriaUseCase) Handle(
	ctx context.Context,
	criteria models.FindUserCriteria,
) ([]user.User, error) {
	queryCriteria := query.FindCriteria{}

	queryCriteria.Username.String = criteria.Username
	queryCriteria.Username.Valid = true
	queryCriteria.DisplayName.String = criteria.DisplayName
	queryCriteria.DisplayName.Valid = true
	queryCriteria.CurrentMeetingId.UUID = criteria.CurrentMeetingId
	queryCriteria.CurrentMeetingId.Valid = true
	queryCriteria.Age.Int32 = int32(criteria.Age)
	queryCriteria.Age.Valid = true
	queryCriteria.PhoneNumber.String = criteria.PhoneNumber
	queryCriteria.PhoneNumber.Valid = true

	if criteria.Username == "" {
		queryCriteria.Username.Valid = false
	}
	if criteria.DisplayName == "" {
		queryCriteria.DisplayName.Valid = false
	}
	if criteria.CurrentMeetingId == uuid.Nil {
		queryCriteria.CurrentMeetingId.Valid = false
	}
	if criteria.Age == 0 {
		queryCriteria.Age.Valid = false
	}
	if criteria.Gender == nil {
		queryCriteria.Gender.Valid = false
	} else {
		queryCriteria.Gender.Int16 = int16(*criteria.Gender)
		queryCriteria.Gender.Valid = true
	}
	if criteria.PhoneNumber == "" {
		queryCriteria.Username.Valid = false
	}

	queryResult, err := useCase.user.FindUsersByCriteria(ctx, queryCriteria)

	if err != nil {
		return nil, err
	}
	if queryResult == nil {
		return []user.User{}, nil
	}
	return queryResult, nil
}
