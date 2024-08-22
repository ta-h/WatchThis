package main

import (
	"net/http"

	"github.com/go-playground/validator/v10"
)

var (
	validate           *validator.Validate
	watchListIdCounter = 68
	watchLists         = make(map[string]*WatchList)
)

type WatchList struct {
	ID   string  `json:"id" db:"id"`
	Name *string `json:"name" db:"name"`
}

func main() {
	validate = validator.New(validator.WithRequiredStructEnabled())

	// dbFile := "watchthis.db"
	// createDB(dbFile)
	// createTable(dbFile)

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
