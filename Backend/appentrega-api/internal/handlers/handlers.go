package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"
)

type Entrega struct {
	ID             int    `json:"id"`
	Status         string `json:"status"`
	Entregador     string `json:"entregador"`
	NomeEntregador string `json:"nome_entregador"`
}

func GetEntregas(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		rows, err := db.Query("SELECT id, status, entregador, nome_entregador FROM entregas")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer rows.Close()

		entregas := []Entrega{}
		for rows.Next() {
			var entrega Entrega
			if err := rows.Scan(&entrega.ID, &entrega.Status, &entrega.Entregador, &entrega.NomeEntregador); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			entregas = append(entregas, entrega)
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(entregas)
	}
}
