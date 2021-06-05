package config

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

var Cfg *viper.Viper

func init() {
	Cfg = viper.New()
	Cfg.AddConfigPath(".")
	Cfg.SetConfigName("config")
	Cfg.SetConfigType("yaml")

	err := Cfg.ReadInConfig()

	if err != nil {
		panic(fmt.Errorf("读取配置文件失败: %s \n", err))
	}

	Cfg.WatchConfig()
	Cfg.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("Config file changed:", e.Name)
	})
}
