package usecases

import (
	"cmd/app/entities/meeting"
	"cmd/app/entities/meeting/repository"
	"context"
	"fmt"
	"github.com/gofrs/uuid/v5"
	"time"
)

type CreateMeetingUseCase struct {
	meetings repository.MeetingsRepository
}

type CreateMeetingCommand struct {
	GatheringPlaceID uuid.UUID
	InitiatorsID     uuid.UUID
	StartTime        time.Time
	EndTime          time.Time
	UsersQuantity    int
	State            meeting.MeetingState
}

func NewCreateMeetingUseCase(meetings repository.MeetingsRepository) *CreateMeetingUseCase {
	return &CreateMeetingUseCase{meetings: meetings}
}

func (useCase *CreateMeetingUseCase) Handle(
	ctx context.Context,
	command CreateMeetingCommand,
) (*meeting.Meeting, error) {
	meeting := meeting.NewMeeting(
		command.GatheringPlaceID,
		command.InitiatorsID,
		command.StartTime,
		command.EndTime,
		command.UsersQuantity,
		command.State,
	)

	_, err := useCase.meetings.Create(ctx, meeting)
	if err != nil {
		return nil, fmt.Errorf("meeting: create meeting %w", err)
	}

	return meeting, nil
}
