package main

import (
	"github.com/IDarar/hub/internal/app"
)

const configPath = "configs/main"

// @title Hub
// @version 0.001
// @description Hub for specified topics

// @host localhost:8080
// @BasePath /api/v1/

// @securityDefinitions.apikey AdminAuth
// @in header
// @name Authorization

// @securityDefinitions.apikey StudentsAuth
// @in header
// @name Authorization
func main() {
	app.Run(configPath)
}
