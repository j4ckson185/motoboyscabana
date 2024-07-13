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
    r.HandleFunc("/deliveries", GetDeliveries).Methods("GET")
    r.HandleFunc("/deliveries/{id}", GetDelivery).Methods("GET")

    http.Handle("/", r)
    log.Println("Starting server on :4004")
    log.Fatal(http.ListenAndServe(":4004", nil))
}

func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
    w.WriteHeader(http.StatusOK)
    w.Write([]byte("APPEntrega API is running"))
}

func GetDeliveries(w http.ResponseWriter, r *http.Request) {
    rows, err := db.Query("SELECT id, status, driver FROM deliveries")
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    defer rows.Close()

    deliveries := []Delivery{}
    for rows.Next() {
        var delivery Delivery
        if err := rows.Scan(&delivery.ID, &delivery.Status, &delivery.Driver); err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }
        deliveries = append(deliveries, delivery)
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(deliveries)
}

func GetDelivery(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id := vars["id"]

    var delivery Delivery
    err := db.QueryRow("SELECT id, status, driver FROM deliveries WHERE id = $1", id).Scan(&delivery.ID, &delivery.Status, &delivery.Driver)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(delivery)
}

type Delivery struct {
    ID       int    `json:"id"`
    Status   string `json:"status"`
    Driver   string `json:"driver"`
}
