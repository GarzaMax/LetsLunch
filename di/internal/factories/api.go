package factories

import (
	"cmd/app/auth"
	gathering_place_api "cmd/app/entities/gatheringPlace/api"
	meeting_api "cmd/app/entities/meeting/api"
	user_api "cmd/app/entities/user/api"
	"cmd/di/internal/lookup"
	"context"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/go-chi/jwtauth/v5"
	"github.com/lestrrat-go/jwx/v2/jwt"
	"time"
)

func CreateRouter(ctx context.Context, c lookup.Container) *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(cors.Handler(cors.Options{
		// AllowedOrigins:   []string{"https://foo.com"}, // Use this to allow specific origin hosts
		AllowedOrigins: []string{"https://*", "http://*"},
		// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))

	r.Post("/api/user/token", c.API().GetTokenHandler(ctx).ServeHTTP)
	r.Post("/api/user/create", c.API().CreateUserHandler(ctx).ServeHTTP)

	r.Route("/api", func(r chi.Router) {
		r.Use(c.AuthConfig(ctx).AuthMiddleware)

		r.Route("/meeting", func(r chi.Router) {
			r.Get("/find", c.API().FindMeetingsHandler(ctx).ServeHTTP)

			r.Route("/find_by_id", func(r chi.Router) {
				r.Get("/{meetingID}", c.API().FindMeetingHandler(ctx).ServeHTTP)
			})

			r.Post("/create", c.API().CreateMeetingHandler(ctx).ServeHTTP)

			r.Route("/update", func(r chi.Router) {
				r.Put("/{meetingID}", c.API().UpdateMeetingHandler(ctx).ServeHTTP)
				r.Delete("/{meetingID}", c.API().DeleteMeetingHandler(ctx).ServeHTTP)
			})
		})

		r.Route("/user", func(r chi.Router) {
			r.Get("/find", c.API().FindUsersHandler(ctx).ServeHTTP)

			r.Route("/find_by_id", func(r chi.Router) {
				r.Get("/{userID}", c.API().FindUserHandler(ctx).ServeHTTP)
			})

			r.Post("/create", c.API().CreateUserHandler(ctx).ServeHTTP)

			r.Route("/update", func(r chi.Router) {
				r.Put("/{userID}", c.API().UpdateUserHandler(ctx).ServeHTTP)
				r.Delete("/{userID}", c.API().DeleteUserHandler(ctx).ServeHTTP)
			})
		})

		r.Route("/gatheringPlace", func(r chi.Router) {
			r.Get("/find", c.API().FindGatheringPlacesHandler(ctx).ServeHTTP)

			r.Route("/find_by_id", func(r chi.Router) {
				r.Get("/{placeID}", c.API().FindGatheringPlaceHandler(ctx).ServeHTTP)
			})

			r.Post("/create", c.API().CreateGatheringPlaceHandler(ctx).ServeHTTP)

			r.Route("/update", func(r chi.Router) {
				r.Put("/{placeID}", c.API().UpdateGatheringPlaceHandler(ctx).ServeHTTP)
				r.Delete("/{placeID}", c.API().DeleteGatheringPlaceHandler(ctx).ServeHTTP)
			})
		})
	})

	return r
}

func CreateAuthConfig(ctx context.Context, c lookup.Container) *auth.AuthConfig {
	return auth.NewAuthConfig(
		c.Repositories().UserRepository(ctx),
		c.TokenAuth(ctx),
	)
}

func CreateTokenAuth(ctx context.Context, c lookup.Container) *jwtauth.JWTAuth {
	return jwtauth.New("HS256", []byte(c.Config(ctx).Secret), nil, jwt.WithAcceptableSkew(30*time.Second))
}

func CreateAPIFindMeetingHandler(ctx context.Context, c lookup.Container) *meeting_api.FindMeetingByIdHandler {
	return meeting_api.NewFindMeetingByIdHandler(
		c.UseCases().FindMeeting(ctx),
	)
}

func CreateAPIFindUserHandler(ctx context.Context, c lookup.Container) *user_api.FindUserByIdHandler {
	return user_api.NewFindUserByIdHandler(
		c.UseCases().FindUser(ctx),
	)
}

func CreateAPIFindGatheringPlaceHandler(ctx context.Context, c lookup.Container) *gathering_place_api.FindGatheringPlaceByIdHandler {
	return gathering_place_api.NewFindGatheringCaseByIdHandler(
		c.UseCases().FindGatheringPlace(ctx),
	)
}

func CreateAPICreateUserHandler(ctx context.Context, c lookup.Container) *user_api.CreateUserHandler {
	return user_api.NewCreateUserHandler(
		c.UseCases().CreateUser(ctx),
	)
}

func CreateAPIUpdateUserHandler(ctx context.Context, c lookup.Container) *user_api.UpdateUserHandler {
	return user_api.NewUpdateUserHandler(
		c.UseCases().UpdateUser(ctx),
	)
}

func CreateAPIDeleteUserHandler(ctx context.Context, c lookup.Container) *user_api.DeleteUserHandler {
	return user_api.NewDeleteUserHandler(
		c.UseCases().DeleteUser(ctx),
	)
}

func CreateAPICreateMeetingHandler(ctx context.Context, c lookup.Container) *meeting_api.CreateMeetingHandler {
	return meeting_api.NewCreateMeetingHandler(
		c.UseCases().CreateMeeting(ctx),
	)
}

func CreateAPIUpdateMeetingHandler(ctx context.Context, c lookup.Container) *meeting_api.UpdateMeetingHandler {
	return meeting_api.NewUpdateMeetingHandler(
		c.UseCases().UpdateMeeting(ctx),
	)
}

func CreateAPIDeleteMeetingHandler(ctx context.Context, c lookup.Container) *meeting_api.DeleteMeetingHandler {
	return meeting_api.NewDeleteMeetingHandler(
		c.UseCases().DeleteMeeting(ctx),
	)
}

func CreateAPICreateGatheringPlaceHandler(ctx context.Context, c lookup.Container) *gathering_place_api.CreateGatheringPlaceHandler {
	return gathering_place_api.NewCreateGatheringPlaceHandler(
		c.UseCases().CreateGatheringPlace(ctx),
	)
}

func CreateAPIUpdateGatheringPlaceHandler(ctx context.Context, c lookup.Container) *gathering_place_api.UpdateGatheringPlaceHandler {
	return gathering_place_api.NewUpdateGatheringPlaceHandler(
		c.UseCases().UpdateGatheringPlace(ctx),
	)
}

func CreateAPIDeleteGatheringPlaceHandler(ctx context.Context, c lookup.Container) *gathering_place_api.DeleteGatheringPlaceHandler {
	return gathering_place_api.NewDeleteGatheringPlaceHandler(
		c.UseCases().DeleteGatheringPlace(ctx),
	)
}

func CreateAPIGetTokenHandler(ctx context.Context, c lookup.Container) *user_api.GetTokenHandler {
	return user_api.NewGetTokenHandler(
		c.UseCases().GetToken(ctx),
	)
}

func CreateAPIFindUsersHandler(ctx context.Context, c lookup.Container) *user_api.FindUsersByCriteriaHandler {
	return user_api.NewFindUsersByCriteriaHandler(
		c.UseCases().FindUsers(ctx),
	)
}

func CreateAPIFindGatheringPlacesHandler(ctx context.Context, c lookup.Container) *gathering_place_api.FindGatheringPlacesByCriteriaHandler {
	return gathering_place_api.NewFindGatheringPlacesByCriteriaHandler(
		c.UseCases().FindGatheringPlaces(ctx),
	)
}

func CreateAPIFindMeetingsHandler(ctx context.Context, c lookup.Container) *meeting_api.FindMeetingsByCriteriaHandler {
	return meeting_api.NewFindMeetingsByCriteriaHandler(
		c.UseCases().FindMeetings(ctx),
	)
}
