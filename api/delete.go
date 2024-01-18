package api

import (
	"Rvvijays/taskReminder/db"
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
)

func DeleteTask(w http.ResponseWriter, r *http.Request) {
	taskId := chi.URLParam(r, "id")

	task, err := db.Get(taskId)
	if err != nil {
		fmt.Println("err", err)
		http.Error(w, "Task not found", http.StatusNotFound)

		return
	}

	if task.ID == taskId {
		err = db.Delete(taskId)
		if err != nil {
			fmt.Println("Err", err)
			http.Error(w, "Error deleting task", http.StatusInternalServerError)

			return
		}
	}

}
