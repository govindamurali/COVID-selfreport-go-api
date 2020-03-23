package main

import (
	"fmt"
	"github.com/govindamurali/COVID-selfreport-go-api/app"
	"github.com/govindamurali/COVID-selfreport-go-api/config"
)

func main() {
	fmt.Println("Starting the application...")
	config := config.GetConfig()

	app := &app.App{}
	app.Initialize(config)
	app.Run(config.Port)
}
