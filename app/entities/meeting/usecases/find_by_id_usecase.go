package usecases

import (
	"cmd/app/entities/meeting"
	"cmd/app/entities/meeting/repository"
	"context"
	"github.com/gofrs/uuid/v5"
	"time"
)

type FindMeetingByIdResponse struct {
	ID               uuid.UUID            `json:"id"`
	GatheringPlaceID uuid.UUID            `json:"gatheringPlaceId"`
	InitiatorsID     uuid.UUID            `json:"initiatorsId"`
	StartTime        time.Time            `json:"startTime"`
	EndTime          time.Time            `json:"endTime"`
	UsersQuantity    int                  `json:"usersQuantity"`
	State            meeting.MeetingState `json:"state"`
}

type FindMeetingByIdUseCase struct {
	meetings repository.MeetingsRepository
}

func NewFindMeetingByIdUseCase(meetings repository.MeetingsRepository) *FindMeetingByIdUseCase {
	return &FindMeetingByIdUseCase{meetings: meetings}
}

func (useCase *FindMeetingByIdUseCase) Handle(ctx context.Context, id uuid.UUID) (*FindMeetingByIdResponse, error) {
	meetings, err := useCase.meetings.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return &FindMeetingByIdResponse{
		ID:               meetings.ID,
		GatheringPlaceID: meetings.GatheringPlaceId,
		InitiatorsID:     meetings.InitiatorsId,
		StartTime:        meetings.StartTime,
		EndTime:          meetings.EndTime,
		UsersQuantity:    meetings.UsersQuantity,
		State:            meetings.State,
	}, nil
}
