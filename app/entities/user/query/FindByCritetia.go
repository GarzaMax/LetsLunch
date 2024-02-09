package query

import (
	"context"
	"database/sql"
	"fmt"
	sq "github.com/Masterminds/squirrel"
	"github.com/gofrs/uuid/v5"
)

func FindUserByCriteria(ctx context.Context, criteria FindCriteria, db *sql.DB) (*sql.Rows, error) {
	psql := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)

	sqlStatement := psql.Select("*").From("users").RunWith(db)
	if criteria.Username.Valid != false {
		sqlStatement = findUserByUsername(sqlStatement, criteria.Username.String)
	}
	if criteria.DisplayName.Valid != false {
		sqlStatement = findUserByDisplayname(sqlStatement, criteria.DisplayName.String)
	}
	if criteria.CurrentMeetingId.Valid != false {
		sqlStatement = findUserByCurrentMeetingId(sqlStatement, criteria.CurrentMeetingId.UUID)
	}
	if criteria.Age.Valid != false {
		sqlStatement = findUserByAge(sqlStatement, int(criteria.Age.Int32))
	}
	if criteria.Gender.Valid != false {
		sqlStatement = findUserByGender(sqlStatement, int(criteria.Gender.Int16))
	}
	if criteria.PhoneNumber.Valid != false {
		sqlStatement = findUserByPhoneNumber(sqlStatement, criteria.PhoneNumber.String)
	}
	var rows, err = sqlStatement.Query()
	if err != nil {
		return nil, fmt.Errorf("database query execution error: %w", err)
	}
	return rows, nil
}
func findUserByCurrentMeetingId(sql sq.SelectBuilder, id uuid.UUID) sq.SelectBuilder {
	return sql.Where(sq.Eq{"current_meeting_id": id})
}
func findUserByUsername(sql sq.SelectBuilder, username string) sq.SelectBuilder {
	return sql.Where(sq.Eq{"username": username})
}
func findUserByDisplayname(sql sq.SelectBuilder, displayname string) sq.SelectBuilder {
	return sql.Where(sq.Eq{"display_name": displayname})
}
func findUserByAge(sql sq.SelectBuilder, age int) sq.SelectBuilder {
	return sql.Where(sq.Eq{"AGE(CURRENT_DATA, birthday)": age})
	// вот здесь как раз можно дельту прикрутить для возраста
}
func findUserByGender(sql sq.SelectBuilder, gender int) sq.SelectBuilder {
	return sql.Where(sq.Eq{"gender": gender})
}

func findUserByPhoneNumber(sql sq.SelectBuilder, phoneNumber string) sq.SelectBuilder {
	return sql.Where(sq.Eq{"phone_number": phoneNumber})
	// вот здесь как раз можно дельту прикрутить для возраста
}
func FindUserHistoryById(ctx context.Context, id uuid.UUID, db *sql.DB) (*sql.Rows, error) {
	const query = "SELECT meeting_id FROM meetings_history WHERE user_id = $1"
	rows, err := db.Query(query, id)
	if err != nil {
		return nil, fmt.Errorf("database query execution error: %w", err)
	}
	return rows, nil
}
