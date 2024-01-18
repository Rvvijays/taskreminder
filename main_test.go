package main

import (
	"Rvvijays/taskReminder/api"
	"Rvvijays/taskReminder/db"
	"Rvvijays/taskReminder/dep"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/go-chi/chi"
	"github.com/rs/xid"
	"github.com/stretchr/testify/assert"
)

func TestCreateTask(t *testing.T) {

	r := chi.NewRouter()

	// Register the createTask handler with the router
	r.Post("/tasks", api.CreateTask)

	// Define the request payload
	payload := `{"title": "Test Task", "description": "Test description", "priority": 1, "dueDateTime": "2024-01-20T15:04:05Z"}`

	// Create a new request with the payload
	req, err := http.NewRequest("POST", "/tasks", strings.NewReader(payload))
	if err != nil {
		t.Fatal(err)
	}

	// Set the Content-Type header
	req.Header.Set("Content-Type", "application/json")

	// Create a new recorder to capture the response
	rec := httptest.NewRecorder()

	// Serve the HTTP request to the recorder
	r.ServeHTTP(rec, req)

	// Check the HTTP status code
	assert.Equal(t, http.StatusOK, rec.Code)

}
func TestGetTask(t *testing.T) {
	// Initialize a new Chi router
	r := chi.NewRouter()

	// Register the getTask handler with the router
	r.Get("/tasks/{id}", api.GetTask)

	// Create a new request
	req, err := http.NewRequest("GET", "/tasks/1", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create a new recorder to capture the response
	rec := httptest.NewRecorder()

	// Serve the HTTP request to the recorder
	r.ServeHTTP(rec, req)

	// Check the HTTP status code
	assert.Equal(t, http.StatusOK, rec.Code)

	// Additional assertions based on your application logic
	// For example, you can assert the response body or headers
}

func TestUpdateTask(t *testing.T) {
	// Initialize a new Chi router
	r := chi.NewRouter()

	// Register the updateTask handler with the router
	r.Put("/tasks/{id}", api.UpdateTask)

	// Define the request payload for updating the task
	payload := `{"title": "Updated Task", "description": "Updated description", "priority": 2, "dueDateTime": "2024-01-21T15:04:05Z"}`

	// Create a new request with the payload
	req, err := http.NewRequest("PUT", "/tasks/1", strings.NewReader(payload))
	if err != nil {
		t.Fatal(err)
	}

	// Set the Content-Type header
	req.Header.Set("Content-Type", "application/json")

	// Create a new recorder to capture the response
	rec := httptest.NewRecorder()

	// Serve the HTTP request to the recorder
	r.ServeHTTP(rec, req)

	// Check the HTTP status code
	assert.Equal(t, http.StatusOK, rec.Code)

	// Additional assertions based on your application logic
	// For example, you can assert the updated data in the database
}

func TestDeleteTask(t *testing.T) {
	// Initialize a new Chi router
	r := chi.NewRouter()

	// Register the deleteTask handler with the router
	r.Delete("/tasks/{id}", api.DeleteTask)

	// Create a new request
	req, err := http.NewRequest("DELETE", "/tasks/1", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create a new recorder to capture the response
	rec := httptest.NewRecorder()

	// Serve the HTTP request to the recorder
	r.ServeHTTP(rec, req)

	// Check the HTTP status code
	assert.Equal(t, http.StatusOK, rec.Code)

	// Additional assertions based on your application logic
	// For example, you can assert that the task with ID 1 is deleted in the database
}

func TestReminderGoroutine(t *testing.T) {
	// Initialize a new mock database
	// mockDB, mock, err := sqlmock.New()
	// if err != nil {
	// 	t.Fatal(err)
	// }
	// defer mockDB.Close()

	// // Replace the global db variable with the mock database
	// db = sql.OpenDB(mockDB)

	// Prepare data for testing
	taskID := xid.New().String()
	taskTitle := "Test Task"
	taskDescription := "Test description"
	taskPriority := 1
	taskDueDateTime := time.Now().Add(1 * time.Minute).Format(time.RFC3339)
	taskTime, _ := dep.APIToIntTime(taskDueDateTime)

	// Insert a task into the database with a due time of 1 minute from now
	// mock.ExpectBegin()
	// mock.ExpectExec("INSERT INTO tasks").WillReturnResult(sqlmock.NewResult(1, 1))
	// mock.ExpectCommit()

	_, err := db.DB.Exec("INSERT INTO tasks1 (id, title, description, priority, time) VALUES ($1, $2, $3, $4, $5)",
		taskID, taskTitle, taskDescription, taskPriority, taskTime)
	assert.NoError(t, err)

	// Start the reminder Goroutine
	api.CheckDueTasks()

	// Wait for a short time to allow the Goroutine to execute
	time.Sleep(2 * time.Second)

	// Perform any additional assertions based on the expected behavior of the Goroutine
	// For example, you can check log output or other side effects of the Goroutine

	// Verify that the database was called with the correct query
	// if err := mock.ExpectationsWereMet(); err != nil {
	// 	t.Errorf("There were unfulfilled expectations: %s", err)
	// }
}
