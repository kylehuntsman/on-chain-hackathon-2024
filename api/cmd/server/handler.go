package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

func generateHash(t Transaction) string {
	hash := sha256.Sum256([]byte(fmt.Sprintf("%v", t)))
	return hex.EncodeToString(hash[:])
}

func (db *Database) saveTransaction(t Transaction) (string, error) {
	hash := generateHash(t)
	_, err := db.Exec("INSERT INTO transactions (hash, chain, amount, address) VALUES (?, ?, ?, ?)", hash, t.Chain, t.Amount, t.Address)
	if err != nil {
		return "", err
	}
	return hash, nil
}

func (db *Database) getTransaction(hash string) (*Transaction, error) {
	row := db.QueryRow("SELECT chain, amount, address FROM transactions WHERE hash = ?", hash)
	t := &Transaction{}
	err := row.Scan(&t.Chain, &t.Amount, &t.Address)
	if err != nil {
		return nil, err
	}
	return t, nil
}
