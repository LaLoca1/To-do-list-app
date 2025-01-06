// This is the entry point of the application

package main

import (
	"log"
	"net/http"

	"github.com/LaLoca1/to-do-list-app-backend/internal/api"
	"github.com/LaLoca1/to-do-list-app-backend/internal/db"
	"github.com/LaLoca1/to-do-list-app-backend/internal/services"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {
	db.InitDB()

	taskService := services.NewTaskService()       // Create TaskService
	taskHandler := api.NewTaskHandler(taskService) // Create the TaskHandler

	r := mux.NewRouter()

	r.HandleFunc("/tasks", taskHandler.GetTasks).Methods("GET")
	r.HandleFunc("/tasks", taskHandler.CreateTask).Methods("POST")
	r.HandleFunc("/tasks/{id:[0-9]+}", taskHandler.UpdateTask).Methods("PUT")
	r.HandleFunc("/tasks/{id:[0-9]+}", taskHandler.DeleteTask).Methods("DELETE")

	// Enable CORS
	corsHandler := handlers.CORS(
		handlers.AllowedOrigins([]string{"http://localhost:3000"}),        // Allow only frontend's origin
		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE"}), // Allow these HTTP methods
		handlers.AllowedHeaders([]string{"Content-Type"}),                 // Allow Content-Type header
	)(r)

	// Start the server
	log.Println("Server started on http://localhost:8080")
	if err := http.ListenAndServe(":8080", corsHandler); err != nil {
		log.Fatal("Error starting server:", err)
	}
}
