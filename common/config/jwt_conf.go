package config

type JwtConf struct {
	Secret  string `json:"secret" mapstructure:"secret"`
	Expires int    `json:"expires" mapstructure:"expires"`
	Issuer  string `json:"issuer" mapstructure:"issuer"`
}
