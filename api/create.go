package api

import (
	"Rvvijays/taskReminder/db"
	"Rvvijays/taskReminder/dep"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/rs/xid"
)

func CreateTask(w http.ResponseWriter, r *http.Request) {

	task := dep.Task{}

	err := json.NewDecoder(r.Body).Decode(&task)
	if err != nil {
		fmt.Println("err", err)
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	task.Time, err = dep.APIToIntTime(task.DueDateTime)
	if err != nil {
		fmt.Println("err", err)
		http.Error(w, "Invalid request payload", http.StatusBadRequest)

		return
	}

	task.ID = xid.New().String()

	err = db.Insert(&task)
	if err != nil {
		fmt.Println("err", err)
		http.Error(w, "Error creating task", http.StatusInternalServerError)

		return
	}

	fmt.Println("task created")

	json.NewEncoder(w).Encode(task)

}
