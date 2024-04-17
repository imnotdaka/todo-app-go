package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Ping() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "created")
	}
}
