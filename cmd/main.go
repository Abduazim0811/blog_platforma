package main

import (
	"blog/internal/https/api"
	"blog/internal/storage"
	"log"
)

func main() {
	db, err := storage.OpenSql()
	if err != nil {
		log.Fatal(err)
	}
	api.Run(db)
}
