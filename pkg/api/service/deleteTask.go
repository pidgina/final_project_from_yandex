package service

import (
	"errors"

	"proj/pkg/db"
)

var NoneID = errors.New("Задача с указанным id не найдена")

func DeleteTask(id string) ([]byte, error) {

	taska, err := GetTaskID(id)
	if err != nil {
		return nil, NoneID
	}

	_, err = db.DBOpen.Exec("DELETE FROM scheduler WHERE id = ?", taska.ID)
	if err != nil {
		return nil, err
	}

	return []byte("{}"), nil

}
