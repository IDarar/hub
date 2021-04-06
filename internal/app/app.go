package app

import (
	"fmt"

	"github.com/IDarar/hub/internal/config"
	"github.com/IDarar/hub/internal/server"
	"github.com/IDarar/hub/internal/service"
	"github.com/IDarar/hub/internal/transport/http"
)

func Run(configPath string) {
	cfg, err := config.Init(configPath)
	if err != nil {
		fmt.Println("Could not init cfg ... ", err)
		return
	}

	handlers := http.NewHandler(service.Users{}, service.Admins{})
	srv := server.NewServer(cfg, handlers.Init())

	srv.Run()
}
