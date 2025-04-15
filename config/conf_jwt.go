package config

type Jwt struct {
	Secret  string `mapstructure:"secrect" json:"secret" yaml:"secret"`   //密钥
	Expires int    `mapstructure:"exipres" json:"expires" yaml:"expires"` //过期时间
	Issuer  string `mapstructure:"issuer" json:"issuer" yaml:"issuer"`    //颁发人
}
