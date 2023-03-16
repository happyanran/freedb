package middleware

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/happyanran/freedb/server/global"
	"github.com/happyanran/freedb/server/model/response"
	"github.com/happyanran/freedb/server/service"
)

var casbinService = new(service.CasbinService)

// CasbinHandler 拦截器
func CasbinHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		if global.FDB_CONFIG.System.Env != "develop" {
			waitUse, _ := service.GetClaims(c)
			//获取请求的PATH
			obj := c.Request.URL.Path
			// 获取请求方法
			act := c.Request.Method
			// 获取用户的角色
			sub := strconv.Itoa(int(waitUse.AuthorityId))

			e := casbinService.Casbin() // 判断策略中是否存在
			success, _ := e.Enforce(sub, obj, act)
			if !success {
				response.FailWithDetailed(gin.H{}, "权限不足", c)
				c.Abort()
				return
			}
		}
		c.Next()
	}
}
