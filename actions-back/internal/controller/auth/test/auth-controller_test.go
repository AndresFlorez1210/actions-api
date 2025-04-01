package test

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	controller "actions-back/internal/controller/auth"
	entity "actions-back/internal/entity/auth"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockAuthService struct {
	mock.Mock
}

func (m *MockAuthService) Register(ctx context.Context, user *entity.User) (string, error) {
	args := m.Called(ctx, user)
	return args.String(0), args.Error(1)
}

func (m *MockAuthService) Login(ctx context.Context, username string, password string) (string, error) {
	args := m.Called(ctx, username, password)
	return args.String(0), args.Error(1)
}

func TestRegister(t *testing.T) {
	mockService := new(MockAuthService)
	controller := controller.NewAuthController(mockService)

	user := &entity.User{
		Username: "testuser",
		Password: "password123",
	}

	body, _ := json.Marshal(user)
	req := httptest.NewRequest("POST", "/register", bytes.NewBuffer(body))
	w := httptest.NewRecorder()

	mockService.On("Register", mock.Anything, mock.MatchedBy(func(u *entity.User) bool {
		return u.Username == user.Username && u.Password == user.Password
	})).Return("test-token", nil)

	controller.Register(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var response map[string]string
	err := json.NewDecoder(w.Body).Decode(&response)
	assert.NoError(t, err)
	assert.Equal(t, "test-token", response["token"])

	mockService.AssertExpectations(t)
}

func TestLogin(t *testing.T) {
	mockService := new(MockAuthService)
	controller := controller.NewAuthController(mockService)

	credentials := struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}{
		Username: "testuser",
		Password: "password123",
	}

	body, _ := json.Marshal(credentials)
	req := httptest.NewRequest("POST", "/login", bytes.NewBuffer(body))
	w := httptest.NewRecorder()

	mockService.On("Login", mock.Anything, credentials.Username, credentials.Password).Return("test-token", nil)

	controller.Login(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var response map[string]string
	err := json.NewDecoder(w.Body).Decode(&response)
	assert.NoError(t, err)
	assert.Equal(t, "test-token", response["token"])

	mockService.AssertExpectations(t)
}

func TestRegisterError(t *testing.T) {
	mockService := new(MockAuthService)
	controller := controller.NewAuthController(mockService)

	body := []byte(`{"invalid": "json"`)
	req := httptest.NewRequest("POST", "/register", bytes.NewBuffer(body))
	w := httptest.NewRecorder()

	controller.Register(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestLoginError(t *testing.T) {
	mockService := new(MockAuthService)
	controller := controller.NewAuthController(mockService)

	body := []byte(`{"invalid": "json"`)
	req := httptest.NewRequest("POST", "/login", bytes.NewBuffer(body))
	w := httptest.NewRecorder()

	controller.Login(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestRegisterServiceError(t *testing.T) {
	mockService := new(MockAuthService)
	controller := controller.NewAuthController(mockService)

	user := &entity.User{
		Username: "testuser",
		Password: "password123",
	}

	body, _ := json.Marshal(user)
	req := httptest.NewRequest("POST", "/register", bytes.NewBuffer(body))
	w := httptest.NewRecorder()

	mockService.On("Register", mock.Anything, mock.MatchedBy(func(u *entity.User) bool {
		return u.Username == user.Username && u.Password == user.Password
	})).Return("", assert.AnError)

	controller.Register(w, req)

	assert.Equal(t, http.StatusInternalServerError, w.Code)

	mockService.AssertExpectations(t)
}

func TestLoginServiceError(t *testing.T) {
	mockService := new(MockAuthService)
	controller := controller.NewAuthController(mockService)

	credentials := struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}{
		Username: "testuser",
		Password: "password123",
	}

	body, _ := json.Marshal(credentials)
	req := httptest.NewRequest("POST", "/login", bytes.NewBuffer(body))
	w := httptest.NewRecorder()

	mockService.On("Login", mock.Anything, credentials.Username, credentials.Password).Return("", assert.AnError)

	controller.Login(w, req)

	assert.Equal(t, http.StatusInternalServerError, w.Code)

	mockService.AssertExpectations(t)
}
