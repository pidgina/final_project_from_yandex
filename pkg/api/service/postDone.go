package service

import (
	"database/sql"
	"time"

	"proj/pkg/db"
)

func PostDone(id string) ([]byte, error) {
	var res string
	now := time.Now()

	taska, err := GetTaskID(id)
	if err != nil {
		return nil, err
	}

	if taska.Repeat == "" {
		jsByte, err := DeleteTask(id)
		if err != nil {
			return nil, err
		}
		return jsByte, nil
	} else {
		res, err = NextDate(now, taska.Date, taska.Repeat)
		if err != nil {
			return nil, err
		}

		_, err = db.DBOpen.Exec("UPDATE scheduler SET date = :date WHERE id = :id",
			sql.Named("date", res),
			sql.Named("id", taska.ID))

		if err != nil {
			return nil, err
		}

		return []byte("{}"), nil

	}

}
