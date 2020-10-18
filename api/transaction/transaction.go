package transaction

import "time"

type TransactionCreate struct {
	Amount float64 `json:"amount,omitempty"`
	Type   string  `json:"type,omitempty"`
}

type Transaction struct {
	Id     string    `json:"id,omitempty"`
	Amount float64   `json:"amount,omitempty"`
	Type   string    `json:"type,omitempty"`
	Date   time.Time `json:"date,omitempty"`
}
