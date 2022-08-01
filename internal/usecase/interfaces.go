// Package usecase implements application business logic. Each logic group in own file.
package usecase

import (
	"context"

	"glogin/internal/entity"
)

//go:generate mockgen -source=interfaces.go -destination=./mocks_test.go -package=usecase_test

type (
	// Login -.
	Login interface {
		Login(context.Context, entity.User) (int, string, error)
	}

	// UserRepo -.
	UserRepo interface {
		Login(context.Context) error
	}
)
