package config

type DatabaseConfiguration struct {
	Driver   string
	Name     string
	Username string
	Password string
	Host     string
	Port     string
	LogLevel string
}
