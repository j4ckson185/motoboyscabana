package handlers

import (
    "database/sql"
    "encoding/json"
    "net/http"
)

type Order struct {
    ID            int     `json:"id"`
    UserID        int     `json:"user_id"`
    Customer      string  `json:"customer"`
    Item          string  `json:"item"`
    Quantity      int     `json:"quantity"`
    Total         float64 `json:"total"`
    PaymentMethod string  `json:"payment_method"`
    Status        string  `json:"status"`
    CreatedAt     string  `json:"created_at"`
}

func GetOrders(db *sql.DB) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        rows, err := db.Query("SELECT id, user_id, customer, item, quantity, total, payment_method, status, created_at FROM orders")
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }
        defer rows.Close()

        orders := []Order{}
        for rows.Next() {
            var order Order
            err := rows.Scan(&order.ID, &order.UserID, &order.Customer, &order.Item, &order.Quantity, &order.Total, &order.PaymentMethod, &order.Status, &order.CreatedAt)
            if err != nil {
                http.Error(w, err.Error(), http.StatusInternalServerError)
                return
            }
            orders = append(orders, order)
        }

        w.Header().Set("Content-Type", "application/json")
        json.NewEncoder(w).Encode(orders)
    }
}
