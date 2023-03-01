package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func main() {
	// Connect to database
	db, err := gorm.Open("sqlite3", "tasks.db")
	if err != nil {
		log.Fatal("Error connecting to database")
	}
	defer db.Close()

	// Set up router
	r := mux.NewRouter()

	// Define endpoints
	r.HandleFunc("/tasks", createTask).Methods("POST")
	r.HandleFunc("/tasks/{id}", updateTask).Methods("PUT")
	r.HandleFunc("/tasks/{id}", deleteTask).Methods("DELETE")

	// Start server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}
	log.Fatal(http.ListenAndServe(":8080", r))
}

// Create a new task
func createTask(w http.ResponseWriter, r *http.Request) {
	// Parse JSON request body
	var task Task
	err := json.NewDecoder(r.Body).Decode(&task)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	// Save task to database
	db, err := gorm.Open("sqlite3", "tasks.db")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()

	db.Create(&task)

	// Return success response
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(task)
}

// Update an existing task
func updateTask(w http.ResponseWriter, r *http.Request) {
	// Parse task ID from URL parameters
	vars := mux.Vars(r)
	id := vars["id"]
	// Parse JSON request body
	var taskUpdates Task
	err := json.NewDecoder(r.Body).Decode(&taskUpdates)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Update task in database
	db, err := gorm.Open("sqlite3", "tasks.db")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()

	var task Task
	db.Where("id = ?", id).First(&task)
	task.Title = taskUpdates.Title
	task.Description = taskUpdates.Description
	task.Status = taskUpdates.Status
	db.Save(&task)

	// Return success response
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(task)
}

// Delete an existing task
func deleteTask(w http.ResponseWriter, r *http.Request) {
	// Parse task ID from URL parameters
	vars := mux.Vars(r)
	id := vars["id"]
	// Delete task from database
	db, err := gorm.Open("sqlite3", "tasks.db")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()

	var task Task
	db.Where("id = ?", id).First(&task)
	db.Delete(&task)

	// Return success response
	w.WriteHeader(http.StatusOK)
}

// Define task struct
type Task struct {
	gorm.Model
	Title       string `json:"title"`
	Description string `json:"description"`
	Status      string `json:"status"`
}
