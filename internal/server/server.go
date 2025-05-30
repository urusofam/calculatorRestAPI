package server

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/urusofam/calculatorRestAPI/internal/server/handlers"
)

func NewServer() (*gin.Engine, error) {
	router := gin.Default()

	err := router.SetTrustedProxies(nil)
	if err != nil {
		return nil, err
	}

	router.Use(cors.Default())
	router.GET("/calculations", handlers.GetCalculations)
	router.POST("/calculations", handlers.PostCalculation)
	router.PATCH("/calculations/:id", handlers.PatchCalculation)
	router.DELETE("/calculations/:id", handlers.DeleteCalculation)

	return router, nil
}
