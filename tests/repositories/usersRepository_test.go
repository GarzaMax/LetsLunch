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
)

func TestCreatingUser(t *testing.T) {
	//set up
	var databaseUsersRepository = repository.NewUsersDatabaseRepository(db)
	var currentUser = user.NewUser()
	currentUser.Username = "Steve"
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
	assert.Equal(t, currentUser, userFromRepository)
	errDeleting := databaseUsersRepository.Delete(ctx, currentUser)
	if errDeleting != nil {
		t.Fatalf("Error in deleting user: %v", errDeleting)
	}
}

func TestFindingByCriteriaUser(t *testing.T) {
	//set up
	var databaseUsersRepository = repository.NewUsersDatabaseRepository(db)
	var firstUser = user.NewUser()
	firstUser.Username = "Steve"
	var ctx = context.Background()
	_, errCreating := databaseUsersRepository.Create(ctx, firstUser)
	if errCreating != nil {
		t.Fatalf("Error in creating user: %v", errCreating)
	}
	var secondUser = user.NewUser()
	secondUser.Username = "Masha"
	_, errCreating = databaseUsersRepository.Create(ctx, secondUser)
	if errCreating != nil {
		t.Fatalf("Error in creating user: %v", errCreating)
	}
	var thirdUser = user.NewUser()
	thirdUser.Username = "Steve"
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
	assert.Equal(t, *firstUser, userFromRepostory)

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
	var currentUser = user.NewUser()
	currentUser.DisplayName = "Katya14"
	var ctx = context.Background()
	_, errCreating := databaseUsersRepository.Create(ctx, currentUser)
	if errCreating != nil {
		t.Fatalf("Error in creating user: %v", errCreating)
	}
	//main part
	currentUser.DisplayName = "Katya15"
	currentUser.Username = "Katya"
	currentUser.Age = 18
	currentUser.Gender = user.Female
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
	assert.Equal(t, currentUser, userFromRepository)
	errDeleting := databaseUsersRepository.Delete(ctx, currentUser)
	if errDeleting != nil {
		t.Fatalf("Error in deleting user: %v", errDeleting)
	}
}

func TestDeletingGatheringUser(t *testing.T) {
	//set up
	var databaseUsersRepository = repository.NewUsersDatabaseRepository(db)
	var currentUser = user.NewUser()
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
