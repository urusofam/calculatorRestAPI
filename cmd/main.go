package main

import (
	"context"
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
	"github.com/urusofam/calculatorRestAPI/config"
	"github.com/urusofam/calculatorRestAPI/http/handlers"
	"log/slog"
	"os"
)

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	config, err := config.LoadConfig()
	if err != nil {
		logger.Error(err.Error())
	}
	logger.Info("config file loaded", slog.String("cfg", config.Server.Host))

	dbURL := fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s",
		config.Database.User,
		config.Database.Pass,
		config.Database.Host,
		config.Database.Port,
		config.Database.Name,
	)

	conn, err := pgx.Connect(context.Background(), dbURL)
	if err != nil {
		logger.Error(err.Error())
	}
	logger.Info("connected to database")
	defer conn.Close(context.Background())

	router := gin.Default()
	err = router.SetTrustedProxies(nil)
	if err != nil {
		logger.Error(err.Error())
	}

	router.Use(cors.Default())
	router.GET("/calculations", handlers.GetCalculations)
	router.POST("/calculations", handlers.PostCalculation)
	router.PATCH("/calculations/:id", handlers.PatchCalculation)
	router.DELETE("/calculations/:id", handlers.DeleteCalculation)

	if err = router.Run(fmt.Sprintf("%s:%d", config.Server.Host, config.Server.Port)); err != nil {
		logger.Error(err.Error())
	}
}
