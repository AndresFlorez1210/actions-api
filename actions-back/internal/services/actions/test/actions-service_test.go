package test

import (
	entity "actions-back/internal/entity/actions"
	actionsService "actions-back/internal/services/actions"
	"actions-back/internal/utils"
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockActionsRepository struct {
	mock.Mock
}

func (m *MockActionsRepository) GetActions(ctx context.Context) ([]entity.Action, error) {
	args := m.Called(ctx)
	return args.Get(0).([]entity.Action), args.Error(1)
}

func TestGetActions(t *testing.T) {
	mockRepo := new(MockActionsRepository)
	service := actionsService.NewActionsService(mockRepo)
	ctx := context.Background()

	expectedActions := []entity.Action{
		{
			Ticker:     "AAPL",
			Company:    "Apple Inc",
			RatingFrom: "Buy",
			RatingTo:   "Buy",
			TargetFrom: "$150.00",
			TargetTo:   "$180.00",
		},
	}

	mockRepo.On("GetActions", ctx).Return(expectedActions, nil)

	result, err := service.GetActions(ctx)

	assert.NoError(t, err)
	assert.Equal(t, expectedActions, result)
	mockRepo.AssertExpectations(t)
}

func TestGetBestActionsForCategory(t *testing.T) {
	tests := []struct {
		name          string
		inputActions  []entity.Action
		expectedCount int
		shouldBeEmpty bool
	}{
		{
			name: "Should return best 3 buy rated actions with increasing target",
			inputActions: []entity.Action{
				{
					Ticker:     "AAPL",
					RatingFrom: "Buy",
					RatingTo:   "Buy",
					TargetFrom: "$150.00",
					TargetTo:   "$180.00",
				},
				{
					Ticker:     "GOOGL",
					RatingFrom: "Buy",
					RatingTo:   "Buy",
					TargetFrom: "$2500.00",
					TargetTo:   "$3000.00",
				},
				{
					Ticker:     "MSFT",
					RatingFrom: "Buy",
					RatingTo:   "Buy",
					TargetFrom: "$300.00",
					TargetTo:   "$350.00",
				},
				{
					Ticker:     "AMZN",
					RatingFrom: "Buy",
					RatingTo:   "Sell",
					TargetFrom: "$100.00",
					TargetTo:   "$90.00",
				},
			},
			expectedCount: 3,
			shouldBeEmpty: false,
		},
		{
			name: "Should return empty when no buy ratings",
			inputActions: []entity.Action{
				{
					Ticker:     "AAPL",
					RatingFrom: "Sell",
					RatingTo:   "Sell",
					TargetFrom: "$150.00",
					TargetTo:   "$180.00",
				},
			},
			expectedCount: 0,
			shouldBeEmpty: true,
		},
		{
			name: "Should filter out decreasing targets",
			inputActions: []entity.Action{
				{
					Ticker:     "AAPL",
					RatingFrom: "Buy",
					RatingTo:   "Buy",
					TargetFrom: "$180.00",
					TargetTo:   "$150.00",
				},
			},
			expectedCount: 0,
			shouldBeEmpty: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := actionsService.GetBestActionsForCategory(tt.inputActions)

			assert.NoError(t, err)
			if tt.shouldBeEmpty {
				assert.Empty(t, result)
			} else {
				assert.Len(t, result, tt.expectedCount)

				for _, action := range result {
					assert.Equal(t, actionsService.BUY_RATING, action.RatingFrom)
					assert.Equal(t, actionsService.BUY_RATING, action.RatingTo)

					targetFrom := utils.StringMoneyToFloat(action.TargetFrom)
					targetTo := utils.StringMoneyToFloat(action.TargetTo)
					assert.Greater(t, targetTo, targetFrom)
				}
			}
		})
	}
}

func TestGetBestActions(t *testing.T) {
	mockRepo := new(MockActionsRepository)
	service := actionsService.NewActionsService(mockRepo)
	ctx := context.Background()

	mockActions := []entity.Action{
		{
			Ticker:     "AAPL",
			RatingFrom: "Buy",
			RatingTo:   "Buy",
			TargetFrom: "$150.00",
			TargetTo:   "$180.00",
		},
		{
			Ticker:     "GOOGL",
			RatingFrom: "Buy",
			RatingTo:   "Buy",
			TargetFrom: "$2500.00",
			TargetTo:   "$3000.00",
		},
		{
			Ticker:     "MSFT",
			RatingFrom: "Sell",
			RatingTo:   "Sell",
			TargetFrom: "$300.00",
			TargetTo:   "$250.00",
		},
	}

	mockRepo.On("GetActions", ctx).Return(mockActions, nil)

	result, err := service.GetBestActions(ctx)

	assert.NoError(t, err)
	assert.NotEmpty(t, result)
	assert.LessOrEqual(t, len(result), 3)

	for _, action := range result {
		assert.Equal(t, actionsService.BUY_RATING, action.RatingFrom)
		assert.Equal(t, actionsService.BUY_RATING, action.RatingTo)

		targetFrom := utils.StringMoneyToFloat(action.TargetFrom)
		targetTo := utils.StringMoneyToFloat(action.TargetTo)
		assert.Greater(t, targetTo, targetFrom)
	}
}
