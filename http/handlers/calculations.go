package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/urusofam/calculatorRestAPI/http/models/calcualtion"
	"net/http"
)

var calculations = []calcualtion.Calculation{}

func GetCalculations(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, calculations)
}
