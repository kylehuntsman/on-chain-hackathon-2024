package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Hello, World!")

	config := Database{
		Host:     "localhost",
		Port:     5432,
		DBName:   "postgres",
		User:     "postgres",
		Password: "postgres",
	}

	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", config.Host, config.Port, config.User, config.Password, config.DBName)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	database := &Database{db, config.Host, config.Port, config.DBName, config.User, config.Password}
	_, err = database.Exec("CREATE TABLE IF NOT EXISTS transactions (hash TEXT PRIMARY KEY, chain TEXT, amount INTEGER, address TEXT)")
	if err != nil {
		panic(err)
	}

	http.HandleFunc("/transaction", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			var t Transaction
			if err := json.NewDecoder(r.Body).Decode(&t); err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			hash, err := database.saveTransaction(t)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			w.Write([]byte(hash))
		} else if r.Method == http.MethodGet {
			hash := r.URL.Query().Get("hash")
			t, err := database.getTransaction(hash)
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
