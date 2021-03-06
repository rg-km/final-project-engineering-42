package main

import (
	"backend/api"
	"backend/repository"
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "db/store-app.db")
	if err != nil {
		panic(err)
	}

	userRepo := repository.NewUserRepository(db)
	bookRepo := repository.NewBookRepository(db)
	cartRepo := repository.NewCartRepository(db)
	mainAPI := api.NewAPI(*userRepo, *bookRepo, *cartRepo)
	mainAPI.Start()
}
