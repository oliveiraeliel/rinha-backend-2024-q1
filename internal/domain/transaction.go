package domain

import "time"

type Transaction struct {
	ID          int       `json:"id"`
	ClientID    int       `json:"cliente_id"`
	Name        string    `json:"nome"`
	Value       int       `json:"valor"`
	Description string    `json:"decricao"`
	Type        string    `json:"tipo"`
	CreatedAt   time.Time `json:"realizado_em"`
}

type CreateTransactionResponse struct {
	Balance int `json:"saldo"`
	Limit   int `json:"limite"`
}
