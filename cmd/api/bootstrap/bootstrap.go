package bootstrap

import (
	dependencycontainer "ch-gateway/internal/shared/dependencyContainer"
	"ch-gateway/internal/shared/platform/server"
	"ch-gateway/internal/shared/service"
	"context"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/kelseyhightower/envconfig"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Run() error {
	var conf config
	err := envconfig.Process("APP", &conf)
	if err != nil {
		return err
	}

	//DataBase setup
	dbURI := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", conf.DbUsername, conf.DbPassword, conf.DbHost, conf.DbPort, conf.DbName)

	db, err := waitForDB(dbURI, conf.DbTimeOut)
	if err != nil {
		return err
	}
	sqlDB, _ := db.DB()
	if err := sqlDB.Ping(); err != nil {
		return fmt.Errorf("failed to connect to the database: %w", err)
	}
	//ConsulServer sertup
	consulService, _ := service.NewConsulService(conf.ConsulAdress, conf.ConsulPort)

	//Dependecy container setup
	container := dependencycontainer.NewContainer(db, conf.SecretKey)
	container.Services.DiscoveryService = consulService

	ctx, srv := server.NewServer(context.Background(), conf.Host, conf.Port, conf.ShutdownTimeout)
	return srv.Run(ctx, container)
}

func waitForDB(dbURI string, timeout time.Duration) (*gorm.DB, error) {
	start := time.Now()
	for {
		db, err := gorm.Open(mysql.Open(dbURI))
		if err == nil {
			return db, nil
		}
		if time.Since(start) > timeout {
			return nil, fmt.Errorf("database not ready: %w", err)
		}
		fmt.Println("Waiting for database to be ready...")
		time.Sleep(2 * time.Second)
	}
}

type config struct {
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
	SecretKey    string `envconfig:"SECRET_KEY" default:"secretos123"`
}
