package query

import (
	"cmd/app/entities/user"
	"context"
	"database/sql"
	"fmt"
)

func Update(ctx context.Context, user *user.User, db *sql.DB) error {
	const query = `UPDATE users 
    SET id = $1, username = $2, display_name = $3, rating = $4, current_meeting_id = $5, age = $6, gender = $7 WHERE id = $1`
	_, err := db.Exec(query, user.ID, user.Username, user.DisplayName, user.Rating, user.CurrentMeetingId, user.Age, user.Gender)
	if err != nil {
		return fmt.Errorf("database query execution error: %w", err)
	}
	return nil
}

/*func Update(ctx context.Context, user *user.User, db *sql.DB) error {
	const query = `UPDATE meetings SET current_meeting_id = $1 WHERE id = $2`
	_, err := db.Exec(query, user.CurrentMeetingId, user.ID)
	if err != nil {
		return fmt.Errorf("database query execution error: %w", err)
	}
	return nil
}*/
