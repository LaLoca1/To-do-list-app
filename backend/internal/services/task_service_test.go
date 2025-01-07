package services

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/LaLoca1/to-do-list-app-backend/internal/models"
)

func TestTaskService_GetTasks(t *testing.T) {
	// Mock database
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("failed to create sqlmock: %v", err)
	}
	defer db.Close()

	service := &TaskService{db: db}

	// Mock database rows
	rows := sqlmock.NewRows([]string{"id", "title", "description", "completed"}).
		AddRow(1, "Task 1", "Description 1", false).
		AddRow(2, "Task 2", "Description 2", true)

	mock.ExpectQuery(`(?i)^SELECT id, title, description, completed FROM tasks$`).
		WillReturnRows(rows)

	// Execute the service method
	tasks, err := service.GetTasks()
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	// Assert results
	if len(tasks) != 2 {
		t.Fatalf("expected 2 tasks, got %d", len(tasks))
	}
	if tasks[0].Title != "Task 1" || tasks[1].Completed != true {
		t.Errorf("unexpected task data")
	}

	// Ensure all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("unmet expectations: %v", err)
	}
}

func TestTaskService_CreateTask(t *testing.T) {
	// Mock database
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("failed to create sqlmock: %v", err)
	}
	defer db.Close()

	service := &TaskService{db: db}

	// Mock task
	task := &models.Task{
		Title:       "New Task",
		Description: "Description",
		Completed:   false,
	}

	// Expect insert query with full query match
	mock.ExpectExec(`(?i)^INSERT INTO tasks \(title, description, completed\) VALUES \(\?, \?, \?\)$`).
		WithArgs(task.Title, task.Description, task.Completed).
		WillReturnResult(sqlmock.NewResult(1, 1)) // What the mock database should return when SQL Query is executed
		// first arg, the last inserted ID, represents id of newly inserted row
		// second arg, number of rows affected by query. so would be 1 for insert operation 

	// Execute the service method
	err = service.CreateTask(task)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	// Ensure all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("unmet expectations: %v", err)
	}
}

func TestTaskService_UpdateTask(t *testing.T) {
	// Mock database
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("failed to create sqlmock: %v", err)
	}
	defer db.Close()

	service := &TaskService{db: db}

	// Mock task
	task := &models.Task{
		Title:       "Updated Task",
		Description: "Updated Description",
		Completed:   true,
	}

	// Expect update query with escaped `?`
	mock.ExpectExec(`(?i)^UPDATE tasks SET title = \?, description = \?, completed = \? WHERE id = \?$`).
		WithArgs(task.Title, task.Description, task.Completed, int64(1)).
		WillReturnResult(sqlmock.NewResult(1, 1))

	// Execute the service method
	err = service.UpdateTask(1, task)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	// Ensure all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("unmet expectations: %v", err)
	}
}

func TestTaskService_DeleteTask(t *testing.T) {
	// Mock database
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("failed to create sqlmock: %v", err)
	}
	defer db.Close()

	service := &TaskService{db: db}

	// Execute delete query 
	mock.ExpectExec(`(?i)^DELETE FROM tasks WHERE id = \?$`).
		WithArgs(int64(1)).
		WillReturnResult(sqlmock.NewResult(1, 1))

	// Execute the service method
	err = service.DeleteTask(1)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	// Ensure all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("unmet expectations: %v", err)
	}
}
