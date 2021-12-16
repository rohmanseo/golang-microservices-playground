package router

import (
	"github.com/gin-gonic/gin"
	"github.com/rohmanseo/golang-clean-arch/controller"
)

func SetupRouter(r *gin.Engine, controller controller.IAuthController) {
	r.Use(gin.Logger())

	r.POST("/login", controller.Login)
	r.POST("/register", controller.Register)
	r.POST("/logout", controller.Logout)

}
