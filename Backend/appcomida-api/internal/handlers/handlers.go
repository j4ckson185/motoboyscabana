package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"
)

type Comida struct {
	ID    int     `json:"id"`
	Nome  string  `json:"nome"`
	Preco float64 `json:"preco"`
}

func GetComidas(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		rows, err := db.Query("SELECT id, nome, preco FROM comidas")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer rows.Close()

		comidas := []Comida{}
		for rows.Next() {
			var comida Comida
			if err := rows.Scan(&comida.ID, &comida.Nome, &comida.Preco); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			comidas = append(comidas, comida)
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(comidas)
	}
}

func CreateComida(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var comida Comida
		if err := json.NewDecoder(r.Body).Decode(&comida); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		err := db.QueryRow(
			"INSERT INTO comidas (nome, preco) VALUES ($1, $2) RETURNING id",
			comida.Nome, comida.Preco,
		).Scan(&comida.ID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(comida)
	}
}
