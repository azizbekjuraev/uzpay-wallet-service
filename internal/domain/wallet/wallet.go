// Package wallet
package wallet

import (
	"errors"

	"github.com/google/uuid"
)

type Wallet struct {
	id      uuid.UUID
	userID  string
	balance int64
}

func NewWallet(userID string) (*Wallet, error) {
	if userID == "" {
		return nil, errors.New("must create user first before creating wallet")
	}

	return &Wallet{
		id:     uuid.New(),
		userID: userID,
	}, nil
}

func (w *Wallet) ID() uuid.UUID {
	return w.id
}

func (w *Wallet) UserID() string {
	return w.userID
}

func (w *Wallet) Balance() int64 {
	return w.balance
}
