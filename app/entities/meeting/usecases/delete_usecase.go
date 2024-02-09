package usecases

import (
	"cmd/app/entities/meeting/repository"
	"context"
	"github.com/gofrs/uuid/v5"
)

type DeleteMeetingUseCase struct {
	meeting repository.MeetingsRepository
}

func NewDeleteMeetingUseCase(meeting repository.MeetingsRepository) *DeleteMeetingUseCase {
	return &DeleteMeetingUseCase{meeting: meeting}
}

func (useCase *DeleteMeetingUseCase) Handle(
	ctx context.Context,
	id uuid.UUID,
) error {
	meeting, err := useCase.meeting.FindByID(ctx, id)
	if err != nil {
		return nil
	}

	err = useCase.meeting.Delete(ctx, meeting)
	return err
}
