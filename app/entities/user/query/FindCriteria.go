package query

import (
	"database/sql"
	"github.com/gofrs/uuid/v5"
)

type FindCriteria struct {
	Username         sql.NullString
	DisplayName      sql.NullString
	CurrentMeetingId uuid.NullUUID
	Age              sql.NullInt32
	Gender           sql.NullInt16
	PhoneNumber      sql.NullString
}
