package internal

import (
	"photographer/internal/config"

	"github.com/kelseyhightower/envconfig"
)

// Config extends general config with db service config
type Config struct {
	*config.Config

	DB_HOST     string `default:"127.0.0.1"`
	DB_PORT     string `default:"5432"`
	DB_NAME     string `default:"photographer"`
	DB_USER     string `default:"postgres"`
	DB_PASSWORD string `default:"postgres"`
}

// NewConfig configure db service config
func (service *DB_SERVICE) NewConfig() {
	service.Config = &Config{}

	// get configs from os
	service.Log.Infoln("getting config from environments")
	if e := envconfig.Process("", service.Config); e != nil {
		service.Log.Fatalln(e)
	}
	service.Log.Infoln("done!")
}

// SyncWithConsul get configs from consul
func (service *DB_SERVICE) SyncWithConsul() {
	service.Log.Infoln("getting config from consul")
	if e := service.Config.Config.SyncConfigWithConsul(service.Client); e != nil {
		service.Log.Fatalln("sync from consul error: ", e)
	}
	service.Log.Infoln("done!")
}
