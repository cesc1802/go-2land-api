package config

import "os"

type databaseConfig struct {
	DbHost string
	DbPort string
}

var DatabaseConf databaseConfig

func ReadDbConf() {
	DatabaseConf.DbHost = os.Getenv("DB_HOST")
	DatabaseConf.DbPort = os.Getenv("DB_PORT")
}
