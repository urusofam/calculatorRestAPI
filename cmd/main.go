package main

import (
	"github.com/urusofam/calculatorRestAPI/config"
)

func main() {
	config, err := config.LoadConfig()
	if err != nil {
		panic(err)
	}
	_ = config.Server
}
