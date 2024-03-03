package transaction

import (
	"context"
	"errors"
	"time"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/oliveiraeliel/rinha-backend-24-q1/internal/domain"
)

type TransactionRepository interface {
	SaveTransaction(ctx context.Context, t *domain.Transaction) (domain.CreateTransactionResponse, error)
	GetExtract(ctx context.Context, clientId int) (domain.Extract, error)
}

type repository struct {
	db    *pgxpool.Pool
}

func NewTransactionRepository(db *pgxpool.Pool) TransactionRepository {
	return &repository{
		db: db,
	}
}

func (r *repository) SaveTransaction(ctx context.Context, t *domain.Transaction) (domain.CreateTransactionResponse, error) {
	var operation string
	if t.Type == "c" {
		operation = "creditar"
	} else if t.Type == "d" {
		operation = "debitar"
	} else {
		return domain.CreateTransactionResponse{}, errors.New("invalid type")
	}

	var novo_saldo int
	var limite int
	var possui_erro bool
	var mensagem string

	err := r.db.QueryRow(
		ctx,
		`
		SELECT novo_saldo, limite, possui_erro, mensagem
		FROM public.`+operation+`($1, $2, $3)`,
		t.ClientID,
		int(t.Value),
		t.Description,
	).Scan(&novo_saldo, &limite, &possui_erro, &mensagem)

	if err != nil {
		return domain.CreateTransactionResponse{}, err
	}

	if possui_erro {
		return domain.CreateTransactionResponse{}, errors.New(mensagem)
	}

	return domain.CreateTransactionResponse{Balance: novo_saldo, Limit: limite}, nil
}

func (r *repository) getClientBalance(ctx context.Context, clientId int) (domain.ExtractHeader, bool) {
	var saldo int
	var limite int

	err := r.db.QueryRow(
		ctx,
		`SELECT saldo, limite FROM clientes WHERE id=$1;`,
		clientId,
	).Scan(&saldo, &limite)

	if err != nil {
		return domain.ExtractHeader{}, false
	}

	extractHeader := domain.ExtractHeader{
		GeneratedAt: time.Now(),
		Limit:       limite,
		Total:     saldo,
	}

	return extractHeader, true
}

func (r *repository) GetExtract(ctx context.Context, clientId int) (domain.Extract, error) {
	extractHeader, ok := r.getClientBalance(ctx, clientId)

	if !ok {
		return domain.Extract{}, errors.New("client not found")
	}

	rows, err := r.db.Query(
		ctx,
		`SELECT valor, tipo, descricao, realizada_em
		FROM transacoes
		WHERE cliente_id=$1
		ORDER BY realizada_em DESC LIMIT 10;`,
		clientId,
	)

	if err != nil {
		return domain.Extract{}, err
	}

	extract := domain.NewExtact(extractHeader)

	for rows.Next() {
		var transaction domain.LastTransaction
		rows.Scan(
			&transaction.Value,
			&transaction.Type,
			&transaction.Description,
			&transaction.CreatedAt,
		)
		extract.LastTransactions = append(extract.LastTransactions, transaction)
	}

	return extract, nil
}
