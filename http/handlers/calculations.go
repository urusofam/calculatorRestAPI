package handlers

import (
	"fmt"
	"github.com/expr-lang/expr"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/urusofam/calculatorRestAPI/http/models/calcualtion"
	"github.com/urusofam/calculatorRestAPI/http/models/calculation_request"
	"net/http"
)

var calculations = []calcualtion.Calculation{}

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

func GetCalculations(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, calculations)
}

func PostCalculation(c *gin.Context) {
	var req calculation_request.CalculationsRequest
	if err := c.BindJSON(&req); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	result, err := calculateExpression(req.Expression)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	calc := calcualtion.Calculation{
		ID:         uuid.NewString(),
		Expression: req.Expression,
		Result:     result,
	}
	calculations = append(calculations, calc)

	c.IndentedJSON(http.StatusCreated, result)
}

func PatchCalculation(c *gin.Context) {
	id := c.Param("id")

	var req calculation_request.CalculationsRequest
	if err := c.BindJSON(&req); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	result, err := calculateExpression(req.Expression)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	for i, calc := range calculations {
		if calc.ID == id {
			calculations[i].Expression = req.Expression
			calculations[i].Result = result
			c.IndentedJSON(http.StatusOK, calculations[i])
			return
		}
	}

	c.IndentedJSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("calculation with id %s not found", id)})
}
