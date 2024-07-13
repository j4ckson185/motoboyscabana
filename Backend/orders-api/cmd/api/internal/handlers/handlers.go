package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"
)

type Order struct {
	ID       int    `json:"id"`
	Status   string `json:"status"`
	Customer string `json:"customer"`
}

func GetOrders(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		rows, err := db.Query("SELECT id, status, customer FROM orders")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer rows.Close()

		orders := []Order{}
		for rows.Next() {
			var order Order
			if err := rows.Scan(&order.ID, &order.Status, &order.Customer); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			orders = append(orders, order)
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(orders)
	}
}
