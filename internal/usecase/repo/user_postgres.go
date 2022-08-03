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
func (r *UserRepo) CheckPass(ctx context.Context, username, password string) (string, int, error) {
	sql, _, err := r.Builder.
		Select("id").
		From("t_user").
		Where("username=? and password=?").ToSql()
	if err != nil {
		return "", -1, fmt.Errorf("UserRepo - checkPass - r.Builder: %w", err)
	}

	userId := ""
	err = r.Pool.QueryRow(ctx, sql, username, password).Scan(&userId)
	if err != nil {
		return "", -2, fmt.Errorf("UserRepo - checkPass - r.Pool.Query: %w", err)
	}

	errcode := 1
	if len(userId) > 0 {
		errcode = 0
	}
	return userId, errcode, nil
}

// CheckWxAccount
func (r *UserRepo) CheckWxAccount(ctx context.Context, openid, unionid string) (string, int, error) {
	userID := ""
	sql, _, err := r.Builder.Select("count(1)").From("t_wxaccount").Where("openid=?").ToSql()
	if err != nil {
		return "", -1, fmt.Errorf("UserRepo - checkWxAccount - r.Builder: %w", err)
	}

	count := 0
	err = r.Pool.QueryRow(ctx, sql, openid).Scan(&count)
	if err != nil {
		return "", -2, fmt.Errorf("UserRepo - checkWxAccount - r.Pool.Query: %w", err)
	}

	errcode := 1
	if count > 0 {
		errcode = 0
	} else {
		insertWxAccountSql, _, err := r.Builder.Insert("t_wxaccount").Columns("openid", "unionid").Values("?", "?").Suffix("RETURNING id").ToSql()
		if err != nil {
			return "", -1, fmt.Errorf("UserRepo - checkWxAccount - r.Builder: %w", err)
		}
		insertUserSql, _, err := r.Builder.Insert("t_user").Columns("username", "password").Values("?", "?").Suffix("RETURNING id").ToSql()
		if err != nil {
			return "", -1, fmt.Errorf("UserRepo - checkWxAccount - r.Builder: %w", err)
		}
		insertUserWxAccountSql, _, err := r.Builder.Insert("t_user_wxaccount").Columns("user_id", "wx_accountid").Values("?", "?").ToSql()
		if err != nil {
			return "", -1, fmt.Errorf("UserRepo - checkWxAccount - r.Builder: %w", err)
		}
		tx, err := r.Pool.Begin(ctx)
		if err != nil {
			return "", -3, fmt.Errorf("UserRepo - checkWxAccount - r.Pool.Begin: %w", err)
		}
		wxAccountID := ""
		err = tx.QueryRow(ctx, insertWxAccountSql, openid, unionid).Scan(&wxAccountID)
		if err != nil {
			tx.Rollback(ctx)
			return "", -4, fmt.Errorf("UserRepo - checkWxAccount - tx.QueryRow.Scan: %w", err)
		}
		userID = ""
		err = tx.QueryRow(ctx, insertUserSql, openid, openid).Scan(&userID)
		if err != nil {
			tx.Rollback(ctx)
			return "", -4, fmt.Errorf("UserRepo - checkWxAccount - tx.QueryRow.Scan: %w", err)
		}
		_, err = tx.Exec(ctx, insertUserWxAccountSql, userID, wxAccountID)
		if err != nil {
			tx.Rollback(ctx)
			return "", -4, fmt.Errorf("UserRepo - checkWxAccount - tx.Exec: %w", err)
		}
		tx.Commit(ctx)
	}

	return userID, errcode, nil
}
