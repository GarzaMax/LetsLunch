package usecases

import (
	"context"
	"time"

	domain "cmd/app/entities/user"
	"cmd/app/entities/user/repository"
	"github.com/gofrs/uuid/v5"
)

type FindUserByIdResponse struct {
	ID               uuid.UUID     `json:"id"`
	Username         string        `json:"username"`
	DisplayName      string        `json:"displayName"`
	CurrentMeetingID uuid.NullUUID `json:"currentMeetingId"`
	MeetingHistory   []uuid.UUID   `json:"meetingHistory"`
	Rating           int           `json:"rating"`
	Birthday         time.Time     `json:"birthday"`
	Gender           domain.Gender `json:"gender"`
}

type FindUserByIdUseCase struct {
	users repository.UsersRepository
}

func NewFindUserByIdUseCase(users repository.UsersRepository) *FindUserByIdUseCase {
	return &FindUserByIdUseCase{users: users}
}

func (useCase *FindUserByIdUseCase) Handle(ctx context.Context, id uuid.UUID) (*FindUserByIdResponse, error) {
	users, err := useCase.users.FindUserByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return &FindUserByIdResponse{
		ID:               users.ID,
		Username:         users.Username,
		DisplayName:      users.DisplayName,
		CurrentMeetingID: users.CurrentMeetingId,
		MeetingHistory:   users.MeetingHistory,
		Rating:           users.Rating,
		Birthday:         users.Birthday,
		Gender:           users.Gender,
	}, nil
}
