package service

import (
	"time"

	"proj/pkg/db"
)

func PostDone(id string) error {
	task, err := db.GetTaskID(id)
	if err != nil {
		return err
	}

	if task.Repeat == "" {
		return db.DeleteTask(id)
	}

	nextDate, err := NextDate(time.Now(), task.Date, task.Repeat)
	if err != nil {
		return err
	}

	return db.UpdateTaskDate(task.ID, nextDate)
}
