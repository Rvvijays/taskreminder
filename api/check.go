package api

import (
	"Rvvijays/taskReminder/db"
	"Rvvijays/taskReminder/dep"
	"fmt"
	"log"
	"sort"
	"time"
)

func CheckDueTasks() {
	for {

		currentTime := time.Now()
		dueTime := currentTime.Add(5 * time.Minute)
		rows, err := db.DB.Query("SELECT id, title, description, priority, time FROM tasks1 WHERE time <= $1", dueTime.Unix())
		if err != nil {
			fmt.Println("Error querying tasks:", err)
			time.Sleep(1 * time.Minute)
			continue
		}

		var tasks []dep.Task

		for rows.Next() {
			var task dep.Task
			err := rows.Scan(&task.ID, &task.Title, &task.Description, &task.Priority, &task.Time)
			if err != nil {
				log.Println("Error scanning task:", err)
				continue
			}

			tasks = append(tasks, task)
		}

		rows.Close()

		sort.Slice(tasks, func(i, j int) bool {
			if tasks[i].Time != tasks[j].Time {
				return tasks[i].Time < tasks[j].Time
			}
			return tasks[i].Priority > tasks[j].Priority
		})

		for _, task := range tasks {
			if task.Time < currentTime.Unix() {
				message := fmt.Sprintf("%s was due at %s and is overdue!", task.Title, time.Unix(task.Time, 0).Format("3:04 pm on Monday"))
				fmt.Println(message)
			} else {
				message := fmt.Sprintf("%s is due at %s", task.Title, time.Unix(task.Time, 0).Format("3:04 pm on Monday"))
				fmt.Println(message)
			}
		}

		time.Sleep(1 * time.Minute)
	}
}
