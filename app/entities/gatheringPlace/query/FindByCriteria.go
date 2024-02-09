package query

import (
	"cmd/app/models"
	"context"
	"database/sql"
	"fmt"
	sq "github.com/Masterminds/squirrel"
	"github.com/gofrs/uuid/v5"
)

func FindByCriteria(ctx context.Context, criteria FindCriteria, db *sql.DB) (*sql.Rows, error) {
	psql := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)

	sqlStatement := psql.Select("*").From("gathering_places").RunWith(db)
	if criteria.Address != nil {
		sqlStatement = findByAddress(sqlStatement, criteria.Address)
	}
	if criteria.OnlyCity.Valid != false {
		sqlStatement = findByCity(sqlStatement, criteria.OnlyCity.String)
	}
	if criteria.InitiatorID.Valid != false {
		sqlStatement = findByInitiatorID(sqlStatement, criteria.InitiatorID.UUID)
	}
	if criteria.Rating != 0 {
		sqlStatement = findByRating(sqlStatement, criteria.Rating)
	}
	if criteria.CuisineType.Valid != false {
		sqlStatement = findByCuisineType(sqlStatement, int(criteria.CuisineType.Int16))
	}
	if criteria.Title.Valid != false {
		sqlStatement = findByTitle(sqlStatement, criteria.Title.String)
	}
	var rows, err = sqlStatement.Query()
	if err != nil {
		return nil, fmt.Errorf("problem with quering to database %w", err)
	}
	return rows, nil
}

func findByInitiatorID(sql sq.SelectBuilder, id uuid.UUID) sq.SelectBuilder {
	return sql.Where(sq.Eq{"initiators_id": id})
}
func findByRating(sql sq.SelectBuilder, rating int) sq.SelectBuilder {
	return sql.Where(sq.Eq{"rating": rating})
}
func findByAddress(sql sq.SelectBuilder, address *models.Address) sq.SelectBuilder {
	return sql.Where(sq.Eq{"country": address.Country, "city": address.City,
		"street_name": address.StreetName, "house_number": address.HouseNumber,
		"building_number": address.BuildingNumber})
}
func findByCuisineType(sql sq.SelectBuilder, cuisineType int) sq.SelectBuilder {
	return sql.Where(sq.Eq{"cuisine_type": cuisineType})
}
func findByTitle(sql sq.SelectBuilder, title string) sq.SelectBuilder {
	return sql.Where(sq.Eq{"title": title})
}
func findByCity(sql sq.SelectBuilder, city string) sq.SelectBuilder {
	return sql.Where(sq.Eq{"city": city})
}
