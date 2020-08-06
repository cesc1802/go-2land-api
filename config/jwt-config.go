package config

import "os"

type JwtConfig struct {
	JwtSecret string
	JwtRefreshSecret string
	JwtExpired string
	JwtRefreshExpired string
}

var JwtConf JwtConfig

func Read() {
	JwtConf.JwtExpired = os.Getenv("JWT_SECRET")
	JwtConf.JwtRefreshExpired = os.Getenv("JWT_REFRESH_SECRET")
	JwtConf.JwtExpired=os.Getenv("JWT_EXPIRED")
	JwtConf.JwtRefreshExpired=os.Getenv("JWT_REFRESH_TOKEN")
}