// Package wallet
package wallet

import (
	"context"
	"errors"
)

var (
	ErrWalletExists = errors.New("wallet already exists")
)

type Repository interface {
	Save(ctx context.Context, w *Wallet) error
}
