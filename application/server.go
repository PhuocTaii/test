package application

import (
	"fmt"
	dbsv "ia-03-backend/application/domains/db/init"
	userhandler "ia-03-backend/application/domains/user/handler"
	userusecase "ia-03-backend/application/domains/user/usecases"
	commonmodels "ia-03-backend/application/models"
	"ia-03-backend/application/routes"
	"ia-03-backend/config/models"
	"ia-03-backend/pkg/db/postgres"
	"ia-03-backend/pkg/http/server"
	"time"
)

func NewServer(cfg *models.Config) (*server.HTTPServer, error) {
	httpServer := server.NewHTTPServer(
		server.AddPort(cfg.Service.Port),
		server.AddName(cfg.Service.Name),
		server.SetStrictSlash(true),
		server.SetGracefulShutdownTimeout(time.Duration(cfg.Service.Timeout)),
	)

	// Create a new Postgres connection
	connection := postgres.Connection{
		Host:     cfg.Postgres.Host,
		Port:     cfg.Postgres.Port,
		User:     cfg.Postgres.User,
		Password: cfg.Postgres.Password,
		Database: cfg.Postgres.Database,
		// Connection pool settings
		MaxOpenConnections:    cfg.Postgres.MaxOpenConnections,
		MaxIdleConnections:    cfg.Postgres.MaxIdleConnections,
		ConnectionMaxIdleTime: time.Duration(cfg.Postgres.ConnectionMaxIdleTime) * time.Second,
		ConnectionMaxLifeTime: time.Duration(cfg.Postgres.ConnectionMaxLifeTime) * time.Second,
		// SSL settings
		SSLMode:     postgres.SSLMode(cfg.Postgres.SSLMode),
		SSLCert:     cfg.Postgres.SSLCertPath,
		SSLKey:      cfg.Postgres.SSLKeyPath,
		SSLRootCert: cfg.Postgres.SSLRootCertPath,
	}
	db, err := postgres.NewDatabase(connection, cfg.Postgres.LogLevel, cfg.Postgres.SlowQueryThreshold)
	if err != nil {
		fmt.Printf("failed to connect database: %v", err)
		panic(err)
	}

	// Initialize the lib
	lib := &commonmodels.Lib{
		Db:  db.Executor,
		Cfg: cfg,
	}
	dbSv := dbsv.NewInit(lib)

	httpServer.AddGroupRoutes(routes.InitHandler(lib, dbSv, userhandler.NewHandler(lib, userusecase.InitUseCase(lib, dbSv))).NewGroupRoutes())
	return httpServer, nil
}
