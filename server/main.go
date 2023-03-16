package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/happyanran/freedb/server/common"
	"github.com/happyanran/freedb/server/global"
	"github.com/happyanran/freedb/server/model"
	"github.com/happyanran/freedb/server/router"
)

func main() {
	global.FDB_VP = common.Viper() // 初始化Viper
	global.FDB_LOG = common.Zap()  // 初始化zap日志库
	global.FDB_DB = common.Gorm()  // gorm连接数据库
	if global.FDB_DB != nil {
		model.RegisterTables() // 初始化表
		db, _ := global.FDB_DB.DB()
		defer db.Close()
	}

	RunServer()
}

func RunServer() {
	if global.FDB_CONFIG.System.Env != "dev" {
		gin.SetMode(gin.ReleaseMode)
	}

	Router := router.Routers()

	address := fmt.Sprintf(":%d", global.FDB_CONFIG.System.Addr)
	s := &http.Server{
		Addr:           address,
		Handler:        Router,
		ReadTimeout:    20 * time.Second,
		WriteTimeout:   20 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	global.FDB_LOG.Error(s.ListenAndServe().Error())
}
