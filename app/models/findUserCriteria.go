package models

import (
	"cmd/app/entities/user"
	"github.com/gofrs/uuid/v5"
)

type FindUserCriteria struct {
	Username         string
	DisplayName      string
	CurrentMeetingId uuid.UUID
	Age              int
	Gender           *user.Gender
	PhoneNumber      string
}
