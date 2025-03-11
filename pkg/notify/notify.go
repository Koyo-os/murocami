package notify

import (
	"github.com/koyo-os/murocami/internal/config"
	"github.com/koyo-os/murocami/pkg/logger"
	"github.com/nikoksr/notify/service/telegram"
)

type NotifyTheme string

var OK_AGENT NotifyTheme = "ok_agent"
var ERROR_AGENT NotifyTheme = "error_agent"

type Notifyler struct{
	tg *telegram.Telegram
	logger *logger.Logger
}

func Init(cfg *config.Config) (*Notifyler, error) {
	service, err := telegram.New(cfg.TelegrammApiToken)
	if err != nil{
		return nil,err
	}

	return &Notifyler{
		tg: service,
		logger: logger.Init(),
	}, nil
}

func (n *Notifyler) Send(themeid NotifyTheme) error {
	
}