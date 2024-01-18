package main

import (
	"Rvvijays/taskReminder/api"
	"Rvvijays/taskReminder/db"
	"log"
	"net/http"

	"github.com/go-chi/chi"
)

func init() {
	db.Init()
	go api.CheckDueTasks()
}

func main() {

	r := chi.NewRouter()

	// r.Use(middleware.Logger)
	// r.Use(middleware.Recoverer)

	r.Post("/tasks", api.CreateTask)
	r.Get("/tasks/{id}", api.GetTask)
	r.Put("/tasks/{id}", api.UpdateTask)
	r.Delete("/tasks/{id}", api.DeleteTask)

	log.Fatal(http.ListenAndServe(":8080", r))

}
