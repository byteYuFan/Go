package main

import (
	"github.com/LingshengCode/awesomeProject/controller"
	"github.com/gin-gonic/gin"
)

func initRouter(r *gin.Engine) {
	r.Static("/static", "./public")

	apiRouter := r.Group("/douyin")
	// 基础接口
	apiRouter.POST("/user/register/", controller.Register)
	apiRouter.POST("/user/login/", controller.Login)
}
