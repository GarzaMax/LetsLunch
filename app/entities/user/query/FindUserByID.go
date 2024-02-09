package query

import (
	"context"
	"database/sql"
	"github.com/gofrs/uuid/v5"
)

func FindUserByID(ctx context.Context, id uuid.UUID, db *sql.DB) *sql.Row {
	const query = `SELECT id, username, display_name, rating, gender, current_meeting_id, phone_number, birthday FROM users WHERE id = $1`
	row := db.QueryRow(query, id)
	return row
}
