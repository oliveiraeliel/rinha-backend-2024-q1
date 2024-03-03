package main

import (
	"context"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/oliveiraeliel/rinha-backend-24-q1/internal/http/router"
	"github.com/oliveiraeliel/rinha-backend-24-q1/internal/transaction"
)

func main() {
	db, err := pgxpool.Connect(
		context.Background(),
		"postgresql://postgres:123@db:5432/rinha",
	)

	if err != nil {
		panic(err)
	}

	defer db.Close()

	repo := transaction.NewTransactionRepository(db)
	router.Run(repo)
}
