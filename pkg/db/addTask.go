package db

import (
	_ "modernc.org/sqlite"
)

func AddTask(task Task) (int64, error) {
	var id int64
	res, err := DBOpen.Exec("INSERT INTO scheduler (date, title, comment, repeat) VALUES (?, ?, ?, ?)",
		task.Date,
		task.Title,
		task.Comment,
		task.Repeat,
	)
	if err != nil {
		return 0, err
	}

	id, err = res.LastInsertId()

	return id, nil

}
