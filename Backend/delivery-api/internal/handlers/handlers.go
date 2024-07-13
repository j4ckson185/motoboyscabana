package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"
)

type Delivery struct {
	ID         int    `json:"id"`
	Status     string `json:"status"`
	DriverName string `json:"driver_name"`
}

func GetDeliveries(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		rows, err := db.Query("SELECT id, status, driver_name FROM deliveries")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer rows.Close()

		deliveries := []Delivery{}
		for rows.Next() {
			var delivery Delivery
			if err := rows.Scan(&delivery.ID, &delivery.Status, &delivery.DriverName); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			deliveries = append(deliveries, delivery)
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(deliveries)
	}
}
