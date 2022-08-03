package usecase

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"guser/internal/entity"
	"guser/internal/usecase/repo"

	"github.com/huibunny/gocore/utils"
)

// wechat code info
type Code2Session struct {
	Code      string
	AppId     string
	AppSecret string
}

// pass
type Code2SessionResult struct {
	OpenId     string `json:"openid"`
	SessionKey string `json:"session_key"`
	UnionId    string `json:"unionid"`
	ErrCode    uint   `json:"errcode"`
	ErrMsg     string `json:"errmsg"`
}

// user unique index
type UserInfo struct {
	OpenId string `json:"openId"`
}

// UserUserCase -.
type UserUserCase struct {
	repo        *repo.UserRepo
	tokenExpire int64
	secret      string
	appid       string
	appSecret   string
}

// New -.
func New(r *repo.UserRepo, tokenExpire int64, secret, appid, appSecret string) *UserUserCase {
	return &UserUserCase{
		repo:        r,
		tokenExpire: tokenExpire,
		secret:      secret,
		appid:       appid,
		appSecret:   appSecret,
	}
}

// Login -.
func (uc *UserUserCase) Login(ctx context.Context, t entity.User) (int, string, error) {
	token := ""
	userID, errcode, err := uc.repo.CheckPass(context.Background(), t.Username, t.Password)
	if err != nil {
		return errcode, token, fmt.Errorf("UserUserCase - Login - s.repo.CheckPass: %w", err)
	}

	now := utils.CurrentTime()
	expireTime := now + uc.tokenExpire
	if errcode == 0 {
		token, err = utils.CreateToken(map[string]interface{}{
			"user_id":     userID,
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
func (uc *UserUserCase) LoginWx(ctx context.Context, code string) (int, string, error) {
	token := ""

	var err error

	var acsJson Code2SessionResult
	cs := Code2Session{
		Code:      code,
		AppId:     uc.appid,
		AppSecret: uc.appSecret,
	}
	api := "https://api.weixin.qq.com/sns/jscode2session?appid=%s&secret=%s&js_code=%s&grant_type=authorization_code"
	res, err := http.DefaultClient.Get(fmt.Sprintf(api, cs.AppId, cs.AppSecret, cs.Code))
	if err != nil {
		fmt.Println("weixin http request failed.")
		return -1, "", nil
	}
	if err := json.NewDecoder(res.Body).Decode(&acsJson); err != nil {
		fmt.Println("decoder error...")
		return -2, "", nil
	}

	var userId string
	var errcode int
	userId, errcode, err = uc.repo.CheckWxAccount(ctx, acsJson.OpenId, acsJson.UnionId)
	if err != nil {
		return errcode, token, fmt.Errorf("UserUserCase - Login - s.repo.checkWxAccount: %w", err)
	}

	now := utils.CurrentTime()
	expireTime := now + uc.tokenExpire
	if errcode == 0 {
		token, err = utils.CreateToken(map[string]interface{}{
			"user_id":     userId,
			"expire_time": expireTime,
		}, uc.secret)
		if err != nil {
			errcode = 2
		} else {
		}
	}

	return errcode, token, nil
}
