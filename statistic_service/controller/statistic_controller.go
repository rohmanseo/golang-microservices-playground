package controller

import "github.com/gin-gonic/gin"

type IStatisticController interface {
	GetStatistic(ctx *gin.Context)
}
