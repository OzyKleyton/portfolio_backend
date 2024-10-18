package api

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/gofiber/fiber/v2"
	"github.com/ozykt4/portfolio_backend/config"
	"github.com/ozykt4/portfolio_backend/config/db"
	"github.com/ozykt4/portfolio_backend/internal/api/handler"
	"github.com/ozykt4/portfolio_backend/internal/api/router"
	"github.com/ozykt4/portfolio_backend/internal/repository"
	"github.com/ozykt4/portfolio_backend/internal/service"
)

func Run(host, port string) error {
	address := fmt.Sprintf("%s:%s", host, port)
	log.Println("Listen app in port ", address)

	app := fiber.New(fiber.Config{
		JSONEncoder: json.Marshal,
		JSONDecoder: json.Unmarshal,
		Prefork:     config.GetConfig().Prefork,
	})

	db, err := db.ConnectDB(config.GetConfig().DBURL)
	if err != nil {
		return err
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	db = db.WithContext(ctx)

	// Loads all repositories

	projectRepo := repository.NewProjectRepository(db)

	// Loads all services
	projectService := service.NewProjectService(projectRepo)

	// Loads all handlers

	projectHandlers := handler.NewProjectHandler(projectService)

	// Setup middlewares

	router.SetupRouter(app,
		projectHandlers.Routes(),
	)

	c := make(chan os.Signal, 1)
	errc := make(chan error, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		fmt.Println("Gracefully shutting down...")
		cancel()
		errc <- app.Shutdown()
	}()

	if err := app.Listen(address); err != nil {
		return err
	}

	err = <-errc

	return err
}
