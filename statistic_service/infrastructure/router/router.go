package router

import (
	"github.com/gin-gonic/gin"
	"github.com/rohmanseo/golang-clean-arch/controller"
)

func SetupRouter(r *gin.Engine, controller controller.IStatisticController) {
	r.Use(gin.Logger())
	r.GET("/statistic", controller.GetStatistic)
}
