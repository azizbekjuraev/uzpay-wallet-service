// Package ledger entry
package ledger

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

type LedgerEntry struct {
	ledgerEntryID uuid.UUID
	userID        int
	amount        int64
	description   string
	createdAt     time.Time
}

func NewLedgerEntry(userID int, amount int64, description string) (*LedgerEntry, error) {
	if userID == 0 {
		return nil, errors.New("must create user first before making any operations")
	}

	return &LedgerEntry{
		ledgerEntryID: uuid.New(),
		userID:        userID,
		amount:        amount,
		description:   description,
		createdAt:     time.Now(),
	}, nil
}
