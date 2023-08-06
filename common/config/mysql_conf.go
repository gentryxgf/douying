package config

import "fmt"

type MysqlConf struct {
	Host         string `json:"host" mapstructure:"host"`
	Port         int    `json:"port" mapstructure:"port"`
	Config       string `json:"config" mapstructure:"config"`
	User         string `json:"user" mapstructure:"user"`
	Password     string `json:"password" mapstructure:"password"`
	DbName       string `json:"db_name" mapstructure:"db_name"`
	LogLevel     string `json:"log_level" mapstructure:"log_level"`
	MaxIdelConns int    `json:"max_idel_conns" mapstructure:"max_idel_conns"`
	MaxOpenConns int    `json:"max_open_conns" mapstructure:"max_open_conns"`
}

func (m MysqlConf) Dsn() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?%s", m.User, m.Password, m.Host, m.Port, m.DbName, m.Config)
}
