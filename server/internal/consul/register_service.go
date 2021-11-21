package consul

import "github.com/hashicorp/consul/api"

func (client *ConsulClient) RegisterService() error {
	client.Agent().ServiceRegister(&api.AgentServiceRegistration{})

	return nil
}
