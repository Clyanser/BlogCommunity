package config

type Email struct {
	Host             string `json:"host" yaml:"host"`
	Port             int    `json:"port" yaml:"port"`
	User             string `json:"user" yaml:"user"` //发送人邮箱
	Password         string `json:"password" yaml:"password"`
	DefaultFromEmail string `json:"defaultFromEmail" yaml:"defaultFromEmail"` //默认的发件人名字
	UseSSL           bool   `json:"use_ssL" yaml:"use_ssL"`                   //是否使用ssl
	UseTls           bool   `json:"use_tls" yaml:"use_tls"`                   //是否使用Tls
}
