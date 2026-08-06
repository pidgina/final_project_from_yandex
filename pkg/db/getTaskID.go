package db

import (
	"log"
	"strconv"
)

func GetTaskID(id string) (Task, error) {
	var res Task

	intStr, err := strconv.Atoi(id)
	if err != nil {
		return Task{}, err
	}

	err = DBOpen.QueryRow("SELECT * FROM scheduler WHERE id = ?", intStr).
		Scan(
			&res.ID,
			&res.Date,
			&res.Title,
			&res.Comment,
			&res.Repeat)

	if err != nil {
		log.Println(err)
		return Task{}, err
	}

	return res, nil

}
