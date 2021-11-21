package config

import (
	"photographer/internal/consul"
)

// Config is general config for services
type Config struct {
	// consul addr
	CONSUL_ADDR string `default:"localhost:8500"`

	// services address
	S_MAIN_ADDR    string `default:"localhost:8080"` // main service address
	S_DB_ADDR      string `default:"localhost:8000"` // database service address
	S_REDIS_ADDR   string `default:"localhost:8010"` // redis service address
	S_API_ADDR     string `default:"localhost:8020"` // api service address
	S_AUTH_ADDR    string `default:"localhost:8030"` // auth service address
	S_BILLING_ADDR string `default:"localhost:8040"` // billing service address
}

// TODO: sync with consul: get other services address and so on
func (c *Config) SyncConfigWithConsul(client *consul.ConsulClient) error {
	// list, meta, _ := client.KV().List("", nil)

	// for _, v := range list {

	// }
	return nil
}

// NewConfig return new config
func NewConfig() *Config {
	return &Config{}
}
