package api

import (
	"Rvvijays/taskReminder/db"
	"Rvvijays/taskReminder/dep"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
)

func UpdateTask(w http.ResponseWriter, r *http.Request) {
	taskId := chi.URLParam(r, "id")

	newTask := dep.Task{}

	err := json.NewDecoder(r.Body).Decode(&newTask)
	if err != nil {
		fmt.Println("Err", err)
		http.Error(w, "Invalid request payload", http.StatusBadRequest)

		return
	}
	newTask.Time, err = dep.APIToIntTime(newTask.DueDateTime)
	if err != nil {
		fmt.Println("err", err)
		http.Error(w, "Invalid request payload", http.StatusBadRequest)

		return
	}

	newTask.ID = taskId

	oldTask, err := db.Get(taskId)
	if err != nil {
		fmt.Println("err", err)
		http.Error(w, "Task not found", http.StatusNotFound)

		return
	}

	if oldTask.ID == newTask.ID {
		err = db.Update(&newTask)
		if err != nil {
			fmt.Println("Err", err)
			http.Error(w, "Error updating task", http.StatusInternalServerError)
			return
		}
	}

}
