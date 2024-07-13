package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/j4ckson185/motoboyscabana/orders-api/internal/handlers"
	_ "github.com/lib/pq"
)

func main() {
	db, err := sql.Open("postgres", os.Getenv("DB_CONN_STRING"))
	if err != nil {
		log.Fatal(err)
	}

	r := mux.NewRouter()
	r.HandleFunc("/orders", handlers.GetOrders(db)).Methods("GET")

	http.Handle("/", r)
	log.Println("Starting server on :4001")
	log.Fatal(http.ListenAndServe(":4001", nil))
}
