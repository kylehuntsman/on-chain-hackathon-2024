package main

import "database/sql"

type Transaction struct {
	Chain   string `json:"chain"`
	Amount  int    `json:"amount"`
	Address string `json:"address"`
}

type Database struct {
	*sql.DB
	Host     string
	Port     int
	DBName   string
	User     string
	Password string
}
