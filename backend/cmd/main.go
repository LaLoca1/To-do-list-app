package main

import (
	"log"
	"net/http"

	"github.com/LaLoca1/to-do-list-app-backend/internal/api"
	"github.com/LaLoca1/to-do-list-app-backend/internal/db"
	"github.com/LaLoca1/to-do-list-app-backend/internal/services"
	"github.com/gorilla/mux"
)

func main() {
	db.InitDB() 

	taskService := services.NewTaskService() // Create TaskService 
	taskHandler := api.NewTaskHandler(taskService)

	r := mux.NewRouter()
    r.HandleFunc("/tasks", taskHandler.GetTasks).Methods("GET")
    r.HandleFunc("/tasks", taskHandler.CreateTask).Methods("POST")
    r.HandleFunc("/tasks/{id:[0-9]+}", taskHandler.UpdateTask).Methods("PUT")
    r.HandleFunc("/tasks/{id:[0-9]+}", taskHandler.DeleteTask).Methods("DELETE")

    log.Fatal(http.ListenAndServe(":8080", r)) // Start the HTTP server
}