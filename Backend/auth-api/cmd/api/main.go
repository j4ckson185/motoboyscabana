package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"auth-api/internal/handlers"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

func main() {
    db, err := sql.Open("postgres", os.Getenv("DB_CONN_STRING"))
    if err != nil {
        log.Fatal(err)
    }

    r := mux.NewRouter()
    r.HandleFunc("/users", handlers.GetUsers(db)).Methods("GET")
    r.HandleFunc("/users", handlers.CreateUser(db)).Methods("POST")

    http.Handle("/", r)
    log.Println("Starting server on :4000")
    log.Fatal(http.ListenAndServe(":4000", nil))
}
