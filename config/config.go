package config

import (
	"github.com/spf13/viper"
	"os"

	"github.com/CasimirYang/entry-task-common/log"
)

func init() {
	viper.AddConfigPath("config")
	viper.SetConfigName("application") // name of config file (without extension)
	viper.SetConfigType("yaml")

	err := viper.ReadInConfig()
	if err != nil {
		log.SugarLogger.Error("read config failed: %v", err)
		os.Exit(1)
	}
}
