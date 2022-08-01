package v1

import (
	"github.com/gin-gonic/gin"
)

type response struct {
	Message string `json:"message" example:"success"`
}

func errorResponse(c *gin.Context, code int, msg string) {
	c.AbortWithStatusJSON(code, response{Message: msg})
}
