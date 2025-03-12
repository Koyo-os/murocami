package agent

import (
	"github.com/koyo-os/murocami/internal/config"
	"github.com/koyo-os/murocami/pkg/logger"
)

type PipeLineRunner struct{
	cfg *config.Config
	logger *logger.Logger
}

func InitPipelineRunner(cfg *config.Config) *PipeLineRunner {
	return &PipeLineRunner{
		cfg: cfg,
		logger: logger.Init(),
	}
}