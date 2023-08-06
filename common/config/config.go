package config

type Config struct {
	ZapConf   ZapConf   `mapstructure:"zap"`
	RedisConf RedisConf `mapstructure:"redis"`
	MysqlConf MysqlConf `mapstructure:"mysql"`
	JwtConf   JwtConf   `mapstructure:"jwt"`
}
