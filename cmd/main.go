package main

import (
	"github.com/urusofam/calculatorRestAPI/config"
	"log/slog"
	"os"
)

func main() {
	config, err := config.LoadConfig()
	if err != nil {
		panic(err)
	}

	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	logger.Info("config file loaded", slog.String("cfg", config.Server.Host))
}
