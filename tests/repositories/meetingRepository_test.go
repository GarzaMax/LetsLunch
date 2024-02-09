package repositories

import (
	"cmd/app/entities/gatheringPlace"
	repositoryPlaces "cmd/app/entities/gatheringPlace/repository"
	"cmd/app/entities/meeting"
	"cmd/app/entities/meeting/query"
	repositoryMeetings "cmd/app/entities/meeting/repository"
	"cmd/app/entities/user"
	"cmd/app/entities/user/repository"
	"context"
	"database/sql"
	"errors"
	"github.com/gofrs/uuid/v5"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/assert"
	"reflect"
	"slices"
	"testing"
	"time"
)

func setUpTestUser(t *testing.T, ctx context.Context) *user.User {
	date, _ := time.Parse(time.DateOnly, "2003-04-16")
	var testUser = user.NewUser("@testUser", "Steve13", date, "+79528123333", user.Male, []byte("1234567890"))
	var databaseUsersRepository = repository.NewUsersDatabaseRepository(db)
	_, err := databaseUsersRepository.Create(ctx, testUser)
	if err != nil {
		t.Fatalf("Error creating user: %v", err)
	}
	return testUser
}

func setUpTestPlace(t *testing.T, ctx context.Context) *gatheringPlace.GatheringPlace {
	var testPlace = gatheringPlace.NewGatheringPlace(testAddress, 500, gatheringPlace.FastFood, 5, "+781245422005", "", "", "")
	var databasePlacesRepository = repositoryPlaces.NewPlacesDatabaseRepository(db)
	_, err := databasePlacesRepository.Create(ctx, testPlace)
	if err != nil {
		t.Fatalf("Error creating place: %v", err)
	}
	return testPlace
}

func TestCreatingMeeting(t *testing.T) {
	//set up
	var ctx = context.Background()
	var testUser = setUpTestUser(t, ctx)
	var testPlace = setUpTestPlace(t, ctx)
	var databaseMeetingsRepository = repositoryMeetings.NewMeetingsDatabaseRepository(db)
	var currentMeeting = meeting.NewMeeting(testPlace.ID, testUser.ID, time.Now().Round(time.Microsecond).UTC(), time.Now().Round(time.Microsecond).UTC().Add(time.Hour*2), 3, meeting.Active)
	//main part
	_, errCreating := databaseMeetingsRepository.Create(ctx, currentMeeting)
	if errCreating != nil {
		t.Fatalf("Error in creating place: %v", errCreating)
	}
	meetingFromRepository, errFinding := databaseMeetingsRepository.FindByID(ctx, currentMeeting.ID)
	if errFinding != nil {
		t.Fatalf("Error in finding place: %v", errFinding)
	}

	//testing equality of found and created entities
	assert.True(t, reflect.DeepEqual(currentMeeting, meetingFromRepository))
	assert.Equal(t, currentMeeting, meetingFromRepository)
	errDeleting := databaseMeetingsRepository.Delete(ctx, currentMeeting)
	if errDeleting != nil {
		t.Fatalf("Error in deleting user: %v", errDeleting)
	}
	errDeleting = repositoryPlaces.NewPlacesDatabaseRepository(db).Delete(ctx, testPlace)
	if errDeleting != nil {
		t.Fatalf("Error in deleting place: %v", errDeleting)
	}
	errDeleting = repository.NewUsersDatabaseRepository(db).Delete(ctx, testUser)
	if errDeleting != nil {
		t.Fatalf("Error in deleting user %v", errDeleting)
	}
}

func TestFindingByCriteriaMeeting(t *testing.T) {
	//set up
	var ctx = context.Background()
	var testUser = setUpTestUser(t, ctx)
	var testPlace = setUpTestPlace(t, ctx)
	var databaseMeetingsRepository = repositoryMeetings.NewMeetingsDatabaseRepository(db)
	var firstMeeting = meeting.NewMeeting(testPlace.ID, testUser.ID, time.Now().Round(time.Microsecond).UTC(), time.Now().Round(time.Microsecond).UTC().Add(time.Hour*2), 3, meeting.Active)
	_, errCreating := databaseMeetingsRepository.Create(ctx, firstMeeting)
	if errCreating != nil {
		t.Fatalf("Error in creating place: %v", errCreating)
	}
	var anotherTestUser = setUpTestUser(t, ctx)
	var secondMeeting = meeting.NewMeeting(testPlace.ID, anotherTestUser.ID, time.Now().Round(time.Microsecond).UTC(), time.Now().Round(time.Microsecond).UTC().Add(time.Hour*2), 3, meeting.Active)
	_, errCreating = databaseMeetingsRepository.Create(ctx, secondMeeting)
	if errCreating != nil {
		t.Fatalf("Error in creating place: %v", errCreating)
	}
	var findingCriteria = query.FindCriteria{InitiatorID: uuid.NullUUID{UUID: testUser.ID, Valid: true}}
	//main part
	meetingsOfFirstTestUser, errFinding := databaseMeetingsRepository.FindByCriteria(ctx, findingCriteria)
	if errFinding != nil {
		t.Fatalf("Error in finding place: %v", errFinding)
	}

	//testing
	// first - check, if a placesWithFastFood contains right restaurants
	assert.True(t, slices.Contains(meetingsOfFirstTestUser, *firstMeeting))
	assert.False(t, slices.Contains(meetingsOfFirstTestUser, *secondMeeting))
	errDeleting := databaseMeetingsRepository.Delete(ctx, firstMeeting)
	errDeleting = databaseMeetingsRepository.Delete(ctx, secondMeeting)
	if errDeleting != nil {
		t.Fatalf("Error in deleting place: %v", errDeleting)
	}
	errDeleting = repository.NewUsersDatabaseRepository(db).Delete(ctx, testUser)
	if errDeleting != nil {
		t.Fatalf("Error in deleting user %v", errDeleting)
	}
	errDeleting = repository.NewUsersDatabaseRepository(db).Delete(ctx, anotherTestUser)
	if errDeleting != nil {
		t.Fatalf("Error in deleting user %v", errDeleting)
	}
}

func TestUpdatingMeeting(t *testing.T) {
	//set up
	var ctx = context.Background()
	var testUser = setUpTestUser(t, ctx)
	var testPlace = setUpTestPlace(t, ctx)
	var databaseMeetingsRepository = repositoryMeetings.NewMeetingsDatabaseRepository(db)
	var firstMeeting = meeting.NewMeeting(testPlace.ID, testUser.ID, time.Now().Round(time.Microsecond).UTC(), time.Now().Round(time.Microsecond).UTC().Add(time.Hour*2), 3, meeting.Active)
	_, errCreating := databaseMeetingsRepository.Create(ctx, firstMeeting)
	if errCreating != nil {
		t.Fatalf("Error in creating place: %v", errCreating)
	}
	//main part
	firstMeeting.State = meeting.Cancelled
	_, errUpdating := databaseMeetingsRepository.Update(ctx, firstMeeting)
	if errUpdating != nil {
		t.Fatalf("Error in updating place: %v", errUpdating)
	}
	meetingFromRepository, errFinding := databaseMeetingsRepository.FindByID(ctx, firstMeeting.ID)
	if errFinding != nil {
		t.Fatalf("Error in finding place: %v", errFinding)
	}

	//testing
	//checking the equality of found and updated entity
	assert.Equal(t, firstMeeting, meetingFromRepository)
	errDeleting := databaseMeetingsRepository.Delete(ctx, firstMeeting)
	if errDeleting != nil {
		t.Fatalf("Error in deleting place: %v", errDeleting)
	}
}

func TestDeletingMeeting(t *testing.T) {
	//set up
	var ctx = context.Background()
	var testUser = setUpTestUser(t, ctx)
	var testPlace = setUpTestPlace(t, ctx)
	var databaseMeetingsRepository = repositoryMeetings.NewMeetingsDatabaseRepository(db)
	var firstMeeting = meeting.NewMeeting(testPlace.ID, testUser.ID, time.Now().Round(time.Microsecond).UTC(), time.Now().Round(time.Microsecond).UTC().Add(time.Hour*2), 3, meeting.Active)
	_, errCreating := databaseMeetingsRepository.Create(ctx, firstMeeting)
	if errCreating != nil {
		t.Fatalf("Error in creating place: %v", errCreating)
	}
	//main part
	errDeleting := databaseMeetingsRepository.Delete(ctx, firstMeeting)
	if errDeleting != nil {
		t.Fatalf("Error in deleting place: %v", errDeleting)
	}
	_, errFinding := databaseMeetingsRepository.FindByID(ctx, firstMeeting.ID)
	//testing
	//checking the right error - no rows with such ID remained
	assert.Contains(t, errFinding.Error(), sql.ErrNoRows.Error())
	assert.Equal(t, "no such meeting: sql: no rows in result set", errFinding.Error())
	assert.True(t, errors.Is(errFinding, sql.ErrNoRows))
}
