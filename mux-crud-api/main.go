package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/gorilla/mux"
)

func main() {

	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	router := mux.NewRouter()
	router.HandleFunc("/users", getUsers(db).Method("GET"))

}