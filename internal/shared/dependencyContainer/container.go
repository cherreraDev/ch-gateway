package dependencycontainer

import (
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
	LoginService domain.LoginService
	UserService  domain.UserService
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

func NewServices(repos Repositories, signingKey string) Services {
	return Services{
		LoginService: loginservices.NewUserPasswordLoginService(repos.UserRepository, signingKey),
		UserService:  crudservice.NewUserService(repos.UserRepository),
	}
}

func NewContainer(db *gorm.DB, signingKey string) Container {
	repos := NewRepositories(db)
	services := NewServices(repos, signingKey)

	return Container{
		Repositories: repos,
		Services:     services,
		signingKey:   signingKey,
	}
}
