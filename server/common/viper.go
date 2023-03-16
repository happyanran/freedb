package common

import (
	"flag"
	"fmt"

	"github.com/happyanran/freedb/server/global"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

func Viper() *viper.Viper {
	var config string

	flag.StringVar(&config, "c", "./config.yaml", "path to config file.")
	flag.Parse()

	v := viper.New()
	v.SetConfigFile(config)
	v.SetConfigType("yaml")

	if err := v.ReadInConfig(); err != nil {
		panic(fmt.Errorf("Fatal error read config file: %s \n", err))
	}

	if err := v.Unmarshal(&global.FDB_CONFIG); err != nil {
		fmt.Println(err)
	}

	v.WatchConfig()

	v.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("config file changed:", e.Name)

		if err := v.Unmarshal(&global.FDB_CONFIG); err != nil {
			fmt.Println(err)
		}
	})

	return v
}
