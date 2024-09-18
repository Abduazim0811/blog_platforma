package storage

import (
	"database/sql"
	"log"
	_ "github.com/lib/pq"
)

func OpenSql() (*sql.DB, error) {
	db, err := sql.Open("postgres", "postgres://postgres:Abdu0811@localhost:5432/n9?sslmode=disable")
	if err != nil {
		log.Println("failed to open sql:", err)
		return nil, err
	}

	if err = db.Ping(); err != nil {
		log.Println("unable to connect database:", err)
		return nil, err
	}
	return db, err
}
