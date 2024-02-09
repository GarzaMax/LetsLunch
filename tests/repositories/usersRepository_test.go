package repositories

import (
	"cmd/app/entities/user"
	"cmd/app/entities/user/query"
	"cmd/app/entities/user/repository"
	"context"
	"database/sql"
	"errors"
	"github.com/gofrs/uuid/v5"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/assert"
	"slices"
	"testing"
	"time"
)

func AreUsersEqual(first *user.User, second *user.User) bool {
	if (first.ID != second.ID) || (first.CurrentMeetingId != second.CurrentMeetingId) ||
		(!time.Time.Equal(first.Birthday, second.Birthday)) || (first.PhoneNumber != second.PhoneNumber) ||
		(first.Gender != second.Gender) || (!slices.Equal(first.MeetingHistory, second.MeetingHistory)) ||
		(first.Rating != second.Rating) || (first.Username != second.Username) || (first.DisplayName != second.DisplayName) {
		return false
	}
	return true
}

func TestCreatingUser(t *testing.T) {
	//set up
	var databaseUsersRepository = repository.NewUsersDatabaseRepository(db)
	date, _ := time.Parse(time.DateOnly, "2003-04-16")
	var currentUser = user.NewUser("Steve", "Steve13", date, "+79528123333", user.Male, []byte("1234567890"))
	currentUser.Rating = 5
	var ctx = context.Background()

	//main part
	_, errCreating := databaseUsersRepository.Create(ctx, currentUser)
	if errCreating != nil {
		t.Fatalf("Error in creating user: %v", errCreating)
	}
	userFromRepository, errFinding := databaseUsersRepository.FindUserByID(ctx, currentUser.ID)
	if errFinding != nil {
		t.Fatalf("Error in finding user: %v", errFinding)
	}

	//testing equality of found and created entities
	//assert.Equal(t, currentUser, userFromRepository)
	assert.True(t, AreUsersEqual(currentUser, userFromRepository))
	errDeleting := databaseUsersRepository.Delete(ctx, currentUser)
	if errDeleting != nil {
		t.Fatalf("Error in deleting user: %v", errDeleting)
	}
}

func TestFindingByCriteriaUser(t *testing.T) {
	//set up
	var databaseUsersRepository = repository.NewUsersDatabaseRepository(db)
	date, _ := time.Parse(time.DateOnly, "2003-04-16")
	var firstUser = user.NewUser("Steve", "Steve13", date, "+79528123333", user.Male, []byte("1234567890"))
	var ctx = context.Background()
	_, errCreating := databaseUsersRepository.Create(ctx, firstUser)
	if errCreating != nil {
		t.Fatalf("Error in creating user: %v", errCreating)
	}
	date, _ = time.Parse(time.DateOnly, "2003-04-16")
	var secondUser = user.NewUser("Masha", "Masha456767", date, "+79528123334", user.Female, []byte("1234567890"))
	_, errCreating = databaseUsersRepository.Create(ctx, secondUser)
	if errCreating != nil {
		t.Fatalf("Error in creating user: %v", errCreating)
	}
	date, _ = time.Parse(time.DateOnly, "2003-04-16")
	var thirdUser = user.NewUser("Steve", "Steve13", date, "+79528123332", user.Male, []byte("1234567890"))
	_, errCreating = databaseUsersRepository.Create(ctx, thirdUser)
	if errCreating != nil {
		t.Fatalf("Error in creating user: %v", errCreating)
	}
	var findingCriteria = query.FindCriteria{Username: sql.NullString{String: "Steve", Valid: true}}

	//main part
	usersWithSameUsername, errFinding := databaseUsersRepository.FindUsersByCriteria(ctx, findingCriteria)
	if errFinding != nil {
		t.Fatalf("Error in finding user: %v", errFinding)
	}

	//testing
	var userFromRepostory = user.User{}
	var IDs = []uuid.UUID{}
	for _, element := range usersWithSameUsername {
		IDs = append(IDs, element.ID)
		if element.ID == firstUser.ID {
			userFromRepostory = element
		}
	}
	// first - check, if were found right users with Username "Steve"
	assert.True(t, slices.Contains(IDs, firstUser.ID))
	assert.True(t, slices.Contains(IDs, thirdUser.ID))
	assert.False(t, slices.Contains(IDs, secondUser.ID))
	//second - check equality of returned entity
	//assert.Equal(t, *firstUser, userFromRepostory)
	assert.True(t, AreUsersEqual(firstUser, &userFromRepostory))
	errDeleting := databaseUsersRepository.Delete(ctx, firstUser)
	errDeleting = databaseUsersRepository.Delete(ctx, secondUser)
	errDeleting = databaseUsersRepository.Delete(ctx, thirdUser)
	if errDeleting != nil {
		t.Fatalf("Error in deleting user: %v", errDeleting)
	}
}

func TestUpdatingUser(t *testing.T) {
	//set up
	var databaseUsersRepository = repository.NewUsersDatabaseRepository(db)
	date, _ := time.Parse(time.DateOnly, "2003-04-16")
	var currentUser = user.NewUser("Katya", "Katya14", date, "+79528123330", user.Female, []byte("1234567890"))
	var ctx = context.Background()
	_, errCreating := databaseUsersRepository.Create(ctx, currentUser)
	if errCreating != nil {
		t.Fatalf("Error in creating user: %v", errCreating)
	}
	//main part
	currentUser.DisplayName = "Katya15"
	_, errUpdating := databaseUsersRepository.Update(ctx, currentUser)
	if errUpdating != nil {
		t.Fatalf("Error in updating user: %v", errUpdating)
	}
	userFromRepository, errFinding := databaseUsersRepository.FindUserByID(ctx, currentUser.ID)
	if errFinding != nil {
		t.Fatalf("Error in finding user: %v", errFinding)
	}

	//testing
	//checking the equality of found and updated entity
	//assert.Equal(t, currentUser, userFromRepository)
	assert.True(t, AreUsersEqual(currentUser, userFromRepository))
	errDeleting := databaseUsersRepository.Delete(ctx, currentUser)
	if errDeleting != nil {
		t.Fatalf("Error in deleting user: %v", errDeleting)
	}
}

func TestDeletingUser(t *testing.T) {
	//set up
	var databaseUsersRepository = repository.NewUsersDatabaseRepository(db)
	date, _ := time.Parse(time.DateOnly, "2003-04-16")
	var currentUser = user.NewUser("@testUser", "Steve13", date, "+79528123333", user.Male, []byte("1234567890"))
	var ctx = context.Background()
	_, errCreating := databaseUsersRepository.Create(ctx, currentUser)
	if errCreating != nil {
		t.Fatalf("Error in creating user: %v", errCreating)
	}
	//main part
	errDeleting := databaseUsersRepository.Delete(ctx, currentUser)
	if errDeleting != nil {
		t.Fatalf("Error in deleting user: %v", errDeleting)
	}
	_, errFinding := databaseUsersRepository.FindUserByID(ctx, currentUser.ID)
	//testing
	//checking the right error - no rows with such ID remained
	assert.Contains(t, errFinding.Error(), sql.ErrNoRows.Error())
	assert.Equal(t, "no such user: sql: no rows in result set", errFinding.Error())
	assert.True(t, errors.Is(errFinding, sql.ErrNoRows))
}
