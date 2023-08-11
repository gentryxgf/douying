package config

type Config struct {
	SystemConf SystemConf `mapstructure:"system"`
	ZapConf    ZapConf    `mapstructure:"zap"`
	RedisConf  RedisConf  `mapstructure:"redis"`
	MysqlConf  MysqlConf  `mapstructure:"mysql"`
	JwtConf    JwtConf    `mapstructure:"jwt"`
	UploadConf UploadConf `mapstructure:"upload"`
}
