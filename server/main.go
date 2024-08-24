package main

import (
	"log"
	"net/http"
	"os"

	"github.com/go-playground/validator/v10"
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

var (
	validate           *validator.Validate
	watchListIdCounter = 68
	watchLists         = make(map[string]*WatchList)
	schema             = `
	CREATE TABLE WatchList (
		id varchar(255),
		name varchar(255)
	);`
)

type WatchList struct {
	ID   string  `json:"id" db:"id"`
	Name *string `json:"name" db:"name"`
}

func main() {
	validate = validator.New(validator.WithRequiredStructEnabled())

	dbFile := "./testDB.db"
	if _, err := os.Stat(dbFile); err == nil {
		if err := os.Remove(dbFile); err != nil {
			log.Fatalf("Failed to remove existing database file: %v", err)
		}
		log.Println("Existing database file removed")
	}

	db, err := sqlx.Connect("sqlite3", "./testDB.db")
	if err != nil {
		log.Fatalln(err)
	}

	db.MustExec(schema)

	// var testQuery string
	testQuery := "INSERT INTO WatchList (id, name) VALUES (:id, :name)"

	testWatchList := WatchList{
		ID:   "1",
		Name: func() *string { name := "testList"; return &name }(),
	}

	tx := db.MustBegin()
	_, err = tx.NamedExec(testQuery, testWatchList)
	if err != nil {
		log.Fatal(err)
	}

	err = tx.Commit()
	if err != nil {
		log.Fatal(err)
	}

	savedWatchLists := []WatchList{}
	err = db.Select(&savedWatchLists, "SELECT * FROM WatchList")
	if err != nil {
		log.Fatalln(err)
	}

	// db, err := sql.Open("sqlite3", dbFile)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// defer db.Close()
	wl := WatchList{}
	structToString(wl)

	http.HandleFunc("GET /watchlist", getAllWatchLists)
	http.HandleFunc("GET /watchlist/{id}", getWatchList)
	http.HandleFunc("POST /watchlist", createWatchList)
	http.HandleFunc("PUT /watchlist/{id}", setWatchList)
	http.HandleFunc("PATCH /watchlist/{id}", patchWatchList)
	http.HandleFunc("DELETE /watchlist/{id}", deleteWatchList)

	http.ListenAndServe(":8080", nil)
}
