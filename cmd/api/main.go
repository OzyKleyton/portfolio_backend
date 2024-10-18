package main

import (
	"github.com/ozykt4/portfolio_backend/config"
	"github.com/ozykt4/portfolio_backend/internal/api"
	"github.com/ozykt4/portfolio_backend/internal/database"
)

func main() {
	config.LoadConfig()
	port := config.GetConfig().Port

	host := "0.0.0.0"
	if config.GetConfig().Prefork {
		host = "0.0.0.0"
	}

	database.InitMigrate()

	if err := (api.Run(host, port)); err != nil {
		panic(err)
	}

}
