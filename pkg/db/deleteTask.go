package db

func DeleteTask(id string) error {
	task, err := GetTaskID(id)
	if err != nil {
		return NoneID
	}

	_, err = DBOpen.Exec(
		"DELETE FROM scheduler WHERE id = ?", task.ID,
	)

	return err
}
