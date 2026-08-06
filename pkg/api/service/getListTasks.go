package service

import (
	"database/sql"
	"time"

	"proj/pkg/db"

	_ "modernc.org/sqlite"
)

type TasksResp struct {
	Tasks []Task `json:"tasks"`
}

const (
	Limit = 50
)

func GetListTask(limit int, search string) (TasksResp, error) {

	var (
		rows *sql.Rows
		err  error
	)
	if search == "" {
		rows, err = db.DBOpen.Query(
			"SELECT * FROM scheduler ORDER BY date LIMIT ?",
			limit,
		)
		if err != nil {
			return TasksResp{}, err
		}
		defer rows.Close()

	} else {
		tims, err := time.Parse("02.01.2006", search)
		if err == nil {
			timesis := tims.Format(LayoutDate)

			rows, err = db.DBOpen.Query("SELECT * FROM scheduler WHERE date = ? ORDER BY date LIMIT ?",
				timesis, limit)
			if err != nil {
				return TasksResp{}, err
			}
			defer rows.Close()
		} else {
			searchFind := "%" + search + "%"
			rows, err = db.DBOpen.Query("SELECT * FROM scheduler WHERE title LIKE ? OR comment LIKE ? ORDER BY date LIMIT ?",
				searchFind, searchFind, limit)
			if err != nil {
				return TasksResp{}, err
			}
			defer rows.Close()
		}

	}

	var res Task
	var resTask TasksResp

	for rows.Next() {
		err := rows.Scan(&res.ID, &res.Date, &res.Title, &res.Comment, &res.Repeat)
		if err != nil {
			return TasksResp{}, err
		}
		resTask.Tasks = append(resTask.Tasks, res)
	}

	if err := rows.Err(); err != nil {
		return TasksResp{}, err
	}

	if resTask.Tasks == nil {
		resTask.Tasks = []Task{}
	}

	return resTask, nil

}
