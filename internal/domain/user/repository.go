// Package user
package user

import (
	"context"
	"errors"
)

var (
	ErrNotFound   = errors.New("user not found")
	ErrEmailTaken = errors.New("email already taken")
)

type Repository interface {
	Save(ctx context.Context, u *User) error
}
