package db

import "errors"

var NoneID = errors.New("Задача с указанным id не найдена")

func UpdateTask(task Task) error {
	if task.ID == "" {
		return NoneID
	}

	_, err := GetTaskID(task.ID)
	if err != nil {
		return NoneID
	}

	_, err = DBOpen.Exec(
		`UPDATE scheduler SET date = ?, title = ?, comment = ?, repeat = ? WHERE id = ?`,
		task.Date,
		task.Title,
		task.Comment,
		task.Repeat,
		task.ID,
	)

	return err
}

func UpdateTaskDate(id string, date string) error {
	_, err := DBOpen.Exec(
		"UPDATE scheduler SET date = ? WHERE id = ?",
		date,
		id,
	)

	return err
}
