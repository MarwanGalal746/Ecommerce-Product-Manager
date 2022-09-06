package main

import (
	"Ecommerce-Product-Manager/pkg/config"
	"fmt"
	"os"
)

func init() {
	err := config.Config()
	if err != nil {
		fmt.Println("application has stopped due to configuration issues")
		os.Exit(1)
	}
}

func main() {
	config.Logger.Info("Application has started")
	fmt.Println("hello")
}
