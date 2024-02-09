package factories

import (
	gathering_place_usecase "cmd/app/entities/gatheringPlace/usecases"
	meeting_usecase "cmd/app/entities/meeting/usecases"
	user_usecase "cmd/app/entities/user/usecases"
	"cmd/di/internal/lookup"
	"context"
)

func CreateUseCasesFindMeeting(ctx context.Context, c lookup.Container) *meeting_usecase.FindMeetingByIdUseCase {
	return meeting_usecase.NewFindMeetingByIdUseCase(
		c.Repositories().MeetingRepository(ctx),
	)
}

func CreateUseCasesFindUser(ctx context.Context, c lookup.Container) *user_usecase.FindUserByIdUseCase {
	return user_usecase.NewFindUserByIdUseCase(
		c.Repositories().UserRepository(ctx),
	)
}

func CreateUseCasesFindGatheringPlace(ctx context.Context, c lookup.Container) *gathering_place_usecase.FindGatheringPlaceByIdUseCase {
	return gathering_place_usecase.NewFindGatheringPlaceByIdUseCase(
		c.Repositories().GatheringPlaceRepository(ctx),
	)
}

func CreateUseCasesCreateUser(ctx context.Context, c lookup.Container) *user_usecase.CreateUserUseCase {
	return user_usecase.NewCreateUserUseCase(
		c.Repositories().UserRepository(ctx),
	)
}

func CreateUseCasesUpdateUser(ctx context.Context, c lookup.Container) *user_usecase.UpdateUserUseCase {
	return user_usecase.NewUpdateUserUseCase(
		c.Repositories().UserRepository(ctx),
	)
}

func CreateUseCasesDeleteUser(ctx context.Context, c lookup.Container) *user_usecase.DeleteUserUseCase {
	return user_usecase.NewDeleteUserUseCase(
		c.Repositories().UserRepository(ctx),
	)
}

func CreateUseCasesCreateMeeting(ctx context.Context, c lookup.Container) *meeting_usecase.CreateMeetingUseCase {
	return meeting_usecase.NewCreateMeetingUseCase(
		c.Repositories().MeetingRepository(ctx),
	)
}

func CreateUseCasesUpdateMeeting(ctx context.Context, c lookup.Container) *meeting_usecase.UpdateMeetingUseCase {
	return meeting_usecase.NewUpdateMeetingUseCase(
		c.Repositories().MeetingRepository(ctx),
	)
}

func CreateUseCasesDeleteMeeting(ctx context.Context, c lookup.Container) *meeting_usecase.DeleteMeetingUseCase {
	return meeting_usecase.NewDeleteMeetingUseCase(
		c.Repositories().MeetingRepository(ctx),
	)
}

func CreateUseCasesCreateGatheringPlace(ctx context.Context, c lookup.Container) *gathering_place_usecase.CreateGatheringPlaceUseCase {
	return gathering_place_usecase.NewCreateGatheringPlaceUseCase(
		c.Repositories().GatheringPlaceRepository(ctx),
	)
}

func CreateUseCasesUpdateGatheringPlace(ctx context.Context, c lookup.Container) *gathering_place_usecase.UpdateGatheringPlaceUseCase {
	return gathering_place_usecase.NewUpdateGatheringPlaceUseCase(
		c.Repositories().GatheringPlaceRepository(ctx),
	)
}

func CreateUseCasesDeleteGatheringPlace(ctx context.Context, c lookup.Container) *gathering_place_usecase.DeleteGatheringPlaceUseCase {
	return gathering_place_usecase.NewDeleteGatheringPlaceUsecase(
		c.Repositories().GatheringPlaceRepository(ctx),
	)
}

func CreateUseCasesGetToken(ctx context.Context, c lookup.Container) *user_usecase.GetTokenUseCase {
	return user_usecase.NewGetTokenUseCase(
		c.Repositories().UserRepository(ctx),
		c.TokenAuth(ctx),
	)
}

func CreateUseCasesFindUsers(ctx context.Context, c lookup.Container) *user_usecase.FindUsersByCriteriaUseCase {
	return user_usecase.NewFindUsersByCriteriaUseCase(
		c.Repositories().UserRepository(ctx),
	)
}

func CreateUseCasesFindGatheringPlaces(ctx context.Context, c lookup.Container) *gathering_place_usecase.FindGatheringPlacesByCriteriaUseCase {
	return gathering_place_usecase.NewFindGatheringPlacesByCriteriaUseCase(
		c.Repositories().GatheringPlaceRepository(ctx),
	)
}

func CreateUseCasesFindMeetings(ctx context.Context, c lookup.Container) *meeting_usecase.FindMeetingsByCriteriaUseCase {
	return meeting_usecase.NewFindMeetingsByCriteriaUseCase(
		c.Repositories().MeetingRepository(ctx),
	)
}
