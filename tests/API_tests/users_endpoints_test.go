package API_tests

import (
	"cmd/app/entities/user"
	"cmd/app/entities/user/api"
	"cmd/app/entities/user/repository"
	"cmd/app/entities/user/usecases"
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/jwtauth/v5"
	"github.com/gofrs/uuid/v5"
	"github.com/gorilla/mux"
	"github.com/lestrrat-go/jwx/v2/jwt"
	_ "github.com/lib/pq"
	"github.com/ory/dockertest/v3"
	"github.com/pressly/goose/v3"
	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"reflect"
	"slices"
	"strings"
	"testing"
	"time"
)

var db *sql.DB

func AreUsersEqual(first *user.User, second *user.User) bool {
	if (first.ID != second.ID) || (first.CurrentMeetingId != second.CurrentMeetingId) ||
		(!time.Time.Equal(first.Birthday, second.Birthday)) || (first.PhoneNumber != second.PhoneNumber) ||
		(first.Gender != second.Gender) || (!slices.Equal(first.MeetingHistory, second.MeetingHistory)) ||
		(first.Rating != second.Rating) || (first.Username != second.Username) || (first.DisplayName != second.DisplayName) {
		return false
	}
	return true
}

func TestGetToken(t *testing.T) {
	var requestBody = `{"password" : "12345", "phoneNumber" : "+79528123333"}`
	date, _ := time.Parse(time.DateOnly, "2003-04-16")
	hashed, _ := bcrypt.GenerateFromPassword([]byte("12345"), 8)
	var testUser = user.NewUser("Steve", "Steve13", date, "+79528123333", user.Male, hashed)
	var usersRepository = repository.NewUsersDatabaseRepository(db)
	_, err := usersRepository.Create(context.Background(), testUser)
	if err != nil {
		t.Fatalf("error creating a user %v", err)
	}

	req := httptest.NewRequest(http.MethodGet, "/api/user/token", strings.NewReader(requestBody))
	w := httptest.NewRecorder()
	req = mux.SetURLVars(req, map[string]string{"phone_number": testUser.PhoneNumber, "password": "12345"})
	var auth = jwtauth.New("HS256", []byte("secret"), nil, jwt.WithAcceptableSkew(30*time.Second))
	var getTokenUsecase = usecases.NewGetTokenUseCase(usersRepository, auth)
	handler := api.GetTokenHandler{UseCase: getTokenUsecase}
	handler.ServeHTTP(w, req)
	res := w.Result()
	defer res.Body.Close()
	var actualResponse api.GetTokenResponse
	if err := json.Unmarshal([]byte(w.Body.String()), &actualResponse); err != nil {
		log.Fatalln(err)
	}
	status := w.Code
	assert.Equal(t, http.StatusOK, status)

	assert.NotEmpty(t, actualResponse)
}

func TestMain(m *testing.M) {
	pool, err := dockertest.NewPool("")
	if err != nil {
		log.Fatalf("Could not construct pool: %s", err)
	}

	//connecting to docker
	err = pool.Client.Ping()
	if err != nil {
		log.Fatalf("Could not connect to Docker: %s", err)
	}

	//creating test container
	resource, err := pool.Run("postgres", "latest", []string{
		"POSTGRES_PASSWORD=secret",
		"POSTGRES_USER=user",
		"POSTGRES_DB=testdb",
	})
	if err != nil {
		log.Fatalf("Could not start resource: %s", err)
	}
	var testPort = resource.GetPort("5432/tcp")

	//retry in case container is not ready to accept connections yet
	if err := pool.Retry(func() error {
		var err error
		db, err = sql.Open("postgres", fmt.Sprintf("host=localhost port=%s user=user password=secret dbname=testdb sslmode=disable", testPort))
		if err != nil {
			return err
		}
		return db.Ping()
	}); err != nil {
		log.Fatalf("Could not connect to database: %s", err)
	}
	if err := goose.SetDialect("postgres"); err != nil {
		panic(err)
	}
	if err := goose.Up(db, "../../migrations"); err != nil {
		panic(err)
	}

	code := m.Run()
	if err = pool.Purge(resource); err != nil {
		log.Fatalf("Could not purge resource: %s", err)
	}
	resource.Close()
	os.Exit(code)
}

func TestCreatingUsersWithSamePhoneNumber(t *testing.T) {
	var requestBody = `{"username": "katya76", "displayName": "katya", "birthday": "2006-12-13T08:08:08Z", "phoneNumber": "89816999983", "gender": 1}`
	req := httptest.NewRequest(http.MethodPost, "/api/user/create", strings.NewReader(requestBody))
	w := httptest.NewRecorder()
	var usersRepository = repository.NewUsersDatabaseRepository(db)
	var createUsecase = usecases.NewCreateUserUseCase(usersRepository)
	handler := api.CreateUserHandler{UseCase: createUsecase}
	handler.ServeHTTP(w, req)
	res := w.Result()
	defer res.Body.Close()
	var response api.JsonCreateUserResponse
	if err := json.Unmarshal([]byte(w.Body.String()), &response); err != nil {
		log.Fatalln(err)
	}
	status := w.Code
	assert.Equal(t, http.StatusOK, status)

	req = httptest.NewRequest(http.MethodPost, "/api/user/create", strings.NewReader(requestBody))
	w = httptest.NewRecorder()
	handler.ServeHTTP(w, req)
	res = w.Result()
	status = w.Code
	assert.Equal(t, http.StatusUnprocessableEntity, status)
	assert.True(t, !response.UserUUID.IsNil())
}

func TestCreatingUserShouldReturnStatus400(t *testing.T) {
	var requestBody = `{"username": "katya78", "birthday": "2006-12-13", "gender": 1}`
	req := httptest.NewRequest(http.MethodPost, "/api/user/create", strings.NewReader(requestBody))
	w := httptest.NewRecorder()
	var usersRepository = repository.NewUsersDatabaseRepository(db)
	var createUsecase = usecases.NewCreateUserUseCase(usersRepository)
	handler := api.CreateUserHandler{UseCase: createUsecase}
	handler.ServeHTTP(w, req)
	res := w.Result()
	defer res.Body.Close()
	status := w.Code
	assert.Equal(t, http.StatusBadRequest, status)
}

func TestFindUserByIdShouldReturnStatus200(t *testing.T) {
	date, _ := time.Parse(time.DateOnly, "2003-04-16")
	var testUser = user.NewUser("Steve", "Steve13", date, "+79528123333", user.Male, nil)
	var usersRepository = repository.NewUsersDatabaseRepository(db)
	_, err := usersRepository.Create(context.Background(), testUser)
	if err != nil {
		t.Fatalf("error creating a user %v", err)
	}
	var responseBody = api.JsonFindUserByIdResponse{
		ID:               testUser.ID,
		Username:         "Steve",
		DisplayName:      "Steve13",
		CurrentMeetingID: uuid.NullUUID{},
		MeetingHistory:   nil,
		Rating:           0,
		Birthday:         date,
		Gender:           user.Male,
	}
	req := httptest.NewRequest(http.MethodGet, "/{userID}", nil)
	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("userID", testUser.ID.String())
	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))
	w := httptest.NewRecorder()
	var findUsecase = usecases.NewFindUserByIdUseCase(usersRepository)
	handler := api.FindUserByIdHandler{UseCase: findUsecase}
	handler.ServeHTTP(w, req)
	res := w.Result()
	defer res.Body.Close()
	var response api.JsonFindUserByIdResponse
	if err := json.Unmarshal([]byte(w.Body.String()), &response); err != nil {
		log.Fatalln(err)
	}
	status := w.Code
	assert.Equal(t, http.StatusOK, status)
	assert.Equal(t, responseBody, response)
}

func TestFindUserByCriteriaShouldReturnStatus200(t *testing.T) {
	date, _ := time.Parse(time.DateOnly, "2003-04-16")
	var testUser = user.NewUser("Steve", "Steve13", date, "+79528123333", user.Male, nil)
	var usersRepository = repository.NewUsersDatabaseRepository(db)
	_, err := usersRepository.Create(context.Background(), testUser)
	if err != nil {
		t.Fatalf("error creating a user %v", err)
	}

	req := httptest.NewRequest(http.MethodGet, "/find?", nil)
	q := req.URL.Query()                        // Get a copy of the query values.
	q.Add("phone_number", testUser.PhoneNumber) // Add a new value to the set.
	req.URL.RawQuery = q.Encode()               // Encode and assign back to the original query.
	w := httptest.NewRecorder()
	req = mux.SetURLVars(req, map[string]string{"phone_number": testUser.PhoneNumber})
	var findUsecase = usecases.NewFindUsersByCriteriaUseCase(usersRepository)
	handler := api.FindUsersByCriteriaHandler{UseCase: findUsecase}
	handler.ServeHTTP(w, req)
	res := w.Result()
	defer res.Body.Close()
	var response []user.User
	if err := json.Unmarshal([]byte(w.Body.String()), &response); err != nil {
		log.Fatalln(err)
	}
	status := w.Code
	assert.Equal(t, http.StatusOK, status)

	var IDs = []uuid.UUID{}
	var userFromRepository = user.User{}
	for _, element := range response {
		IDs = append(IDs, element.ID)
		if element.ID == testUser.ID {
			userFromRepository = element
		}
	}
	assert.True(t, AreUsersEqual(testUser, &userFromRepository))
}

func TestUpdateUserShouldReturnStatus200(t *testing.T) {
	date, _ := time.Parse(time.DateOnly, "2003-04-16")
	var testUser = user.NewUser("Katya", "Katya13", date, "+79528123333", user.Female, nil)
	var usersRepository = repository.NewUsersDatabaseRepository(db)
	_, err := usersRepository.Create(context.Background(), testUser)
	if err != nil {
		t.Fatalf("")
	}
	testUser.Username = "Katya78"
	usersRepository.Update(context.Background(), testUser)
	var requestBody = `{"username": "katya78", "gender": 1}`
	req := httptest.NewRequest(http.MethodPut, "/api/user/update/{user_id}", strings.NewReader(requestBody))
	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("userID", testUser.ID.String())
	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))
	w := httptest.NewRecorder()
	var updateUsecase = usecases.NewUpdateUserUseCase(usersRepository)
	handler := api.UpdateUserHandler{UseCase: updateUsecase}
	handler.ServeHTTP(w, req)
	res := w.Result()
	defer res.Body.Close()
	var responseID api.JsonUpdateUserResponse
	if err := json.Unmarshal([]byte(w.Body.String()), &responseID); err != nil {
		log.Fatalln(err)
	}
	status := w.Code
	assert.Equal(t, http.StatusOK, status)

	req = httptest.NewRequest(http.MethodGet, "/{userID}", nil)
	rctx = chi.NewRouteContext()
	rctx.URLParams.Add("userID", testUser.ID.String())
	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))
	w = httptest.NewRecorder()
	var findUsecase = usecases.NewFindUserByIdUseCase(usersRepository)
	newHandler := api.FindUserByIdHandler{UseCase: findUsecase}
	newHandler.ServeHTTP(w, req)
	res = w.Result()
	defer res.Body.Close()
	var response api.JsonFindUserByIdResponse
	if err := json.Unmarshal([]byte(w.Body.String()), &response); err != nil {
		log.Fatalln(err)
	}
	var responseBody = api.JsonFindUserByIdResponse{}
	responseBody.Username = "katya78"
	/*responseBody.DisplayName = "Katya13"
	responseBody.CurrentMeetingID = uuid.NullUUID{}
	responseBody.Rating = 0
	responseBody.Birthday = date*/
	responseBody.Gender = user.Female
	responseBody.ID = testUser.ID
	assert.Equal(t, responseBody, response)
	// alright if it is true old values to be forgotten
}

func TestDeletingUser(t *testing.T) {
	date, _ := time.Parse(time.DateOnly, "2003-04-16")
	var testUser = user.NewUser("Katya", "Katya13", date, "+79528123333", user.Female, nil)
	var usersRepository = repository.NewUsersDatabaseRepository(db)
	req := httptest.NewRequest(http.MethodDelete, "/api/user/update/{user_id}", nil)
	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("userID", testUser.ID.String())
	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))
	w := httptest.NewRecorder()
	var deleteUsecase = usecases.NewDeleteUserUseCase(usersRepository)
	handler := api.DeleteUserHandler{UseCase: deleteUsecase}
	handler.ServeHTTP(w, req)
	res := w.Result()
	defer res.Body.Close()
	status := w.Code
	assert.Equal(t, http.StatusInternalServerError, status)

	_, err := usersRepository.Create(context.Background(), testUser)
	if err != nil {
		t.Fatalf("Error in creating user %v", err)
	}
	req = httptest.NewRequest(http.MethodDelete, "/api/user/update/{userID}", nil)
	rctx = chi.NewRouteContext()
	rctx.URLParams.Add("userID", testUser.ID.String())
	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))
	w = httptest.NewRecorder()
	handler.ServeHTTP(w, req)
	res = w.Result()
	defer res.Body.Close()
	status = w.Code
	var responseID api.JsonUpdateUserResponse
	json.Unmarshal([]byte(w.Body.String()), &responseID)
	assert.Equal(t, http.StatusOK, status)
	assert.True(t, reflect.DeepEqual(responseID, api.JsonUpdateUserResponse{}))
}

func TestUpdateUserShouldReturnStatus500(t *testing.T) {
	date, _ := time.Parse(time.DateOnly, "2003-04-16")
	var testUser = user.NewUser("Katya", "Katya78", date, "+79528123333", user.Female, nil)
	var usersRepository = repository.NewUsersDatabaseRepository(db)
	usersRepository.Update(context.Background(), testUser)
	var requestBody = `{"username": "katya78", "gender": 1}`
	req := httptest.NewRequest(http.MethodPut, "/api/user/update/{user_id}", strings.NewReader(requestBody))
	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("userID", testUser.ID.String())
	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))
	w := httptest.NewRecorder()
	var updateUsecase = usecases.NewUpdateUserUseCase(usersRepository)
	handler := api.UpdateUserHandler{UseCase: updateUsecase}
	handler.ServeHTTP(w, req)
	res := w.Result()
	defer res.Body.Close()
	status := w.Code
	assert.Equal(t, http.StatusInternalServerError, status)
}
