package config

import (
	"time"

	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	//Server configuration
	Host            string        `envconfig:"HOST" default:"127.0.0.1"`
	Port            string        `envconfig:"PORT" default:"8080"`
	ShutdownTimeout time.Duration `envconfig:"SHUTDOWN_TIMEOUT" default:"10s"`
	//Database configuration
	DbUsername string        `envconfig:"DB_USERNAME" default:"user"`
	DbPassword string        `envconfig:"DB_PASSWORD" default:"userpassword"`
	DbHost     string        `envconfig:"DB_HOST" default:"localhost"`
	DbPort     uint          `envconfig:"DB_PORT" default:"3306"`
	DbName     string        `envconfig:"DB_NAME" default:"mydb"`
	DbTimeOut  time.Duration `envconfig:"DB_TIMEOUT" default:"10s"`
	//Consul configuration
	ConsulAdress string `envconfig:"CONSUL_ADRESS" default:"consul"`
	ConsulPort   string `envconfig:"CONSUL_PORT" default:"8500"`
	//Other
	SecretKey   string `envconfig:"SECRET_KEY" default:"secretos123"`
	Environment string `envconfig:"ENVIRONMENT" default:"dev"`
}

var GlobalConfig *Config

func LoadConfig() error {
	GlobalConfig = &Config{}
	err := envconfig.Process("", GlobalConfig)
	if err != nil {
		return err
	}
	return nil
}

func (cf *Config) GetDBconfigurations() (
	DBusername, DBpassword, DBhost string, DBtimeout time.Duration,
	DBport uint, DBname string) {
	return cf.DbUsername, cf.DbPassword, cf.DbHost, cf.DbTimeOut, cf.DbPort, cf.DbName
}
func (cf *Config) GetServerconfigurations() (host, port string, shutdownTimeout time.Duration) {
	return cf.Host, cf.Port, cf.ShutdownTimeout
}
