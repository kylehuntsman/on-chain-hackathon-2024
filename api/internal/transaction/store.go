package transaction

import (
	"database/sql"

	"github.com/google/uuid"
)

type Transaction struct {
	Chain   string `json:"chain"`
	Amount  int    `json:"amount"`
	Address string `json:"address"`
}

type Store struct {
	*sql.DB
	Host     string
	Port     int
	DBName   string
	User     string
	Password string
}

func (ts *Store) SaveTransaction(t Transaction) (string, error) {
	uuid := uuid.New()
	_, err := ts.Exec("INSERT INTO transactions (uuid, chain, amount, address) VALUES (?, ?, ?, ?)", uuid, t.Chain, t.Amount, t.Address)
	if err != nil {
		return "", err
	}
	return uuid.String(), nil
}

func (ts *Store) GetTransaction(id string) (*Transaction, error) {
	uuid, err := uuid.Parse(id)
	if err != nil {
		return nil, err
	}

	row := ts.QueryRow("SELECT chain, amount, address FROM transactions WHERE uuid = ?", uuid)
	t := &Transaction{}
	err = row.Scan(&t.Chain, &t.Amount, &t.Address)
	if err != nil {
		return nil, err
	}
	return t, nil
}
