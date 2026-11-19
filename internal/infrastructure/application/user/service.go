// Package user
package userapp

import (
	"context"
	"uzpay/internal/domain/user"
)

type Service struct {
	repo   user.Repository
	hasher user.PasswordHasher
}

func NewService(repo user.Repository, hasher user.PasswordHasher) *Service {
	return &Service{repo: repo, hasher: hasher}
}

func (s *Service) Register(ctx context.Context, username string, email string, plainPassword string, phoneNumber string) error {
	hash, _ := s.hasher.Hash(plainPassword)
	u, err := user.NewUser(username, email, phoneNumber, hash)
	if err != nil {
		return err
	}
	err = s.repo.Save(ctx, u)
	if err != nil {
		return err
	}
	return err
}
