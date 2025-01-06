package discovery

type DiscoveryServer interface {
	GetServiceAdress(serviceName string) (string, error)
}
