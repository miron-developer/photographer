package consul

import (
	"fmt"

	"github.com/hashicorp/consul/api"
)

// ConsulClient type
type ConsulClient struct {
	*api.Client

	UnregisterService func() error
}

// NewClient return new consul api client
func NewClient(consulAgentAddr string) (*ConsulClient, error) {
	consulConf := api.DefaultConfig()
	consulConf.Address = consulAgentAddr

	client, e := api.NewClient(consulConf)
	if e != nil {
		return nil, fmt.Errorf("consul client create error: %v", e.Error())
	}

	return &ConsulClient{Client: client}, nil
}
