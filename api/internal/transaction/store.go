package transaction

import (
	"database/sql"
	"fmt"

	"github.com/google/uuid"
)

type Transaction struct {
	Chain   string `json:"chain"`
	Amount  int    `json:"amount"`
	Address string `json:"address"`
}

type Store struct {
	DB       *sql.DB
	Host     string
	Port     int
	DBName   string
	User     string
	Password string
}

func (s *Store) SaveTransaction(t Transaction) (string, error) {
	uuid := uuid.New()
	fmt.Println(t)
	_, err := s.DB.Exec("INSERT INTO transactions (uuid, chain, amount, address) VALUES ($1, $2, $3, $4)", uuid, t.Chain, t.Amount, t.Address)
	if err != nil {
		return "", err
	}
	return uuid.String(), nil
}

func (s *Store) GetTransaction(id string) (*Transaction, error) {
	uuid, err := uuid.Parse(id)
	if err != nil {
		return nil, err
	}

	row := s.DB.QueryRow("SELECT chain, amount, address FROM transactions WHERE uuid = $1", uuid)
	t := &Transaction{}
	err = row.Scan(&t.Chain, &t.Amount, &t.Address)
	if err != nil {
		return nil, err
	}
	return t, nil
}
