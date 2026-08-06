package db

import (
	"errors"
)

var NoneID = errors.New("Задача с указанным id не найдена")

func UpdateTask(task Task) ([]byte, error) {
	if task.ID == "" {
		return nil, NoneID
	}

	_, err := GetTaskID(task.ID)
	if err != nil {
		return nil, NoneID
	}

	_, err = DBOpen.Exec(`UPDATE scheduler SET date = ?, title = ?, comment = ?, repeat = ? WHERE id = ?`,
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
func UpdateTaskDate(id string, date string) error {
	_, err := DBOpen.Exec(
		"UPDATE scheduler SET date = ? WHERE id = ?",
		date,
		id,
	)

	return err
}
