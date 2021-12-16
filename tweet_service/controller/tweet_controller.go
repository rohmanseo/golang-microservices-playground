package controller

import "github.com/gin-gonic/gin"

type ITweetController interface {
	Add(ctx *gin.Context)
}
