package service

import (
	"fmt"
	"time"

	"github.com/hashicorp/consul/api"
)

type ConsulService struct {
	Client *api.Client
}

func NewConsulService(address, port string) (ConsulService, error) {
	consulURL := fmt.Sprintf("http://%s:%s", address, port)
	const maxRetries = 5
	const retryInterval = 2 * time.Second

	consulConfig := api.DefaultConfig()
	consulConfig.Address = consulURL

	var consulClient *api.Client
	var err error

	for i := 0; i < maxRetries; i++ {
		consulClient, err = api.NewClient(consulConfig)
		if err == nil {
			return ConsulService{
				Client: consulClient,
			}, nil
		}
		fmt.Printf("Attempt %d/%d: Error connecting to Consul: %v\n", i+1, maxRetries, err)
		time.Sleep(retryInterval)
	}

	return ConsulService{}, fmt.Errorf("failed to create Consul client after %d attempts: %v", maxRetries, err)
}

func (cs ConsulService) GetServiceAdress(serviceName string) (string, error) {
	services, _, err := cs.Client.Health().Service(serviceName, "", true, nil)
	if err != nil {
		return "", fmt.Errorf("error obteniendo el servicio '%s' de Consul: %v", serviceName, err)
	}

	if len(services) == 0 {
		return "", fmt.Errorf("no se encontraron instancias del servicio '%s' en Consul", serviceName)
	}

	service := services[0]
	adress := fmt.Sprintf("http://%s:%d", service.Service.Address, service.Service.Port)

	return adress, nil
}
