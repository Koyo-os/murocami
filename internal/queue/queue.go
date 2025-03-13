package queue

import (
	"github.com/koyo-os/murocami/internal/config"
	"github.com/koyo-os/murocami/pkg/logger"
)

type QueueRunner struct{
	logger *logger.Logger
	cfg *config.Config
	queueConfig *config.QueueConfig
}

func Init(cfg *config.Config) *QueueRunner {
	logger := logger.Init()

	quecfg, err := config.InitQueueConfig()
	if err != nil{
		logger.Errorf("cant load queue config: %v",err)
		return nil
	}

	
}