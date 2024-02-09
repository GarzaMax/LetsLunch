package usecases

import (
	"cmd/app/entities/meeting"
	"cmd/app/entities/meeting/repository"
	"context"
	"github.com/gofrs/uuid/v5"
	"time"
)

type UpdateMeetingResponse struct {
	ID uuid.UUID `json:"id"`
}

type UpdateMeetingUseCase struct {
	meeting repository.MeetingsRepository
}

type UpdateMeetingCommand struct {
	GatheringPlaceID uuid.UUID
	InitiatorsID     uuid.UUID
	StartTime        time.Time
	EndTime          time.Time
	UsersQuantity    int
	State            meeting.MeetingState
}

func NewUpdateMeetingUseCase(meeting repository.MeetingsRepository) *UpdateMeetingUseCase {
	return &UpdateMeetingUseCase{meeting: meeting}
}

func (useCase *UpdateMeetingUseCase) Handle(
	ctx context.Context,
	command UpdateMeetingCommand,
	id uuid.UUID,
) (*UpdateMeetingResponse, error) {
	meeting, err := useCase.meeting.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}

	meeting.GatheringPlaceId = command.GatheringPlaceID
	meeting.InitiatorsId = command.InitiatorsID
	meeting.StartTime = command.StartTime
	meeting.EndTime = command.EndTime
	meeting.UsersQuantity = command.UsersQuantity
	meeting.State = command.State

	_, err = useCase.meeting.Update(ctx, meeting)
	if err != nil {
		return nil, err
	}

	return &UpdateMeetingResponse{ID: id}, nil
}
