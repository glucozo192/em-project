package configs

import "fmt"

type Database struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
	Database string `mapstructure:"database"`
}

func (e *Database) Address() string {
	return fmt.Sprintf(`%s:%d`, e.Host, e.Port)
}
