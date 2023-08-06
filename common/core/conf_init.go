package core

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"mini-tiktok/common/config"
	"mini-tiktok/common/global"
)

const filename = "./settings.yaml"

func InitConf() {
	var c = &config.Config{}

	viper.SetConfigFile(filename)
	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("viper.ReadInConfig Failed, err : ", err.Error())
		return
	}
	if err := viper.Unmarshal(c); err != nil {
		fmt.Println("viper.Unmarshal Failed, err : ", err.Error())
		return
	}

	viper.WatchConfig()
	viper.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println("配置文件配修改了")
		if err := viper.Unmarshal(c); err != nil {
			fmt.Println("viper.Unmarshal Failed, err : ", err.Error())
			return
		}
	})

	global.Config = c
	return
}
