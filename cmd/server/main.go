package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"

	"github.com/koyo-os/murocami/internal/config"
	"github.com/koyo-os/murocami/internal/handler"
	"github.com/koyo-os/murocami/internal/server"
	"github.com/koyo-os/murocami/pkg/logger"
)

const ART = `


███╗   ███╗██╗   ██╗██████╗  ██████╗  ██████╗ █████╗ ███╗   ███╗██╗
████╗ ████║██║   ██║██╔══██╗██╔═══██╗██╔════╝██╔══██╗████╗ ████║██║
██╔████╔██║██║   ██║██████╔╝██║   ██║██║     ███████║██╔████╔██║██║
██║╚██╔╝██║██║   ██║██╔══██╗██║   ██║██║     ██╔══██║██║╚██╔╝██║██║
██║ ╚═╝ ██║╚██████╔╝██║  ██║╚██████╔╝╚██████╗██║  ██║██║ ╚═╝ ██║██║
╚═╝     ╚═╝ ╚═════╝ ╚═╝  ╚═╝ ╚═════╝  ╚═════╝╚═╝  ╚═╝╚═╝     ╚═╝╚═╝
                                                                   


`

func main() {
	fmt.Print(ART)

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	logger := logger.Init()

	logger.Info("setup config...")

	cfg, err := config.Init()
	if err != nil {
		logger.Errorf("cant get config: %v", err)
		return
	}

	logger.Info("setup directory...")

	os.Remove(cfg.TempDirName)

	s := server.Init(cfg)

	h := handler.InitHandler(cfg)
	mux := http.NewServeMux()

	h.Routes(mux)
	s.SetHandler(mux)

	go func() {
		<-ctx.Done()
		logger.Info("murocami stopped!")
		s.Stop(ctx)

		os.Remove(cfg.TempDirName)
	}()

	err = s.Start()
	if err != nil && err != http.ErrServerClosed {
		logger.Errorf("cant run murocami: %v", err)
	}
}

