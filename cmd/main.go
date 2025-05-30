package main

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
	"github.com/urusofam/calculatorRestAPI/config"
	"github.com/urusofam/calculatorRestAPI/internal/server/router"
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

	conn, err := pgx.Connect(context.Background(), dbURL)
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}
	logger.Info("connected to database")
	defer func(conn *pgx.Conn, ctx context.Context) {
		err := conn.Close(ctx)
		if err != nil {
			logger.Error(err.Error())
		}
	}(conn, context.Background())

	rout, err := router.InitRouter()
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}

	if err = rout.Run(fmt.Sprintf("%s:%d", cfg.Server.Host, cfg.Server.Port)); err != nil {
		logger.Error(err.Error())
	}
}
