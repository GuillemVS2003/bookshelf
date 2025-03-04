package main

import (
	"database/sql"
	"fmt"
	"sync"

	_ "github.com/lib/pq"
)

const (
	host = "localhost"
	port = 5432
	user = "postgres"
	password = "123zxc"
	database = "postgres"
)

func connect() *sql.DB {

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
				host, port, user, password, database)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	return db
}

var lock = &sync.Mutex{}
var db *sql.DB

func getDb() *sql.DB {

	if db == nil {
		lock.Lock()
		defer lock.Unlock()

		if db == nil {
			db = connect()
		}
	}

	return db
}

func findBooks() ([]Book, error) {

	rows, err := getDb().Query("select * from book")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var books []Book

	for rows.Next() {
		var book Book
		err := rows.Scan(&book.Id, &book.Name, &book.Author,
				 &book.Published, &book.Pages, &book.Lang, &book.Publisher)
		if err != nil {
			return books, err
		}
		books = append(books, book)
	}

	err = rows.Err()
	if err != nil {
		return books, err
	}

	return books, nil
}
