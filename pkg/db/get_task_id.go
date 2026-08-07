package db

import (
	"database/sql"
	"errors"
	"log"
	"strconv"
)

func GetTaskID(id string) (Task, error) {
	var res Task

	intStr, err := strconv.Atoi(id)
	if err != nil {
		return Task{}, err
	}

	err = DBOpen.QueryRow("SELECT * FROM scheduler WHERE id = :id",
		sql.Named("id", intStr),
	).Scan(
		&res.ID,
		&res.Date,
		&res.Title,
		&res.Comment,
		&res.Repeat)

	if errors.Is(err, sql.ErrNoRows) {
		log.Println(err)
		return Task{}, err
	}

	if err != nil {
		log.Println(err)
		return Task{}, err
	}

	return res, nil

}
