package main

import (
    "database/sql"
    "encoding/json"
    "log"
    "net/http"
    "os"

    "github.com/gorilla/mux"
    _ "github.com/lib/pq"
)

var db *sql.DB

func main() {
    var err error
    connStr := os.Getenv("DB_CONN_STRING")
    db, err = sql.Open("postgres", connStr)
    if err != nil {
        log.Fatal(err)
    }

    r := mux.NewRouter()
    r.HandleFunc("/health", HealthCheckHandler)
    r.HandleFunc("/restaurants", GetRestaurants).Methods("GET")
    r.HandleFunc("/restaurants/{id}", GetRestaurant).Methods("GET")

    http.Handle("/", r)
    log.Println("Starting server on :4005")
    log.Fatal(http.ListenAndServe(":4005", nil))
}

func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
    w.WriteHeader(http.StatusOK)
    w.Write([]byte("WebRestaurante API is running"))
}

func GetRestaurants(w http.ResponseWriter, r *http.Request) {
    rows, err := db.Query("SELECT id, name, address FROM restaurants")
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    defer rows.Close()

    restaurants := []Restaurant{}
    for rows.Next() {
        var restaurant Restaurant
        if err := rows.Scan(&restaurant.ID, &restaurant.Name, &restaurant.Address); err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }
        restaurants = append(restaurants, restaurant)
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(restaurants)
}

func GetRestaurant(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id := vars["id"]

    var restaurant Restaurant
    err := db.QueryRow("SELECT id, name, address FROM restaurants WHERE id = $1", id).Scan(&restaurant.ID, &restaurant.Name, &restaurant.Address)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(restaurant)
}

type Restaurant struct {
    ID      int    `json:"id"`
    Name    string `json:"name"`
    Address string `json:"address"`
}
