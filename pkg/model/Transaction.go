package model

import "time"

type Transaction struct {
	ID              string
	TransactionType string
	Amount          float64
	CreatedAt       time.Time
}
