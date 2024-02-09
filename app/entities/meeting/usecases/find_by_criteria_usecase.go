package usecases

import (
	"cmd/app/entities/meeting"
	"cmd/app/entities/meeting/query"
	"cmd/app/entities/meeting/repository"
	"context"
	"github.com/gofrs/uuid/v5"
)

type FindMeetingsByCriteriaUseCase struct {
	meeting repository.MeetingsRepository
}

func NewFindMeetingsByCriteriaUseCase(meetings repository.MeetingsRepository) *FindMeetingsByCriteriaUseCase {
	return &FindMeetingsByCriteriaUseCase{meeting: meetings}
}

func (useCase *FindMeetingsByCriteriaUseCase) Handle(
	ctx context.Context,
	criteria query.FindCriteria,
) ([]meeting.Meeting, error) {

	if criteria.InitiatorID.UUID != uuid.Nil {
		criteria.InitiatorID.Valid = true
	}
	if criteria.GatheringPlaceId.UUID != uuid.Nil {
		criteria.GatheringPlaceId.Valid = true
	}

	queryResult, err := useCase.meeting.FindByCriteria(ctx, criteria)

	if err != nil {
		return nil, err
	}
	if queryResult == nil {
		return []meeting.Meeting{}, nil
	}
	return queryResult, nil
}
