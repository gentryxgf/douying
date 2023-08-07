package config

import "fmt"

type SystemConf struct {
	Host string `json:"host" mapstructure:"host"`
	Port int    `json:"port" mapstructure:"port"`
	Env  string `json:"env" mapstructure:"env"`
}

func (s SystemConf) Addr() string {
	return fmt.Sprintf("%s:%d", s.Host, s.Port)
}
