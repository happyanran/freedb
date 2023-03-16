package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/happyanran/freedb/server/middleware"
)

func Routers() *gin.Engine {
	Router := gin.Default()

	PublicGroup := Router.Group("")
	{
		// 健康监测
		PublicGroup.GET("/health", func(c *gin.Context) {
			c.JSON(http.StatusOK, "ok")
		})
	}
	{
		//systemRouter.InitInitRouter(PublicGroup) // 自动初始化相关
	}

	PrivateGroup := Router.Group("/api")
	PrivateGroup.Use(middleware.JWTAuth()).Use(middleware.CasbinHandler())
	{
		//systemRouter.InitApiRouter(PrivateGroup)                 // 注册功能api路由

	}

	return Router
}
