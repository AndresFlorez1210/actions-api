package test

import (
	entity "actions-back/internal/entity/auth"
	authService "actions-back/internal/services/auth"
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockUserRepository struct {
	mock.Mock
}

func (m *MockUserRepository) Create(ctx context.Context, user *entity.User) error {
	args := m.Called(ctx, user)
	return args.Error(0)
}

func (m *MockUserRepository) GetByUsername(ctx context.Context, username string) (*entity.User, error) {
	args := m.Called(ctx, username)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*entity.User), args.Error(1)
}

func TestRegister(t *testing.T) {
	mockRepo := new(MockUserRepository)
	service := authService.NewAuthService(mockRepo)
	ctx := context.Background()

	user := &entity.User{
		Username: "testuser",
		Password: "password123",
	}

	mockRepo.On("Create", ctx, mock.Anything).Return(nil)

	token, err := service.Register(ctx, user)

	assert.NoError(t, err)
	assert.NotEmpty(t, token)
	mockRepo.AssertExpectations(t)
}

func TestLogin(t *testing.T) {
	mockRepo := new(MockUserRepository)
	service := authService.NewAuthService(mockRepo)
	ctx := context.Background()

	successUser := &entity.User{
		ID:       "123",
		Username: "testuser",
		Password: "password123",
	}
	successUser.HashPassword()

	wrongPasswordUser := &entity.User{
		ID:       "123",
		Username: "testuser",
		Password: "wrongpassword",
	}
	wrongPasswordUser.HashPassword()

	tests := []struct {
		name        string
		username    string
		password    string
		mockUser    *entity.User
		mockError   error
		expectError bool
		expectToken bool
	}{
		{
			name:        "Login successful",
			username:    "testuser",
			password:    "password123",
			mockUser:    successUser,
			expectError: false,
			expectToken: true,
		},
		{
			name:        "Invalid credentials",
			username:    "testuser",
			password:    "wrongpassword",
			mockUser:    wrongPasswordUser,
			expectError: true,
			expectToken: false,
		},
		{
			name:        "User not found",
			username:    "nonexistent",
			password:    "password123",
			mockUser:    nil,
			mockError:   assert.AnError,
			expectError: true,
			expectToken: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo.On("GetByUsername", ctx, tt.username).Return(tt.mockUser, tt.mockError)

			token, err := service.Login(ctx, tt.username, tt.password)

			if tt.expectError {
				assert.Error(t, err)
				assert.Empty(t, token)
			} else {
				assert.NoError(t, err)
				assert.NotEmpty(t, token)
			}
			mockRepo.AssertExpectations(t)
		})
	}
}
