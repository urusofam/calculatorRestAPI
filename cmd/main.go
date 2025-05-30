package main

import (
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
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

	router := gin.Default()
	err = router.SetTrustedProxies(nil)
	if err != nil {
		logger.Error(err.Error())
	}

	router.Use(cors.Default())
	router.GET("/calculations", handlers.GetCalculations)
	router.POST("/calculations", handlers.PostCalculation)

	if err = router.Run(fmt.Sprintf("%s:%d", config.Server.Host, config.Server.Port)); err != nil {
		logger.Error(err.Error())
	}
}
