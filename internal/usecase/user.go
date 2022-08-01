package usecase

import (
	"context"
	"fmt"

	"glogin/internal/entity"
	"glogin/internal/usecase/repo"

	"github.com/huibunny/gocore/utils"
)

// UserUserCase -.
type UserUserCase struct {
	repo        *repo.UserRepo
	tokenExpire uint
	secret      string
}

// New -.
func New(r *repo.UserRepo, tokenExpire uint, secret string) *UserUserCase {
	return &UserUserCase{
		repo:        r,
		tokenExpire: tokenExpire,
		secret:      secret,
	}
}

// Login -.
func (uc *UserUserCase) Login(ctx context.Context, t entity.User) (int, string, error) {
	token := ""
	errcode, err := uc.repo.CheckPass(context.Background(), t.Username, t.Password)
	if err != nil {
		return errcode, token, fmt.Errorf("UserUserCase - Login - s.repo.Store: %w", err)
	}

	if errcode == 0 {
		token, err = utils.CreateToken(map[string]interface{}{
			"username": t.Username,
			"password": t.Password,
			"expire":   uc.tokenExpire,
		}, uc.secret)
		if err != nil {
			errcode = 2
		} else {
		}
	}

	return errcode, token, nil
}
