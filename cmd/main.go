package main

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
	"github.com/urusofam/calculatorRestAPI/config"
	"github.com/urusofam/calculatorRestAPI/internal/server"
	"log/slog"
	"os"
)

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	config, err := config.LoadConfig()
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}
	logger.Info("config file loaded", slog.String("cfg", config.Server.Host))

	dbURL := fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s?sslmode=disable",
		config.Database.User,
		config.Database.Pass,
		config.Database.Host,
		config.Database.Port,
		config.Database.Name,
	)

	conn, err := pgx.Connect(context.Background(), dbURL)
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}
	logger.Info("connected to database")
	defer conn.Close(context.Background())

	router, err := server.InitServer()
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}

	if err = router.Run(fmt.Sprintf("%s:%d", config.Server.Host, config.Server.Port)); err != nil {
		logger.Error(err.Error())
	}
}
