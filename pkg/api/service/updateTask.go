package service

import (
	"proj/pkg/db"
)

func UpdateTask(task Task) ([]byte, error) {
	if task.ID == "" {
		return nil, NoneID
	}

	_, err := GetTaskID(task.ID)
	if err != nil {
		return nil, NoneID
	}

	_, err = db.DBOpen.Exec(`UPDATE scheduler SET date = ?, title = ?, comment = ?, repeat = ? WHERE id = ?`,
		task.Date,
		task.Title,
		task.Comment,
		task.Repeat,
		task.ID,
	)
	if err != nil {
		return nil, err
	}

	return []byte("{}"), nil
}
