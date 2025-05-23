// This is there the HTTP requests are handled

package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/LaLoca1/to-do-list-app-backend/internal/models"
	"github.com/gorilla/mux"
)

// TaskServiceInterface defines the methods required for the service
type TaskServiceInterface interface {
	GetTasks() ([]models.Task, error)
	CreateTask(task *models.Task) error
	UpdateTask(id int64, task *models.Task) error
	DeleteTask(id int64) error
}

type TaskHandler struct {
	service TaskServiceInterface
}

// MockTaskService is a mock for TaskService. However, Go doesn't allow you to directly assign a mock type to a variable of 
// another type unless the mock explicitly implements the same interface.

func NewTaskHandler(service TaskServiceInterface) *TaskHandler {
	return &TaskHandler{service: service}
}

// This handles the GET requests to retrieve all tasks
func (h *TaskHandler) GetTasks(w http.ResponseWriter, r *http.Request) {
	tasks, err := h.service.GetTasks()
	if err != nil {
		http.Error(w, "Error retrieving tasks", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tasks)
}

// CreateTask handles POST requests to create a new task
func (h *TaskHandler) CreateTask(w http.ResponseWriter, r *http.Request) {
	var newTask models.Task
	if err := json.NewDecoder(r.Body).Decode(&newTask); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	err := h.service.CreateTask(&newTask)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newTask)
}

// UpdateTask handles PUT requests to update a task
func (h *TaskHandler) UpdateTask(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		http.Error(w, "Invalid task ID", http.StatusBadRequest)
		return
	}

	var updatedTask models.Task
	if err := json.NewDecoder(r.Body).Decode(&updatedTask); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	err = h.service.UpdateTask(id, &updatedTask)
	if err != nil {
		http.Error(w, "Error updating task", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(updatedTask)
}

// DeleteTask handles DELETE requests to remove a task
func (h *TaskHandler) DeleteTask(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		http.Error(w, "Invalid task ID", http.StatusBadRequest)
		return
	}

	err = h.service.DeleteTask(id)
	if err != nil {
		http.Error(w, "Error deleting task", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
