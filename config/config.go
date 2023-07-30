package config

import (
	"fmt"

	"github.com/spf13/viper"
)

// 读取config配置
func InitConfig() {
	viper.AddConfigPath("config")
	viper.SetConfigName("douyin")
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println("Read Config failed: ", err)
	}
	fmt.Println("Douyin config inited!")
}

// 填充Mysql连接字符串
func DBConnectString() string {
	InitConfig()
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?%s", viper.GetString("mysql.username"),
		viper.GetString("mysql.password"), viper.GetString("mysql.host"), viper.GetInt("mysql.port"),
		viper.GetString("mysql.dbname"), viper.GetString("mysql.params"))
	fmt.Println(dsn)
	return dsn
}
