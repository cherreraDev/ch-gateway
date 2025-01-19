package bootstrap

import (
	"ch-gateway/cmd/api/bootstrap/config"
	dependencycontainer "ch-gateway/internal/shared/dependencyContainer"
	"ch-gateway/internal/shared/platform/server"
	"context"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Run() error {
	err := config.LoadConfig()
	if err != nil {
		return err
	}

	//DataBase setup
	DBuserName, DBpassword, DBhost, DBtimeout, DBport, DBname := config.GlobalConfig.GetDBconfigurations()
	dbURI := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", DBuserName, DBpassword, DBhost, DBport, DBname)

	db, err := waitForDB(dbURI, DBtimeout)
	if err != nil {
		return err
	}

	//Dependecy container setup
	container, err := dependencycontainer.NewContainer(db)
	if err != nil {
		return err
	}
	host, port, shutdownTimeout := config.GlobalConfig.GetServerconfigurations()
	ctx, srv := server.NewServer(context.Background(), host, port, shutdownTimeout)
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
