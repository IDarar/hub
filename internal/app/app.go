package app

import (
	"github.com/IDarar/hub/internal/config"
	"github.com/IDarar/hub/internal/repository"
	"github.com/IDarar/hub/internal/repository/postgres"
	"github.com/IDarar/hub/internal/server"
	"github.com/IDarar/hub/internal/service"
	"github.com/IDarar/hub/internal/transport/http"
	"github.com/IDarar/hub/pkg/auth"
	"github.com/IDarar/hub/pkg/logger"
)

func Run(configPath string) {
	cfg, err := config.Init(configPath)
	if err != nil {
		logger.Error(err)
		return
	}
	tokenManager, err := auth.NewManager("cfg.Auth.JWT.SigningKey")
	if err != nil {
		logger.Error(err)
		return
	}
	db, err := postgres.NewPostgresDB(cfg)
	if err != nil {
		logger.Error(err)
		return
	}
	repos := repository.NewRepositories(db)
	services := service.NewServices(service.Deps{
		Repos: repos,
	})
	handlers := http.NewHandler(services, tokenManager)
	srv := server.NewServer(cfg, handlers.Init())

	srv.Run()
}
