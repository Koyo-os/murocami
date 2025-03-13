package queue

import (
	"time"

	"github.com/bytedance/sonic"
	"github.com/koyo-os/murocami/internal/config"
	"github.com/koyo-os/murocami/internal/models"
	"github.com/koyo-os/murocami/pkg/logger"
	"github.com/nats-io/nats.go"
)

type QueueRunner struct{
	logger *logger.Logger
	cfg *config.Config
	queueConfig *config.QueueConfig
	conn *nats.Conn
}

func Init(cfg *config.Config) *QueueRunner {
	logger := logger.Init()

	quecfg, err := config.InitQueueConfig()
	if err != nil{
		logger.Errorf("cant load queue config: %v",err)
		return nil
	}

	var conn *nats.Conn

	if quecfg.NatsUrl == "default" {
		conn, err = nats.Connect(nats.DefaultURL)
		if err != nil{
			logger.Errorf("cant connect nats: %v",err)
			return nil
		}
		defer conn.Close()
	} else {
		conn,err = nats.Connect(quecfg.NatsUrl)
		if err != nil{
			logger.Errorf("cant connect nats: %v",err)
			return nil
		}
		defer conn.Close()
	}

	return &QueueRunner{
		queueConfig: quecfg,
		logger: logger,
		cfg: cfg,
		conn: conn,
	}
}

func (q *QueueRunner) AddInfo(ok bool, message string) error {
	block := models.Block{
		Ok: ok,
		Message: message,
		TimeStamp: time.Now().Format("2006-01-02 15:04:05"),
	}

	body, err := sonic.Marshal(&block)
	if err != nil{
		q.logger.Errorf("cant marshal: %v",err)
		return err
	}

	err = q.conn.Publish(q.queueConfig.NatsTheme, body)
	if err != nil{
		q.logger.Errorf("cant public notify %v",err)
	}

	q.logger.Info("published notify")

	return nil
}