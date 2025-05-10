package config

type Redis struct {
	Ip       string `json:"ip" yaml:"ip"`
	Port     string `json:"port" yaml:"port"`
	Password string `json:"password" yaml:"password"`
	PoolSize int    `json:"pool_size" yaml:"poolSize"`
}

func (r Redis) Addr() string {
	return r.Ip + ":" + r.Port
}
