package repo

import (
	"context"
	"fmt"

	"guser/pkg/postgres"
)

const _defaultEntityCap = 64

// UserRepo -.
type UserRepo struct {
	*postgres.Postgres
}

// New -.
func New(pg *postgres.Postgres) *UserRepo {
	return &UserRepo{pg}
}

// CheckPass -.
func (r *UserRepo) CheckPass(ctx context.Context, username, password string) (int, error) {
	sql, _, err := r.Builder.
		Select("count(1)").
		From("t_user").
		Where("username=? and password=?").ToSql()
	if err != nil {
		return -1, fmt.Errorf("UserRepo - checkPass - r.Builder: %w", err)
	}

	var count int
	r.Pool.QueryRow(ctx, sql, username, password).Scan(&count)
	if err != nil {
		return -1, fmt.Errorf("UserRepo - checkPass - r.Pool.Query: %w", err)
	}

	errcode := 1
	if count > 0 {
		errcode = 0
	}
	return errcode, nil
}
