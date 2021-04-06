package app

import (
	"hub/internal/server"
	"hub/internal/service"
	"hub/internal/transport/http"
)

//53.43 31.45(another)
func Run(configPath string) {
	//TODO services := service.NewServices()
	handlers := http.NewHandler(service.Users{}, service.Admins{})
	srv := server.NewServer(handlers.Init())

	srv.Run()
}
