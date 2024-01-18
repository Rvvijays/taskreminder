package api

import (
	"Rvvijays/taskReminder/db"
	"Rvvijays/taskReminder/dep"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
)

func GetTask(w http.ResponseWriter, r *http.Request) {

	taskId := chi.URLParam(r, "id")

	task, err := db.Get(taskId)
	if err != nil {
		fmt.Println("err", err)
		http.Error(w, "Task not found", http.StatusNotFound)

		return
	}

	task.DueDateTime = dep.IntTOAPITime(task.Time)

	json.NewEncoder(w).Encode(task)

}
