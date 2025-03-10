package main

import (
	"net/http"

	"github.com/koyo-os/murocami/internal/config"
	"github.com/koyo-os/murocami/internal/handler"
	"github.com/koyo-os/murocami/internal/server"
	"github.com/koyo-os/murocami/pkg/logger"
)

func main() {
	logger := logger.Init()

	cfg, err := config.Init()
	if err != nil{
		logger.Errorf("cant get config: %v", err)
		return
	}

	s := server.Init(cfg)

	h := handler.InitHandler(cfg)
	mux := http.NewServeMux()
	
	h.Routes(mux)
	s.SetHandler(mux)

	err = s.Start()
	if err != nil{
		logger.Errorf("cant run murocami: %v",err)
	}
}