package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/j4ckson185/motoboyscabana/appcomida-api/internal/handlers"
	_ "github.com/lib/pq"
)

func main() {
	db, err := sql.Open("postgres", os.Getenv("DB_CONN_STRING"))
	if err != nil {
		log.Fatal(err)
	}

	r := mux.NewRouter()
	r.HandleFunc("/comidas", handlers.GetComidas(db)).Methods("GET")
	r.HandleFunc("/comidas", handlers.CreateComida(db)).Methods("POST")

	http.Handle("/", r)
	log.Println("Starting server on :4003")
	log.Fatal(http.ListenAndServe(":4003", nil))
}
