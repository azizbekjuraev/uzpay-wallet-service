// Package postgres
package postgres

import (
	"context"
	"errors"
	"fmt"
	"uzpay/internal/domain/user"

	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
)

type UserRepository struct {
	db *pgxpool.Pool
}

func NewUserRepository(db *pgxpool.Pool) *UserRepository {
	return &UserRepository{db: db}
}

var _ user.Repository = (*UserRepository)(nil)

func (r *UserRepository) Save(ctx context.Context, u *user.User) error {
	var query = `
		INSERT INTO users (user_id, username, email, password, phone_number)
		VALUES($1, $2, $3, $4, $5)
	`
	_, err := r.db.Exec(ctx, query, u.ID(), u.Username(), u.Email(), u.Password(), u.PhoneNumber())
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == "23505" {
			return user.ErrEmailTaken
		}
		return fmt.Errorf("save user: %w", err)
	}
	return nil
}
