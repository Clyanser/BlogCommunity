package config

type QiNiu struct {
	AccessKey string  `mapstructure:"access_key" json:"access_key" yaml:"access_key"`
	SecretKey string  `mapstructure:"secret_key" json:"secret_key" yaml:"secret_key"`
	Bucket    string  `mapstructure:"bucket" json:"bucket" yaml:"bucket"`
	CDN       string  `mapstructure:"cdn" json:"cdn" yaml:"cdn"`
	Zone      string  `mapstructure:"zone" json:"zone" yaml:"zone"`
	Size      float64 `mapstructure:"size" json:"size" yaml:"size"`
}
