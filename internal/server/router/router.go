package router

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/urusofam/calculatorRestAPI/internal/server/handlers"
	"github.com/urusofam/calculatorRestAPI/internal/storage"
)

func InitRouter(s *storage.Storage) (*gin.Engine, error) {
	router := gin.Default()

	err := router.SetTrustedProxies(nil)
	if err != nil {
		return nil, err
	}

	router.Use(cors.Default())
	router.GET("/calculations", handlers.GetCalculations(s))
	router.POST("/calculations", handlers.PostCalculation(s))
	router.PATCH("/calculations/:id", handlers.PatchCalculation(s))
	router.DELETE("/calculations/:id", handlers.DeleteCalculation(s))

	return router, nil
}
