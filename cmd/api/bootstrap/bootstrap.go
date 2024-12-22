package bootstrap

import (
	"ch-gateway/internal/shared/platform/server"
	"context"
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/kelseyhightower/envconfig"
)

func Run() error {
	var conf config
	err := envconfig.Process("APP", &conf)
	if err != nil {
		return err
	}

	// Change according to the database to be used
	dbURI := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", conf.DbUsername, conf.DbPassword, conf.DbHost, conf.DbPort, conf.DbName)

	db, err := sql.Open("mysql", dbURI)
	if err != nil {
		return err
	}
	if err := db.Ping(); err != nil {
		return fmt.Errorf("failed to connect to the database: %w", err)
	}

	ctx, srv := server.NewServer(context.Background(), conf.Host, conf.Port, conf.ShutdownTimeout)
	return srv.Run(ctx)
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
}
