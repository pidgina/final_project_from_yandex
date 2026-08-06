package service

import (
	"time"

	"proj/pkg/db"
)

func PostDone(id string) ([]byte, error) {
	task, err := db.GetTaskID(id)
	if err != nil {
		return nil, err
	}

	if task.Repeat == "" {
		return db.DeleteTask(id)
	}

	nextDate, err := NextDate(
		time.Now(),
		task.Date,
		task.Repeat,
	)
	if err != nil {
		return nil, err
	}

	if err := db.UpdateTaskDate(task.ID, nextDate); err != nil {
		return nil, err
	}

	return []byte("{}"), nil
}
