package repositories

import (
	"cmd/app/entities/gatheringPlace"
	"cmd/app/entities/gatheringPlace/query"
	repositoryPlaces "cmd/app/entities/gatheringPlace/repository"
	"cmd/app/models"
	"context"
	"database/sql"
	"errors"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/assert"
	"slices"
	"testing"
)

var testAddress = models.Address{Country: "Russia", City: "Saint-Petersburg",
	StreetName: "Kronverkskiy prospekt", HouseNumber: "", BuildingNumber: 49}

func TestMain(m *testing.M) {
	setUpConnection()
	_ = m.Run()
	closeConnection()
}
func TestCreatingGatheringPlace(t *testing.T) {
	//set up
	var databasePlacesRepository = repositoryPlaces.NewPlacesDatabaseRepository(db)
	var currentPlace = gatheringPlace.NewGatheringPlace(testAddress, 500, gatheringPlace.FastFood, 5, "+781245422005", "", "", "")
	var ctx = context.Background()

	//main part
	_, errCreating := databasePlacesRepository.Create(ctx, currentPlace)
	if errCreating != nil {
		t.Fatalf("Error in creating place: %v", errCreating)
	}
	placeFromRepository, errFinding := databasePlacesRepository.FindByID(ctx, currentPlace.ID)
	if errFinding != nil {
		t.Fatalf("Error in finding place: %v", errFinding)
	}

	//testing equality of found and created entities
	assert.Equal(t, currentPlace, placeFromRepository)
	errDeleting := databasePlacesRepository.Delete(ctx, currentPlace)
	if errDeleting != nil {
		t.Fatalf("Error in deleting place: %v", errDeleting)
	}
}

func TestFindingByCriteriaGatheringPlace(t *testing.T) {
	//set up
	var databasePlacesRepository = repositoryPlaces.NewPlacesDatabaseRepository(db)
	var firstPlace = gatheringPlace.NewGatheringPlace(testAddress, 500, gatheringPlace.FastFood, 5, "+781245422005", "", "", "")
	var ctx = context.Background()
	_, errCreating := databasePlacesRepository.Create(ctx, firstPlace)
	if errCreating != nil {
		t.Fatalf("Error in creating place: %v", errCreating)
	}
	var secondPlace = gatheringPlace.NewGatheringPlace(testAddress, 500, gatheringPlace.Eastern, 5, "+781245422005", "", "", "")
	_, errCreating = databasePlacesRepository.Create(ctx, secondPlace)
	if errCreating != nil {
		t.Fatalf("Error in creating place: %v", errCreating)
	}
	var thirdPlace = gatheringPlace.NewGatheringPlace(testAddress, 500, gatheringPlace.FastFood, 5, "+781245422005", "", "", "")
	_, errCreating = databasePlacesRepository.Create(ctx, thirdPlace)
	if errCreating != nil {
		t.Fatalf("Error in creating place: %v", errCreating)
	}

	//main part
	var findingCriteria = query.FindCriteria{CuisineType: sql.NullInt16{Int16: gatheringPlace.FastFood, Valid: true}}
	placesWithFastFood, errFinding := databasePlacesRepository.FindByCriteria(ctx, findingCriteria)
	if errFinding != nil {
		t.Fatalf("Error in finding place: %v", errFinding)
	}
	findingCriteria.CuisineType = sql.NullInt16{Int16: gatheringPlace.Eastern, Valid: true}
	placesWithEastern, errFinding := databasePlacesRepository.FindByCriteria(ctx, findingCriteria)
	if errFinding != nil {
		t.Fatalf("Error in finding place: %v", errFinding)
	}

	//testing
	// first - check, if a placesWithFastFood contains right restaurants
	assert.True(t, slices.Contains(placesWithFastFood, *firstPlace))
	assert.True(t, slices.Contains(placesWithFastFood, *thirdPlace))
	assert.False(t, slices.Contains(placesWithFastFood, *secondPlace))
	// second - check, if a placesWithEastern contains right restaurants
	assert.True(t, slices.Contains(placesWithEastern, *secondPlace))
	assert.False(t, slices.Contains(placesWithEastern, *firstPlace))
	errDeleting := databasePlacesRepository.Delete(ctx, firstPlace)
	errDeleting = databasePlacesRepository.Delete(ctx, secondPlace)
	errDeleting = databasePlacesRepository.Delete(ctx, thirdPlace)
	if errDeleting != nil {
		t.Fatalf("Error in deleting place: %v", errDeleting)
	}
}

func TestUpdatingGatheringPlace(t *testing.T) {
	//set up
	var databasePlacesRepository = repositoryPlaces.NewPlacesDatabaseRepository(db)
	var firstPlace = gatheringPlace.NewGatheringPlace(testAddress, 500, gatheringPlace.FastFood, 5, "+781245422005", "", "", "")
	var ctx = context.Background()
	_, errCreating := databasePlacesRepository.Create(ctx, firstPlace)
	if errCreating != nil {
		t.Fatalf("Error in creating place: %v", errCreating)
	}
	//main part
	firstPlace.Address = models.Address{Country: "Romania", City: "Bucharest"}
	_, errUpdating := databasePlacesRepository.Update(ctx, firstPlace)
	if errUpdating != nil {
		t.Fatalf("Error in updating place: %v", errUpdating)
	}
	placeFromRepository, errFinding := databasePlacesRepository.FindByID(ctx, firstPlace.ID)
	if errFinding != nil {
		t.Fatalf("Error in finding place: %v", errFinding)
	}

	//testing
	//checking the equality of found and updated entity
	assert.Equal(t, firstPlace, placeFromRepository)
	errDeleting := databasePlacesRepository.Delete(ctx, firstPlace)
	if errDeleting != nil {
		t.Fatalf("Error in deleting place: %v", errDeleting)
	}
}

func TestDeletingGatheringPlace(t *testing.T) {
	//set up
	var databasePlacesRepository = repositoryPlaces.NewPlacesDatabaseRepository(db)
	var firstPlace = gatheringPlace.NewGatheringPlace(testAddress, 500, gatheringPlace.FastFood, 5, "+781245422005", "", "", "")
	var ctx = context.Background()
	_, errCreating := databasePlacesRepository.Create(ctx, firstPlace)
	if errCreating != nil {
		t.Fatalf("Error in creating place: %v", errCreating)
	}
	//main part
	errDeleting := databasePlacesRepository.Delete(ctx, firstPlace)
	if errDeleting != nil {
		t.Fatalf("Error in deleting place: %v", errDeleting)
	}
	_, errFinding := databasePlacesRepository.FindByID(ctx, firstPlace.ID)
	//testing
	//checking the right error - no rows with such ID remained
	assert.Contains(t, errFinding.Error(), sql.ErrNoRows.Error())
	assert.Equal(t, "no such gathering place: sql: no rows in result set", errFinding.Error())
	assert.True(t, errors.Is(errFinding, sql.ErrNoRows))
}
