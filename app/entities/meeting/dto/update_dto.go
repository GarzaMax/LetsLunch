package dto

import (
	"cmd/app/entities/meeting"
	"github.com/gofrs/uuid/v5"
	"time"
)

type UpdateMeetingDto struct {
	GatheringPlaceID uuid.UUID            `json:"gatheringPlaceId"`
	InitiatorsID     uuid.UUID            `json:"initiatorsId"`
	StartTime        time.Time            `json:"startTime"`
	EndTime          time.Time            `json:"endTime"`
	UsersQuantity    int                  `json:"usersQuantity"`
	State            meeting.MeetingState `json:"state"`
}
