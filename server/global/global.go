package global

import (
	"github.com/happyanran/freedb/server/config"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"golang.org/x/sync/singleflight"
	"gorm.io/gorm"
)

var (
	FDB_CONFIG       config.Server
	FDB_VP           *viper.Viper
	FDB_LOG          *zap.SugaredLogger
	FDB_DB           *gorm.DB
	FDB_Singleflight = &singleflight.Group{}
)
