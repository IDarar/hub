package main

import (
	"github.com/IDarar/hub/internal/app"
)

const configPath = "configs/main"

func main() {
	app.Run(configPath)
}
