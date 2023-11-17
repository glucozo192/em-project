package configs

import "fmt"

type Endpoint struct {
	Host  string `mapstructure:"host" json:"host"`
	Port  int    `mapstructure:"port" json:"port"`
	Token string `mapstructure:"token" json:"token"`
}

func (e *Endpoint) Address() string {
	return fmt.Sprintf(`%s:%d`, e.Host, e.Port)
}
