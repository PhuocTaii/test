package main

import (
	"ia-03-backend/application"
	"ia-03-backend/config/models"
	"ia-03-backend/docs"
	"ia-03-backend/pkg/config"
	"ia-03-backend/pkg/http/server"
	"os"
)

// @title                        Web Advanced IA-03 API
// @version                      v0.0.1
// @description                  API system for Web Advanced IA-03
// @termsOfService               https://swagger.io/

// @license.name                 Apache 2.0
// @license.url                  https://www.apache.org/licenses/LICENSE-2.0.html

// @host                         localhost:5000
// @BasePath                     /api

// @securitydefinitions.apikey   Jwt
// @in                           header
// @name                         Authorization

// @securitydefinitions.apikey   Accept-Language
// @in                           header
// @name                         Accept-Language

// @securitydefinitions.apikey   UID
// @in                           header
// @name                         uid
func main() {
	cfg, err := config.NewConfig(config.GetConfigPath(os.Getenv("ENVIRONMENT"), "json"), &models.Config{})
	if err != nil {
		panic(err)
	}

	docs.SwaggerInfo.Host = cfg.Service.Host
	server.Start(application.NewServer(cfg))
}
