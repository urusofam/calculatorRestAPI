package handlers

import (
	"fmt"
	"github.com/expr-lang/expr"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/urusofam/calculatorRestAPI/internal/server/models/calculation"
	"github.com/urusofam/calculatorRestAPI/internal/server/models/calculation_request"
	"github.com/urusofam/calculatorRestAPI/internal/storage"
	"net/http"
)

func calculateExpression(expression string) (string, error) {
	program, err := expr.Compile(expression)
	if err != nil {
		return "", err
	}

	result, err := expr.Run(program, nil)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%v", result), nil
}

func GetCalculations(strg *storage.Storage) gin.HandlerFunc {
	return func(c *gin.Context) {
		calculations, err := strg.GetCalculations()
		if err != nil {
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.IndentedJSON(http.StatusOK, calculations)
	}
}

func PostCalculation(strg *storage.Storage) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req calculation_request.CalculationsRequest
		if err := c.BindJSON(&req); err != nil {
			c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}

		result, err := calculateExpression(req.Expression)
		if err != nil {
			c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}

		calc := calculation.Calculation{
			ID:         uuid.NewString(),
			Expression: req.Expression,
			Result:     result,
		}

		err = strg.AddCalculation(calc)
		if err != nil {
			c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}

		c.IndentedJSON(http.StatusCreated, result)
	}
}

func PatchCalculation(s *storage.Storage) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")

		var req calculation_request.CalculationsRequest
		if err := c.BindJSON(&req); err != nil {
			c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}

		result, err := calculateExpression(req.Expression)
		if err != nil {
			c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}

		err = s.UpdateCalculation(req.Expression, result, id)
		if err != nil {
			c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
		c.IndentedJSON(http.StatusOK, id)
	}
}

func DeleteCalculation(s *storage.Storage) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")

		err := s.DeleteCalculation(id)
		if err != nil {
			c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
		c.Status(http.StatusNoContent)
	}
}
