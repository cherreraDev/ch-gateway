package dependencycontainer

import (
	"ch-gateway/cmd/api/bootstrap/config"
	"ch-gateway/internal/shared/domain/discovery"
	"ch-gateway/internal/shared/service"
	"ch-gateway/internal/user/domain"
	"ch-gateway/internal/user/platform/storage/repositories"
	crudservice "ch-gateway/internal/user/service/crudService"
	loginservices "ch-gateway/internal/user/service/loginServices"

	"gorm.io/gorm"
)

type Repositories struct {
	UserRepository domain.UserRepository
}

type Services struct {
	LoginService     domain.LoginService
	UserService      domain.UserService
	DiscoveryService discovery.DiscoveryServer
}

type Container struct {
	Repositories Repositories
	Services     Services
	SigningKey   string
}

func NewRepositories(db *gorm.DB) Repositories {
	return Repositories{
		UserRepository: repositories.NewGormUserRepository(db),
	}
}

func NewServices(repos Repositories) (Services, error) {
	dicoveryServer, err := loadDiscoveryServer()
	if err != nil {
		return Services{}, err
	}
	return Services{
		LoginService:     loginservices.NewUserPasswordLoginService(repos.UserRepository, config.GlobalConfig.SecretKey),
		UserService:      crudservice.NewUserService(repos.UserRepository),
		DiscoveryService: dicoveryServer,
	}, nil
}

func NewContainer(db *gorm.DB) (Container, error) {
	repos := NewRepositories(db)
	services, err := NewServices(repos)
	if err != nil {
		return Container{}, err
	}

	return Container{
		Repositories: repos,
		Services:     services,
		SigningKey:   config.GlobalConfig.SecretKey,
	}, nil
}

func loadDiscoveryServer() (discovery.DiscoveryServer, error) {
	var discoveryServer discovery.DiscoveryServer
	var err error = nil
	switch config.GlobalConfig.Environment {
	case "dev":
		discoveryServer = service.NewLocalDiscovery()
	case "pre":
		//ConsulServer sertup
		discoveryServer, err = service.NewConsulService(config.GlobalConfig.ConsulAdress, config.GlobalConfig.ConsulPort)
	default:
		discoveryServer = service.NewLocalDiscovery()
	}
	return discoveryServer, err

}
