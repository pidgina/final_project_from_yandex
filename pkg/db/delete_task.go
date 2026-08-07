package db

import "database/sql"

func DeleteTask(id string) error {
	task, err := GetTaskID(id)
	if err != nil {
		return NoneID
	}

	_, err = DBOpen.Exec(
		"DELETE FROM scheduler WHERE id = :id", sql.Named("id", task.ID),
	)

	return err
}
