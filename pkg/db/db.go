package db

import (
	"database/sql"
	"log"

	_ "modernc.org/sqlite"
)

const Schema string = `CREATE TABLE IF NOT EXISTS scheduler (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    date CHAR(8) NOT NULL DEFAULT "",
    title VARCHAR(20) NOT NULL DEFAULT "",
    comment TEXT NOT NULL DEFAULT "",
    repeat VARCHAR(128) NOT NULL DEFAULT ""
);

CREATE INDEX IF NOT EXISTS idx_date ON scheduler (date);`

var DBOpen *sql.DB

func Init(dbFile string) error {
	var err error

	DBOpen, err = sql.Open("sqlite", dbFile)
	if err != nil {
		log.Println(err)
		return err
	}

	_, err = DBOpen.Exec(Schema)
	if err != nil {
		log.Println(err)
		DBOpen.Close()
		return err
	}

	return nil
}
