package query

import (
	"cmd/app/entities/user"
	"context"
	"database/sql"
	"fmt"
)

func Update(ctx context.Context, user *user.User, db *sql.DB) error {
	const query = `UPDATE users 
    SET id = $1, username = $2, display_name = $3, rating = $4, current_meeting_id = $5, gender = $6, phone_number = $7, birthday = $8, hashed_password = $9 WHERE id = $1`
	_, err := db.Exec(query, user.ID, user.Username, user.DisplayName, user.Rating, user.CurrentMeetingId, user.Gender, user.PhoneNumber, user.Birthday, user.HashedPassword)
	if err != nil {
		return fmt.Errorf("database query execution error: %w", err)
	}
	return nil
}
