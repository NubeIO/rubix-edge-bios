package config

type ServerConfiguration struct {
	Secret                     string
	AccessTokenExpireDuration  int
	RefreshTokenExpireDuration int
	LimitCountPerRequest       float64
	LogLevel                   string
}
