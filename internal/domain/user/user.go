// Package user
package user

import (
	"errors"
	"net/mail"
	"strings"

	"github.com/google/uuid"
)

type Email struct {
	value string
}

func NewEmail(address string) (Email, error) {
	address = strings.TrimSpace(address)
	if _, err := mail.ParseAddress(address); err != nil {
		return Email{}, errors.New("invalid email format")
	}
	return Email{value: address}, nil
}

func (e Email) String() string {
	return e.value
}

type User struct {
	id          uuid.UUID
	username    string
	email       Email
	password    string
	phoneNumber string
}

func NewUser(username string, emailStr string, phoneNumber string, hashedPassword string) (*User, error) {
	username = strings.TrimSpace(username)
	if len(username) < 3 {
		return nil, errors.New("username must be at least 3 characters long")
	}

	email, err := NewEmail(emailStr)
	if err != nil {
		return nil, err
	}

	if hashedPassword == "" {
		return nil, errors.New("password cannot be empty")
	}

	return &User{
		id:          uuid.New(),
		username:    username,
		email:       email,
		password:    hashedPassword,
		phoneNumber: phoneNumber,
	}, nil
}

func (u *User) ID() uuid.UUID {
	return u.id
}

func (u *User) Username() string {
	return u.username
}

func (u *User) Email() string {
	return u.email.String()
}

func (u *User) ChangeEmail(newEmailStr string) error {
	newEmail, err := NewEmail(newEmailStr)
	if err != nil {
		return err
	}
	u.email = newEmail
	return nil
}

func (u *User) Password() string {
	return u.password
}
