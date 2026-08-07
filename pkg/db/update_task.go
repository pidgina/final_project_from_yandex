package db

import (
	"database/sql"
	"errors"
)

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
		`UPDATE scheduler SET date = :date, title = :title, comment = :comment, repeat = :repeat WHERE id = :id`,
		sql.Named("date", task.Date),
		sql.Named("title", task.Title),
		sql.Named("comment", task.Comment),
		sql.Named("repeat", task.Repeat),
		sql.Named("id", task.ID),
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
