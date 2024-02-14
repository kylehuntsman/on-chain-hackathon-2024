package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	_ "github.com/lib/pq"

	"github.com/kylehuntsman/on-chain-hackathon-2024/internal/transaction"
)

func main() {
	fmt.Println("Hello, World!")

	store := transaction.Store{
		Host:     "localhost",
		Port:     5432,
		DBName:   "postgres",
		User:     "postgres",
		Password: "postgres",
	}

	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", store.Host, store.Port, store.User, store.Password, store.DBName)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	store.DB = db

	_, err = store.DB.Exec("CREATE TABLE IF NOT EXISTS transactions (uuid TEXT PRIMARY KEY, chain TEXT, amount INTEGER, address TEXT)")
	if err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/transaction", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			var t transaction.Transaction
			if err := json.NewDecoder(r.Body).Decode(&t); err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			uuid, err := store.SaveTransaction(t)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			w.Write([]byte(uuid))
		} else if r.Method == http.MethodGet {
			uuid := r.URL.Query().Get("uuid")
			t, err := store.GetTransaction(uuid)
			if err != nil {
				http.Error(w, err.Error(), http.StatusNotFound)
				return
			}
			json.NewEncoder(w).Encode(t)
		} else {
			http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		}
	})

	http.ListenAndServe(":8080", nil)
}
