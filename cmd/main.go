package main

import (
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/urusofam/calculatorRestAPI/config"
	"github.com/urusofam/calculatorRestAPI/internal/database"
	"github.com/urusofam/calculatorRestAPI/internal/server/handlers"
	"github.com/urusofam/calculatorRestAPI/internal/server/repositories"
	"github.com/urusofam/calculatorRestAPI/internal/server/services"
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
	logger.Info("config file loaded")

	db, err := database.InitDB(cfg.Database.User, cfg.Database.Pass,
		cfg.Database.Host, cfg.Database.Port, cfg.Database.Name)
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}
	logger.Info("connected to database")
	defer db.Close()

	router := gin.Default()

	err = router.SetTrustedProxies(nil)
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}

	calcRepo := repositories.NewCalculationRepository(db)
	calcService := services.NewCalculationService(calcRepo)
	calcHandler := handlers.NewCalculationHandler(calcService)

	router.Use(cors.Default())
	router.GET("/calculations", calcHandler.GetCalculations)
	router.POST("/calculations", calcHandler.PostCalculation)
	router.PATCH("/calculations/:id", calcHandler.PatchCalculation)
	router.DELETE("/calculations/:id", calcHandler.DeleteCalculation)

	if err = router.Run(fmt.Sprintf("%s:%d", cfg.Server.Host, cfg.Server.Port)); err != nil {
		logger.Error(err.Error())
	}
}
