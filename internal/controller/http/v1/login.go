package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"glogin/internal/entity"
	"glogin/internal/usecase"
	"glogin/pkg/logger"
)

type loginRoutes struct {
	t usecase.Login
	l logger.Interface
}

func newLoginRoutes(handler *gin.RouterGroup, t usecase.Login, l logger.Interface) {
	r := &loginRoutes{t, l}

	h := handler.Group("/user")
	{
		h.POST("/login", r.login)
		h.POST("/loginwx", r.loginWx)
	}
}

type loginResponse struct {
	ErrCode int    `json:"errcode" example:"0 - success, 1 - username or password not correct"`
	Token   string `json:"token" example:"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHBpcmUiOjM2MDAsInBhc3N3b3JkIjoiMTIzNDU2IiwidXNlcm5hbWUiOiJhbGljZSJ9.u9Pha5vRrJ5meQasanfshl4hLBghLDzVF0rkX6ZdKLw"`
}

type doLoginRequest struct {
	Username string `json:"username" binding:"required"  example:"alice"`
	Password string `json:"password" binding:"required"  example:"123456"`
}

// @Summary     Login
// @Description Login system
// @ID          login
// @Tags  	    login
// @Accept      json
// @Produce     json
// @Param       request body doLoginRequest true "Login System"
// @Success     200 {object} loginResponse
// @Failure     400 {object} response
// @Failure     500 {object} response
// @Router      /user/login [post]
func (r *loginRoutes) login(c *gin.Context) {
	var request doLoginRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		r.l.Error(err, "http - v1 - doTranslate")
		errorResponse(c, http.StatusBadRequest, "invalid request body")

		return
	}

	errcode, token, err := r.t.Login(
		c.Request.Context(),
		entity.User{
			Username: request.Username,
			Password: request.Password,
		},
	)
	if err != nil {
		r.l.Error(err, "http - v1 - doLogin")
		errorResponse(c, http.StatusInternalServerError, "login service problems")

		return
	}

	c.JSON(http.StatusOK, loginResponse{ErrCode: errcode, Token: token})
}

// @Summary     LoginWx
// @Description Login system By Weixin
// @ID          loginWx
// @Tags  	    loginWx
// @Accept      json
// @Produce     json
// @Param       request body doLoginRequest true "Login System By Weixin"
// @Success     200 {object} loginResponse
// @Failure     400 {object} response
// @Failure     500 {object} response
// @Router      /user/login [post]
func (r *loginRoutes) loginWx(c *gin.Context) {
	var request doLoginRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		r.l.Error(err, "http - v1 - doTranslate")
		errorResponse(c, http.StatusBadRequest, "invalid request body")

		return
	}

	errcode, token, err := r.t.Login(
		c.Request.Context(),
		entity.User{
			Username: request.Username,
			Password: request.Password,
		},
	)
	if err != nil {
		r.l.Error(err, "http - v1 - doLogin")
		errorResponse(c, http.StatusInternalServerError, "login service problems")

		return
	}

	c.JSON(http.StatusOK, loginResponse{ErrCode: errcode, Token: token})
}
