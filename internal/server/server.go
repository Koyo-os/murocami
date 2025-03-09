package server

import (
	"context"
	"fmt"
	"net/http"

	"github.com/koyo-os/murocami/internal/config"
	"github.com/koyo-os/murocami/pkg/logger"
)

type Server struct{
	logger *logger.Logger
	*http.Server
}

func Init(cfg *config.Config) *Server {
	return &Server{
		logger: logger.Init(),
		Server: &http.Server{
			Addr: fmt.Sprintf("%s:%s", cfg.Host, cfg.Port),
		},
	}
}

func (s *Server) Start() error {
	s.logger.Info("starting server...")

	return s.ListenAndServe()
}

func (s *Server) Stop(ctx context.Context) {
	s.logger.Info("server stopping...")

	s.Shutdown(ctx)
}