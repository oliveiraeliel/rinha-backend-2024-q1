package domain

import "time"

type Extract struct {
	ExtractHeader    ExtractHeader     `json:"saldo"`
	LastTransactions []LastTransaction `json:"ultimas_transacoes"`
}

type ExtractHeader struct {
	Total       int       `json:"total"`
	Limit       int       `json:"limite"`
	GeneratedAt time.Time `json:"data_extrato"`
}

type LastTransaction struct {
	Value       int       `json:"valor"`
	Type        string    `json:"tipo"`
	Description string    `json:"descricao"`
	CreatedAt   time.Time `json:"realizado_em"`
}

func NewExtact(header ExtractHeader) Extract {
	return Extract{ExtractHeader: header, LastTransactions: make([]LastTransaction, 0)}
}
