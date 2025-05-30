package main

import (
	"fmt"
	"github.com/urusofam/calculatorRestAPI/config"
	"github.com/urusofam/calculatorRestAPI/internal/server/router"
	"github.com/urusofam/calculatorRestAPI/internal/storage"
	"log/slog"
	"os"
)

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	cfg, err := config.LoadConfig()
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}
	logger.Info("config file loaded", slog.String("cfg", cfg.Server.Host))

	dbURL := fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s?sslmode=disable",
		cfg.Database.User,
		cfg.Database.Pass,
		cfg.Database.Host,
		cfg.Database.Port,
		cfg.Database.Name,
	)

	strg, err := storage.NewStorage(dbURL)
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}
	logger.Info("connected to database", slog.String("url", dbURL))
	defer strg.Close()

	rout, err := router.InitRouter(strg)
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}

	if err = rout.Run(fmt.Sprintf("%s:%d", cfg.Server.Host, cfg.Server.Port)); err != nil {
		logger.Error(err.Error())
	}
}
