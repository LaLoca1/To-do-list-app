package api

import (
	"encoding/json"
	// "os"
	// "io" 
	"net/http"
	"net/http/httptest"
	// "strings"
	"testing"

	"github.com/LaLoca1/to-do-list-app-backend/internal/models"
	// "github.com/LaLoca1/to-do-list-app-backend/internal/services"
	// "github.com/gorilla/mux"
	"github.com/stretchr/testify/mock"
)

// MockTaskService is a mock of the TaskService interface
type MockTaskService struct {
	mock.Mock
}

func (m *MockTaskService) GetTasks() ([]models.Task, error) {
	args := m.Called() 
	return args.Get(0).([]models.Task), args.Error(1) 
}

func (m *MockTaskService) CreateTask(task *models.Task) error {
	args := m.Called(task) 
	return args.Error(0) 
}

func (m *MockTaskService) UpdateTask(id int64, task *models.Task) error {
	args := m.Called(id, task) 
	return args.Error(0) 
}

func (m *MockTaskService) DeleteTask(id int64) error {
	args := m.Called(id) 
	return args.Error(0) 
}

// Test for GetTasks 
func TestTaskHandler_GetTasks(t *testing.T) {
	// Mock the TaskService
	mockService := new(MockTaskService)
	mockTasks := []models.Task{
		{ID: 1, Title: "Test Task 1", Description: "Task Description", Completed: false},
		{ID: 2, Title: "Test Task 2", Description: "Task Description", Completed: true},
	}
	mockService.On("GetTasks").Return(mockTasks, nil)

	// Initialize the handler with the mocked service
	handler := NewTaskHandler(mockService)

	// Create a request and recorder
	req, err := http.NewRequest("GET", "/tasks", nil)
	if err != nil {
		t.Fatalf("could not create request: %v", err)
	}

	rr := httptest.NewRecorder()
	handler.GetTasks(rr, req)

	// Check the status code
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("expected status %v, got %v", http.StatusOK, status)
	}

	// Check the response body
	var tasks []models.Task
	err = json.NewDecoder(rr.Body).Decode(&tasks)
	if err != nil {
		t.Fatalf("could not decode response body: %v", err)
	}
	if len(tasks) != len(mockTasks) {
		t.Fatalf("expected %d tasks, got %d", len(mockTasks), len(tasks))
	}
	for i, task := range tasks {
		if task != mockTasks[i] {
			t.Errorf("expected task %v, got %v", mockTasks[i], task)
		}
	}

	// Verify the mock method was called
	mockService.AssertExpectations(t)
}
