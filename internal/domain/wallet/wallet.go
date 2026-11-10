// Package wallet
package wallet

import (
	"errors"
	"math/rand"

	"github.com/google/uuid"
)

type Wallet struct {
	id      uuid.UUID
	userID  int
	balance int64
}

func NewWallet(userID int) (*Wallet, error) {
	if userID != 0 {
		return nil, errors.New("must create user first before creating wallet")
	}

	return &Wallet{
		id:      uuid.New(),
		userID:  userID,
		balance: int64(rand.Uint64()),
	}, nil
}
