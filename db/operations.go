package db

import (
	"Rvvijays/taskReminder/dep"
	"database/sql"
	"fmt"
	"log"
)

func Get(taskId string) (*dep.Task, error) {

	task := dep.Task{}
	err := DB.QueryRow("SELECT id, title, description, priority, time FROM tasks1 WHERE id = $1", taskId).
		Scan(&task.ID, &task.Title, &task.Description, &task.Priority, &task.Time)

	if err != nil {
		fmt.Println("err", err)

		if err == sql.ErrNoRows {
			return nil, err
		}

		return nil, err
	}

	return &task, nil

}

func Insert(task *dep.Task) error {
	_, err := DB.Exec("INSERT INTO tasks1 (id, title, description, priority, time) VALUES ($1, $2, $3, $4,$5)",
		task.ID, task.Title, task.Description, task.Priority, task.Time)
	if err != nil {
		fmt.Println(err)
		return err
	}

	// lastInsertID, err := result.LastInsertId()
	// if err != nil {
	// 	fmt.Println(err)
	// 	return err
	// }
	// fmt.Println("data added.. ", lastInsertID)

	// task.ID = lastInsertID

	return nil
}

func Update(task *dep.Task) error {
	_, err := DB.Exec("UPDATE tasks1 SET title = $1, description = $2, priority = $3, time = $4 WHERE id = $5",
		task.Title, task.Description, task.Priority, task.Time, task.ID)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func Delete(taskId string) error {
	_, err := DB.Exec("DELETE FROM tasks1 WHERE id = $1", taskId)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}
