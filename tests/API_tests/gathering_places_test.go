package API_tests

import (
	"cmd/app/entities/gatheringPlace"
	"cmd/app/entities/gatheringPlace/api"
	"cmd/app/entities/gatheringPlace/repository"
	"cmd/app/entities/gatheringPlace/usecases"
	"cmd/app/models"
	"context"
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/assert"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestFindByCriteriaPlace(t *testing.T) {
	var testPlace = gatheringPlace.NewGatheringPlace(models.Address{
		Country:        "Russia",
		City:           "Moscow",
		StreetName:     "Unknown",
		HouseNumber:    "",
		BuildingNumber: 44,
	}, 1000, gatheringPlace.Russian, 5, "+79816999986", "", "Foodtime", "")
	var secondTestPlace = gatheringPlace.NewGatheringPlace(models.Address{
		Country:        "Russia",
		City:           "Moscow",
		StreetName:     "Unknown",
		HouseNumber:    "",
		BuildingNumber: 44,
	}, 1000, gatheringPlace.Russian, 4, "+79816999985", "", "FoodLunch", "")
	var placesRepository = repository.NewPlacesDatabaseRepository(db)
	_, err := placesRepository.Create(context.Background(), testPlace)
	if err != nil {
		t.Fatalf("Error in creating place %v", err)
	}
	_, err = placesRepository.Create(context.Background(), secondTestPlace)
	if err != nil {
		t.Fatalf("Error in creating place %v", err)
	}

	expectedResponse := []gatheringPlace.GatheringPlace{*testPlace}
	var actualResponse []gatheringPlace.GatheringPlace

	req := httptest.NewRequest(http.MethodGet, "/api/gatheringPlace/find?", nil)
	q := req.URL.Query()
	q.Add("rating", "5")
	req.URL.RawQuery = q.Encode()
	w := httptest.NewRecorder()
	var findByCriteriaUsecase = usecases.NewFindGatheringPlacesByCriteriaUseCase(placesRepository)
	handler := api.FindGatheringPlacesByCriteriaHandler{UseCase: findByCriteriaUsecase}

	handler.ServeHTTP(w, req)
	res := w.Result()
	defer res.Body.Close()
	status := w.Code
	assert.Equal(t, http.StatusOK, status)
	if err := json.Unmarshal([]byte(w.Body.String()), &actualResponse); err != nil {
		log.Fatalln(err)
	}
	assert.ElementsMatch(t, expectedResponse, actualResponse)
}

func TestCreatingPlace(t *testing.T) {
	var requestBody = `{
    "address": {
      "country": "Russia",
      "city": "Piter",
      "streetName ": "kronverksky",
      "buildingNumber": 5,
      "houseNumber": "49"
    },
    "averagePrice": 1,
    "cuisineType": 0,
    "rating": 1,
    "phoneNumber": "88005553535" 
     }`

	var placesRepository = repository.NewPlacesDatabaseRepository(db)
	var createUsecase = usecases.NewCreateGatheringPlaceUseCase(placesRepository)
	handler := api.CreateGatheringPlaceHandler{UseCase: createUsecase}
	req := httptest.NewRequest(http.MethodPost, "/api/gatheringPlace/create", strings.NewReader(requestBody))
	w := httptest.NewRecorder()
	handler.ServeHTTP(w, req)
	res := w.Result()
	defer res.Body.Close()
	var actualResponse api.JsonCreateGatheringPlaceResponse
	if err := json.Unmarshal([]byte(w.Body.String()), &actualResponse); err != nil {
		log.Fatalln(err)
	}
	status := w.Code
	assert.Equal(t, http.StatusOK, status)
	assert.True(t, !actualResponse.GatheringPlaceID.IsNil())

	req = httptest.NewRequest(http.MethodPost, "/api/gatheringPlace/create", strings.NewReader(requestBody))
	w = httptest.NewRecorder()
	handler.ServeHTTP(w, req)
	res = w.Result()
	status = w.Code
	assert.Equal(t, http.StatusOK, status)
}

func TestFindByIDPlace(t *testing.T) {
	var testPlace = gatheringPlace.NewGatheringPlace(models.Address{
		Country:        "Russia",
		City:           "Moscow",
		StreetName:     "Unknown",
		HouseNumber:    "",
		BuildingNumber: 44,
	}, 1000, gatheringPlace.FastFood, 5, "+79816999986", "", "Foodtime", "")
	var placesRepository = repository.NewPlacesDatabaseRepository(db)
	_, err := placesRepository.Create(context.Background(), testPlace)
	if err != nil {
		return
	}

	var expectedResponse = api.JsonFindGatheringPlaceByIdResponse{
		ID:          testPlace.ID,
		Address:     testPlace.Address,
		AvgPrice:    testPlace.AveragePrice,
		CusineType:  testPlace.CuisineType,
		Rating:      testPlace.Rating,
		PhoneNumber: testPlace.PhoneNumber}
	var actualResponse api.JsonFindGatheringPlaceByIdResponse

	req := httptest.NewRequest(http.MethodGet, "/api/gatheringPlace/find_by_id/{gatheringPlaceID}", nil)
	w := httptest.NewRecorder()
	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("placeID", testPlace.ID.String())
	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))

	var findByIDUsecase = usecases.NewFindGatheringPlaceByIdUseCase(placesRepository)
	handler := api.FindGatheringPlaceByIdHandler{UseCase: findByIDUsecase}

	handler.ServeHTTP(w, req)
	res := w.Result()
	defer res.Body.Close()
	status := w.Code
	assert.Equal(t, http.StatusOK, status)
	if err := json.Unmarshal([]byte(w.Body.String()), &actualResponse); err != nil {
		log.Fatalln(err)
	}
	assert.Equal(t, expectedResponse, actualResponse)
}

func TestDeletePlace(t *testing.T) {
	var testPlace = gatheringPlace.NewGatheringPlace(models.Address{
		Country:        "Russia",
		City:           "Moscow",
		StreetName:     "Unknown",
		HouseNumber:    "",
		BuildingNumber: 44,
	}, 1000, gatheringPlace.FastFood, 5, "+79816999986", "", "Foodtime", "")
	var placesRepository = repository.NewPlacesDatabaseRepository(db)

	req := httptest.NewRequest(http.MethodDelete, "/api/gatheringPlace/update/{gatheringPlaceID}", nil)
	w := httptest.NewRecorder()
	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("placeID", testPlace.ID.String())
	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))

	var deletePlaceUsecase = usecases.NewDeleteGatheringPlaceUsecase(placesRepository)
	handler := api.DeleteGatheringPlaceHandler{UseCase: deletePlaceUsecase}

	handler.ServeHTTP(w, req)
	res := w.Result()
	defer res.Body.Close()
	status := w.Code
	assert.Equal(t, http.StatusInternalServerError, status)

	_, err := placesRepository.Create(context.Background(), testPlace)
	if err != nil {
		return
	}
	req = httptest.NewRequest(http.MethodDelete, "/api/gatheringPlace/update/{gatheringPlaceID}", nil)
	w = httptest.NewRecorder()
	rctx = chi.NewRouteContext()
	rctx.URLParams.Add("placeID", testPlace.ID.String())
	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))

	handler.ServeHTTP(w, req)
	res = w.Result()
	defer res.Body.Close()
	status = w.Code
	assert.Equal(t, http.StatusOK, status)
}
