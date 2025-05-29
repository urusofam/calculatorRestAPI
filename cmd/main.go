package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/urusofam/calculatorRestAPI/config"
	"log/slog"
	"net/http"
	"os"
)

func main() {
	config, err := config.LoadConfig()
	if err != nil {
		panic(err)
	}

	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	logger.Info("config file loaded", slog.String("cfg", config.Server.Host))

	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Привет, Gin!")
	})

	router.Run(fmt.Sprintf("%s:%d", config.Server.Host, config.Server.Port))
}
