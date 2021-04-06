package server

import (
	"log"
	"net/http"
)

//30.27 another
type Server struct {
	httpServer *http.Server
}

//TODO make config to server 29.51 another
func NewServer( /*cfg *config.Config*/ handler http.Handler) *Server {
	return &Server{
		httpServer: &http.Server{
			Addr:    ":8080",
			Handler: handler,
			//TODO add left fields with config
		},
	}
}
func (s *Server) Run() error {
	//TODO err
	err := s.httpServer.ListenAndServe()
	if err != nil {
		log.Fatal("Could not start server ... ", err)
		return err
	}
	return nil
}
