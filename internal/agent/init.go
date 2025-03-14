package agent

import (
	"github.com/koyo-os/murocami/internal/agent/history"
	"github.com/koyo-os/murocami/internal/config"
	"github.com/koyo-os/murocami/internal/queue"
	"github.com/koyo-os/murocami/pkg/logger"
	"github.com/koyo-os/murocami/pkg/notify"
)

func initAgent(cfg *config.Config, logger *logger.Logger, dir string, que *config.QueueConfig) (*Agent,error) {
	if cfg.SaveHistory && cfg.SendNotify && que.UseQueue && cfg.UseScpForCD {
		notify,err := notify.Init(cfg)
		if err != nil{
			logger.Errorf("cant init notify: %v",err)
			return nil,err
		}

		return &Agent{
			Logger: logger,
			cfg: cfg,
			notify: notify,
			pipeRunner: InitPipelineRunner(cfg),
			queue: queue.Init(cfg),
			history: history.Init(cfg),
			queueCFG: que,
			Dir: dir,
		}, nil
	} else if cfg.SaveHistory && cfg.SendNotify && cfg.UseScpForCD {
		notify,err := notify.Init(cfg)
		if err != nil{
			logger.Errorf("cant init notify: %v",err)
			return nil,err
		}

		return &Agent{
			Logger: logger,
			cfg: cfg,
			notify: notify,
			pipeRunner: InitPipelineRunner(cfg),
			history: history.Init(cfg),
			Dir: dir,
		}, nil
	} else if cfg.SaveHistory && que.UseQueue && cfg.UseScpForCD {
		return &Agent{
			Logger: logger,
			cfg: cfg,
			pipeRunner: InitPipelineRunner(cfg),
			queue: queue.Init(cfg),
			history: history.Init(cfg),
			queueCFG: que,
			Dir: dir,
		}, nil
	} else if que.UseQueue && cfg.SendNotify && cfg.UseScpForCD {
		notify,err := notify.Init(cfg)
		if err != nil{
			logger.Errorf("cant init notify: %v",err)
			return nil,err
		}

		return &Agent{
			Logger: logger,
			cfg: cfg,
			notify: notify,
			pipeRunner: InitPipelineRunner(cfg),
			queue: queue.Init(cfg),
			queueCFG: que,
			Dir: dir,
		}, nil
	} else if que.UseQueue && cfg.SendNotify && cfg.SaveHistory {
		notify,err := notify.Init(cfg)
		if err != nil{
			logger.Errorf("cant init notify: %v",err)
			return nil,err
		}

		return &Agent{
			Logger: logger,
			cfg: cfg,
			notify: notify,
			queue: queue.Init(cfg),
			queueCFG: que,
			Dir: dir,
		}, nil
	} else if que.UseQueue {
		notify,err := notify.Init(cfg)
		if err != nil{
			logger.Errorf("cant init notify: %v",err)
			return nil,err
		}

		return &Agent{
			Logger: logger,
			cfg: cfg,
			notify: notify,
			pipeRunner: InitPipelineRunner(cfg),
			queue: queue.Init(cfg),
			queueCFG: que,
			Dir: dir,
		}, nil
	} else {
			notify,err := notify.Init(cfg)
			if err != nil{
				logger.Errorf("cant init notify: %v",err)
				return nil,err
			}
	
			return &Agent{
				Logger: logger,
				cfg: cfg,
				notify: notify,
				pipeRunner: InitPipelineRunner(cfg),
				queue: queue.Init(cfg),
				history: history.Init(cfg),
				queueCFG: que,
				Dir: dir,
			}, nil
	}
}