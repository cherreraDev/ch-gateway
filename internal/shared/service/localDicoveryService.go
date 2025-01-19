package service

import (
	"errors"
	"fmt"
)

type LocalDiscovery struct {
	ServicesAdress map[string]string
}

func NewLocalDiscovery() LocalDiscovery {
	return LocalDiscovery{
		ServicesAdress: map[string]string{
			"tasks-service":   "localhost:8081",
			"company-service": "localhost:8082",
		},
	}
}

func (ld LocalDiscovery) GetServiceAdress(serviceName string) (string, error) {
	adress, exists := ld.ServicesAdress[serviceName]
	if exists {
		return "", errors.New("error discovering service")
	}

	return fmt.Sprintf("http://%s", adress), nil
}
