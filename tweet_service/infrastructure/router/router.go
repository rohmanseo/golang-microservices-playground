package router

import (
	"github.com/gin-gonic/gin"
	"github.com/rohmanseo/golang-clean-arch/controller"
)

func SetupRouter(r *gin.Engine, controller controller.ITweetController) {
	r.Use(gin.Logger())
	//todo : add auth middleware
	r.POST("/tweet/add", controller.Add)

}
