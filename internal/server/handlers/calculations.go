package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/urusofam/calculatorRestAPI/internal/server/models"
	"github.com/urusofam/calculatorRestAPI/internal/server/services"
	"net/http"
)

type CalculationHandler struct {
	service services.CalculationService
}

func NewCalculationHandler(s services.CalculationService) *CalculationHandler {
	return &CalculationHandler{service: s}
}

func (h *CalculationHandler) GetCalculations(c *gin.Context) {
	calculations, err := h.service.GetCalculations()

	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	c.IndentedJSON(http.StatusOK, calculations)
}

func (h *CalculationHandler) PostCalculation(c *gin.Context) {
	var req models.CalculationsRequest

	if err := c.BindJSON(&req); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	result, err := h.service.AddCalculation(req.Expression)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	c.IndentedJSON(http.StatusCreated, result)
}

func (h *CalculationHandler) PatchCalculation(c *gin.Context) {
	id := c.Param("id")

	var req models.CalculationsRequest
	if err := c.BindJSON(&req); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	result, err := h.service.UpdateCalculation(req.Expression, id)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	c.IndentedJSON(http.StatusOK, result)
}

func (h *CalculationHandler) DeleteCalculation(c *gin.Context) {
	id := c.Param("id")

	err := h.service.DeleteCalculation(id)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	c.Status(http.StatusNoContent)
}
