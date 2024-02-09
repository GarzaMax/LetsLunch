package query

import (
	"cmd/app/models"
	"database/sql"
	"github.com/gofrs/uuid/v5"
)

type FindCriteria struct {
	OnlyCity    sql.NullString
	Address     *models.Address
	InitiatorID uuid.NullUUID
	CuisineType sql.NullInt16
	Rating      int
	PhoneNumber sql.NullString
	Title       sql.NullString
}
