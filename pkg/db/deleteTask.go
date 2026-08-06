package db

func DeleteTask(id string) ([]byte, error) {
	task, err := GetTaskID(id)
	if err != nil {
		return nil, NoneID
	}

	_, err = DBOpen.Exec(
		"DELETE FROM scheduler WHERE id = ?",
		task.ID,
	)
	if err != nil {
		return nil, err
	}

	return []byte("{}"), nil
}
