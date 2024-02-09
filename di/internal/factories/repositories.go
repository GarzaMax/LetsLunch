package factories

import (
	gathering_place_repository "cmd/app/entities/gatheringPlace/repository"
	meeting_repository "cmd/app/entities/meeting/repository"
	user_repository "cmd/app/entities/user/repository"

	"cmd/di/internal/lookup"
	"context"
)

func CreateRepositoriesMeetingRepository(ctx context.Context, c lookup.Container) meeting_repository.MeetingsRepository {
	return meeting_repository.NewMeetingsDatabaseRepository(c.DB(ctx))
}

func CreateRepositoriesUserRepository(ctx context.Context, c lookup.Container) user_repository.UsersRepository {
	return user_repository.NewUsersDatabaseRepository(c.DB(ctx))
}

func CreateRepositoriesGatheringPlaceRepository(ctx context.Context, c lookup.Container) gathering_place_repository.PlacesRepository {
	return gathering_place_repository.NewPlacesDatabaseRepository(c.DB(ctx))
}
