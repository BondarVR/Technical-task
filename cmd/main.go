package main

import (
	"log"
	technical_task "technical-task"
	"technical-task/pkg/config"
	"technical-task/pkg/handler"
	"technical-task/pkg/logger"
	"technical-task/pkg/repository"
	"technical-task/pkg/service"
)

func main() {
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatal(err)
	}

	lgr, err := logger.New(cfg)
	if err != nil {
		lgr.Fatalf("failed to initialize logger: %v", err.Error())
	}

	db, err := repository.NewClient(
		cfg.MongoUser,
		cfg.MongoPassword,
		cfg.MongoHost,
		cfg.MongoPort,
		cfg.NameDatabase)
	if err != nil {
		lgr.Fatalf("failed to initialize db: %v", err.Error())
	}

	repo := repository.NewRepository(db)
	services := service.NewService(repo)
	handlers := handler.NewHandler(services)

	server := new(technical_task.Server)
	lgr.Infof("Server run on http://%s:%s", cfg.ServerHost, cfg.ServerPort)
	err = server.Run(cfg.ServerHost, cfg.ServerPort, handlers.InitRoutes())
	if err != nil {
		log.Fatal(err)
	}
}
