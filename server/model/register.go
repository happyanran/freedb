package model

import "github.com/happyanran/freedb/server/global"

func RegisterTables() {
	db := global.FDB_DB

	err := db.AutoMigrate()

	if err != nil {
		global.FDB_LOG.Panicf("register table failed: %v", err)
	}

	global.FDB_LOG.Info("register table success")
}
