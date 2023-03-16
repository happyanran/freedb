package common

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/happyanran/freedb/server/global"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

// Gorm 初始化数据库并产生数据库全局变量
func Gorm() *gorm.DB {
	switch global.FDB_CONFIG.System.DbType {
	case "mysql":
		return GormMysql()
	// case "pgsql":
	// 	return GormPgSql()
	default:
		return GormMysql()
	}
}

type writer struct {
	logger.Writer
}

func NewWriter(w logger.Writer) *writer {
	return &writer{Writer: w}
}

func (w *writer) Printf(message string, data ...interface{}) {
	global.FDB_LOG.Info(fmt.Sprintf(message+"\n", data...))
}

func GormConfig(prefix string, singular bool, logMode string) *gorm.Config {
	var lm logger.LogLevel

	switch logMode {
	case "silent", "Silent":
		lm = logger.Silent
	case "error", "Error":
		lm = logger.Error
	case "warn", "Warn":
		lm = logger.Warn
	case "info", "Info":
		lm = logger.Info
	default:
		lm = logger.Info
	}

	return &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   prefix,
			SingularTable: singular,
		},
		DisableForeignKeyConstraintWhenMigrating: true,
		Logger: logger.New(NewWriter(log.New(os.Stdout, "\r\n", log.LstdFlags)), logger.Config{
			SlowThreshold: 200 * time.Millisecond,
			LogLevel:      lm,
			Colorful:      false,
		}),
	}
}
