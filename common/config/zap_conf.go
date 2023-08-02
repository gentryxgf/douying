package config

type ZapConf struct {
	Format       string `json:"format" mapstructure:"format"`
	Mode         string `json:"mode" mapstructure:"mode"`
	WarnFilename string `json:"warn_filename" mapstructure:"warn_filename"`
	InfoFilename string `json:"info_filename" mapstructure:"info_filename"`
	MaxSize      int    `json:"max_size" mapstructure:"max_size"`
	MaxAge       int    `json:"max_age" mapstructure:"max_age"`
	MaxBackups   int    `json:"max_backups" mapstructure:"max_backups"`
	LocalTime    bool   `json:"local_time" mapstructure:"local_time"`
}
