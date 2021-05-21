package server

import (
	"log"
	"net/http"

	"github.com/IDarar/hub/internal/config"
)

type Server struct {
	httpServer *http.Server
}

func NewServer(cfg *config.Config, handler http.Handler) *Server {
	return &Server{
		httpServer: &http.Server{
			Addr:           ":" + cfg.HTTP.Port,
			Handler:        handler,
			ReadTimeout:    cfg.HTTP.ReadTimeout,
			WriteTimeout:   cfg.HTTP.WriteTimeout,
			MaxHeaderBytes: cfg.HTTP.MaxHeaderMegabytes << 20,
		},
	}
}
func (s *Server) Run() error {
	err := s.httpServer.ListenAndServe()
	if err != nil {
		log.Fatal("Could not start server ... ", err)
		return err
	}
	return nil
}
