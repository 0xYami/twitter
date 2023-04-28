package main

import (
	"github.com/0xYami/twitter/config"
	"github.com/0xYami/twitter/server"
	"go.uber.org/zap"
)

func main() {
	logger, err := zap.NewProduction()
	if err != nil {
		panic("Failed to create logger")
	}
	defer logger.Sync()

	cfg, err := config.Load(".")
	if err != nil {
		logger.Panic("Failed to load config")
	}

	s := server.CreateNewServer(cfg, logger)

	s.InitDB()
	s.MountHandlers()

	err = s.Start()
	if err != nil {
		logger.Panic("Failed to start server")
	}
}
