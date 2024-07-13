package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"delivery-api/internal/handlers"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

func main() {
	db, err := sql.Open("postgres", os.Getenv("DB_CONN_STRING"))
	if err != nil {
		log.Fatal(err)
	}

	r := mux.NewRouter()
	r.HandleFunc("/deliveries", handlers.GetDeliveries(db)).Methods("GET")

	http.Handle("/", r)
	log.Println("Starting server on :4006")
	log.Fatal(http.ListenAndServe(":4006", nil))
}
