package db

import (
	"database/sql"
	"time"

	_ "modernc.org/sqlite"
)

const (
	Limit = 50
)

func GetListTask(limit int, search string) ([]Task, error) {
	var (
		rows *sql.Rows
		err  error
	)

	if search == "" {
		rows, err = DBOpen.Query(
			"SELECT * FROM scheduler ORDER BY date LIMIT ?",
			limit,
		)
	} else {
		searchDate, parseErr := time.Parse("02.01.2006", search)

		if parseErr == nil {
			rows, err = DBOpen.Query(
				"SELECT * FROM scheduler WHERE date = ? ORDER BY date LIMIT ?",
				searchDate.Format("20060102"),
				limit,
			)
		} else {
			searchValue := "%" + search + "%"

			rows, err = DBOpen.Query(
				`SELECT * FROM scheduler WHERE title LIKE ? OR comment LIKE ? ORDER BY date LIMIT ?`,
				searchValue,
				searchValue,
				limit,
			)
		}
	}

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	tasks := make([]Task, 0)

	for rows.Next() {
		var task Task

		if err := rows.Scan(
			&task.ID,
			&task.Date,
			&task.Title,
			&task.Comment,
			&task.Repeat,
		); err != nil {
			return nil, err
		}

		tasks = append(tasks, task)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return tasks, nil
}
