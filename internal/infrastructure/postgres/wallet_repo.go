// Package postgres
package postgres

import (
	"context"
	"errors"
	"fmt"
	"uzpay/internal/domain/wallet"

	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
)

type WalletRepository struct {
	db *pgxpool.Pool
}

func NewWalletRepository(db *pgxpool.Pool) *WalletRepository {
	return &WalletRepository{db: db}
}

var _ wallet.Repository = (*WalletRepository)(nil)

func (r *WalletRepository) Save(ctx context.Context, w *wallet.Wallet) error {
	var query = `
		INSERT INTO wallets (wallet_id, user_id, balance)
		VALUES($1, $2, $3)
	`
	_, err := r.db.Exec(ctx, query, w.ID(), w.UserID(), w.Balance())
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == "23505" {
			return wallet.ErrWalletExists
		}
		return fmt.Errorf("save wallet: %w", err)
	}
	return nil
}
