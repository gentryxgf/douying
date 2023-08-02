package config

import "fmt"

type RedisConf struct {
	Host     string `json:"host" mapstructure:"host"`
	Port     int    `json:"port" mapstructure:"port"`
	Password string `json:"password" mapstructure:"password"`
	PoolSize int    `json:"pool_size" mapstructure:"pool_size"`
	Db       int    `json:"db" mapstructure:"db"`
}

func (r RedisConf) Addr() string {
	return fmt.Sprintf("%s:%d", r.Host, r.Port)
}
