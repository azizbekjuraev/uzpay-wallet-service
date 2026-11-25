// Package walletapp
package walletapp

import (
	"context"
	"uzpay/internal/domain/wallet"
)

type Service struct {
	repo wallet.Repository
}

func NewService(repo wallet.Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) Register(ctx context.Context, userID string) (*wallet.Wallet, error) {
	w, err := wallet.NewWallet(userID)
	if err != nil {
		return nil, err
	}
	err = s.repo.Save(ctx, w)
	if err != nil {
		return nil, err
	}
	return w, nil
}
