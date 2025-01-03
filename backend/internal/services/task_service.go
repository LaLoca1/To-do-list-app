package services

import (
	"database/sql"
	"errors"
	"log"

	"github.com/LaLoca1/to-do-list-app-backend/internal/models"
	"github.com/LaLoca1/to-do-list-app-backend/internal/db"
)

type TaskService struct {
	db *sql.DB 
}

// This creates a new TaskService
func NewTaskService() *TaskService {
	return &TaskService{
		db: db.GetDB(), 
	}
}

// Returns all tasks from the database
func (s *TaskService) GetTasks() ([]models.Task, error) {
	rows, err := s.db.Query("Select id, title, description, completed FROM tasks") 
	if err != nil {
		log.Println("Error fetching tasks:", err) 
		return nil, err
	}
	defer rows.Close() 

	var tasks []models.Task
	for rows.Next() {
		var task models.Task 
		err := rows.Scan(&task.ID, &task.Title, &task.Description, &task.Completed) 
		if err != nil {
			log.Println("Error scanning task:", err) 
			return nil, err
		}
		tasks = append(tasks, task) 
	}
	return tasks, nil 
}

// Adds a new task to the database 
func (s *TaskService) CreateTask(task *models.Task) error {
	if task.Title == "" {
		return errors.New("task title is required") 
	}

	_, err := s.db.Exec("INSERT INTO tasks (title, description, completed) VALUES (?, ?, ?)", task.Title, task.Description, task.Completed)
	if err != nil {
		log.Println("Error inserting task:", err) 
		return err 
	}
	return nil 
}

func (s *TaskService) UpdateTask(id int64, task *models.Task) error {
	if task.Title == "" {
		return errors.New("task title is required") 
	}

	_, err := s.db.Exec("UPDATE tasks SET title = ?, description = ?, completed = ? WHERE id = ?", task.Title, task.Description, task.Completed, id) 
	if err != nil {
		log.Println("Error updating task:", err) 
		return err 
	}
	return nil 
}

func (s *TaskService) DeleteTask(id int64) error {
	_, err := s.db.Exec("DELETE FROM tasks WHERE id = ?", id) 
	if err != nil {
		log.Println("Error deleting task:", err) 
		return err 
	}
	return nil 
}


