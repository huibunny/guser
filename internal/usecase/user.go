package usecase

import (
	"context"
	"fmt"

	"guser/internal/entity"
	"guser/internal/usecase/repo"

	"github.com/huibunny/gocore/utils"
)

// UserUserCase -.
type UserUserCase struct {
	repo        *repo.UserRepo
	tokenExpire int64
	secret      string
}

// New -.
func New(r *repo.UserRepo, tokenExpire int64, secret string) *UserUserCase {
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

	now := utils.CurrentTime()
	expireTime := now + uc.tokenExpire
	if errcode == 0 {
		token, err = utils.CreateToken(map[string]interface{}{
			"username":    t.Username,
			"password":    t.Password,
			"expire_time": expireTime,
		}, uc.secret)
		if err != nil {
			errcode = 2
		} else {
		}
	}

	return errcode, token, nil
}

// Login -.
func (uc *UserUserCase) LoginWx(ctx context.Context, t entity.User) (int, string, error) {
	token := ""
	errcode, err := uc.repo.CheckPass(context.Background(), t.Username, t.Password)
	if err != nil {
		return errcode, token, fmt.Errorf("UserUserCase - Login - s.repo.Store: %w", err)
	}

	if errcode == 0 {
		token, err = utils.CreateToken(map[string]interface{}{
			"username":    t.Username,
			"password":    t.Password,
			"expire_time": utils.CurrentTime() + uc.tokenExpire,
		}, uc.secret)
		if err != nil {
			errcode = 2
		} else {
		}
	}

	return errcode, token, nil
}
