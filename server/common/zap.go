package common

import (
	"fmt"

	"github.com/happyanran/freedb/server/global"
	"go.uber.org/zap"
)

func Zap() *zap.SugaredLogger {
	config := zap.NewDevelopmentConfig()
	config.Level = zap.NewAtomicLevelAt(global.FDB_CONFIG.Zap.TransportLevel())

	logger, err := config.Build()
	if err != nil {
		panic(fmt.Errorf("Fatal error zap initial: %s \n", err))
	}

	return logger.Sugar()
}
