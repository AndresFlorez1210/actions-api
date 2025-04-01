package test

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	controller "actions-back/internal/controller/actions"
	entity "actions-back/internal/entity/actions"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockActionsService struct {
	mock.Mock
}

func (m *MockActionsService) GetActions(ctx context.Context) ([]entity.Action, error) {
	args := m.Called(ctx)
	return args.Get(0).([]entity.Action), args.Error(1)
}

func (m *MockActionsService) GetBestActions(ctx context.Context) ([]entity.Action, error) {
	args := m.Called(ctx)
	return args.Get(0).([]entity.Action), args.Error(1)
}

func TestGetActions(t *testing.T) {
	mockService := new(MockActionsService)
	controller := controller.NewActionsController(mockService)

	testActions := []entity.Action{
		{
			Ticker:     "AAPL",
			Company:    "Apple Inc",
			RatingFrom: "Buy",
			RatingTo:   "Buy",
			TargetFrom: "$150.00",
			TargetTo:   "$180.00",
		},
	}

	mockService.On("GetActions", mock.Anything).Return(testActions, nil)

	req := httptest.NewRequest("GET", "/actions", nil)
	w := httptest.NewRecorder()

	controller.GetActions(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var response []entity.Action
	err := json.NewDecoder(w.Body).Decode(&response)
	assert.NoError(t, err)
	assert.Equal(t, testActions, response)

	mockService.AssertExpectations(t)
}

func TestGetBestActions(t *testing.T) {
	mockService := new(MockActionsService)
	controller := controller.NewActionsController(mockService)

	testActions := []entity.Action{
		{
			Ticker:     "AAPL",
			Company:    "Apple Inc",
			RatingFrom: "Buy",
			RatingTo:   "Buy",
			TargetFrom: "$150.00",
			TargetTo:   "$180.00",
		},
		{
			Ticker:     "GOOGL",
			Company:    "Google Inc",
			RatingFrom: "Buy",
			RatingTo:   "Buy",
			TargetFrom: "$2500.00",
			TargetTo:   "$3000.00",
		},
	}

	mockService.On("GetBestActions", mock.Anything).Return(testActions, nil)

	req := httptest.NewRequest("GET", "/best-actions", nil)
	w := httptest.NewRecorder()

	controller.GetBestActions(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var response []entity.Action
	err := json.NewDecoder(w.Body).Decode(&response)
	assert.NoError(t, err)
	assert.Equal(t, testActions, response)

	mockService.AssertExpectations(t)
}

func TestGetActionsError(t *testing.T) {
	mockService := new(MockActionsService)
	controller := controller.NewActionsController(mockService)

	mockService.On("GetActions", mock.Anything).Return([]entity.Action{}, assert.AnError)

	req := httptest.NewRequest("GET", "/actions", nil)
	w := httptest.NewRecorder()

	controller.GetActions(w, req)

	assert.Equal(t, http.StatusInternalServerError, w.Code)

	mockService.AssertExpectations(t)
}

func TestGetBestActionsError(t *testing.T) {
	mockService := new(MockActionsService)
	controller := controller.NewActionsController(mockService)

	mockService.On("GetBestActions", mock.Anything).Return([]entity.Action{}, assert.AnError)

	req := httptest.NewRequest("GET", "/best-actions", nil)
	w := httptest.NewRecorder()

	controller.GetBestActions(w, req)

	assert.Equal(t, http.StatusInternalServerError, w.Code)

	mockService.AssertExpectations(t)
}
