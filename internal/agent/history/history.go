package history

import (
	"io"
	"os"
	"time"

	"github.com/bytedance/sonic"
	"github.com/koyo-os/murocami/internal/config"
	"github.com/koyo-os/murocami/internal/models"
	"github.com/koyo-os/murocami/pkg/logger"
)

type AgentHistory struct{
	cfg *config.Config
	logger *logger.Logger
	file *os.File
}

func Init(cfg *config.Config) *AgentHistory {
	logger := logger.Init()

	file, err := os.Open(cfg.FileHistory)
	if err != nil{
		logger.Errorf("cant open %s: %v", cfg.FileHistory, err)
		return nil
	}

	return &AgentHistory{
		cfg: cfg,
		file: file,
		logger: logger,
	}
}

func (a *AgentHistory) Save(ok bool, message string) error {
	body, err := io.ReadAll(a.file)
	if err != nil{
		a.logger.Errorf("cant get body: %v",body)
		return err
	}

	if string(body) == "" {
		block := models.Block{
			Message: message,
			Ok: ok,
			TimeStamp: time.Now().Format("2006-01-02 15:04:05"),
		}

		newbody,err := sonic.Marshal(&block)
		if err != nil{
			a.logger.Errorf("cant marshal new body: %v", err)
			return err
		}

		a.file.Write(newbody)
	}

	var blocks []models.Block

	if err = sonic.Unmarshal(body, &blocks);err != nil{
		a.logger.Errorf("error unmarshal body: %v",err)
		return err
	}

	blocks = append(blocks, models.Block{
		Ok: ok,
		Message: message,
		TimeStamp: time.Now().Format("2006-01-02 15:04:05"),
	})

	newBody, err := sonic.Marshal(&blocks)
	if err != nil{
		a.logger.Errorf("cant get new body: %v",err)
		return err
	}

	_, err = a.file.Write(newBody)
	if err != nil{
		a.logger.Errorf("cant write new body %v",err)
		return err
	}

	return nil
}