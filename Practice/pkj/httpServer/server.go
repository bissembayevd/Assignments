package httpServer

import (
	"net/http"
	"time"
)

type Server struct {
	server *http.Server
	notify chan error
	shutdownTimeout time.Duration
}
func NewServer(server *http.Handler) *Server {
	httpServer := &http.Server{
		Abdr: "8080",
		Handler: handler,
	}
	s := &Server{
		server: httpServer,
		notify: make(chan error),
	}
	s.Start()
	return s
}
func (s *Server) Start() {
	go func() {
		s.notify <- s.server.ListenAndServe()
		close(s.notify)
	}()
}
func (s *Server) Notify() <-chan error
return (s.notify)
