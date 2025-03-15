package agent

import (
	"os"
	"os/exec"

	"github.com/koyo-os/murocami/internal/agent/history"
	"github.com/koyo-os/murocami/internal/config"
	"github.com/koyo-os/murocami/internal/queue"
	"github.com/koyo-os/murocami/internal/utils"
	"github.com/koyo-os/murocami/pkg/logger"
	"github.com/koyo-os/murocami/pkg/notify"
)

type Agent struct{
	Dir string
	Logger *logger.Logger
	cfg *config.Config
	pipeRunner *PipeLineRunner
	history *history.AgentHistory
	queue *queue.QueueRunner
	queueCFG *config.QueueConfig
	notify *notify.Notifyler
}

const ERROR_MESSAGE = `
CI process not complited`

const SUCCESS_MESSAGE = `
CI process successfully complited`

func Init(cfg *config.Config) (*Agent, error) {
	logger := logger.Init()

	logger.Infof("Creating temp for %s", cfg.TempDirName)
	err := os.Mkdir(cfg.TempDirName, 0755)
	if err != nil{
		return nil,err
	}

	quecfg, err := config.InitQueueConfig()
	if err != nil{
		logger.Errorf("cant get que config: %v",err)
		return nil, err
	}

	return initAgent(cfg, logger, cfg.TempDirName, quecfg)
}

func (a *Agent) Run(url string) (bool, error) {
	okAgent := true

	a.Logger.Infof("starting agent for %s", url)

	if err := utils.CloneRepo(url, a.Dir, a.Logger);err != nil{
		a.Logger.Error(err)
		okAgent = false
		return okAgent, err
	}

	if err := a.RunTests();err != nil{
		a.Logger.Error(err)
		okAgent = false
	}

	if err := a.RunBuild();err != nil{
		a.Logger.Errorf("cant build: %v", err)
		okAgent = false
		return okAgent, err
	}

	if err := a.RunLint();err != nil{
		a.Logger.Errorf("error lint: %v", err)
		return okAgent, err
	}

	if a.cfg.UseScpForCD {
		if err := a.pipeRunner.RunPipeline();err != nil{
			a.Logger.Errorf("error run pipeline: %v",err)
			okAgent = false
			return okAgent, err
		}
	}

	if a.cfg.SaveHistory {
		var message string

		if okAgent {
			message = SUCCESS_MESSAGE
		} else {
			message = ERROR_MESSAGE
		}

		if err := a.history.Save(okAgent, message);err != nil{
			a.Logger.Errorf("cant save history: %v",err)
		}
	}

	return okAgent, nil
}

func (a *Agent) RunTests() error {
	a.Logger.Info("starting test...")

	cmd := exec.Command("go", "test", "./...")
	cmd.Dir = a.Dir
	output, err := cmd.CombinedOutput()
	if err != nil{
		a.Logger.Error(err)
	}

	a.Logger.Infof("test output: %s", output)
	return nil
}

func (a *Agent) RunBuild() error {
	a.Logger.Info("starting build...")

	cmd := exec.Command("go", "build", "-o", a.cfg.OutputPoint, a.cfg.InputPoint)
	cmd.Dir = a.cfg.TempDirName
	output,err := cmd.CombinedOutput()
	if err != nil{
		a.Logger.Error(err)
		return err
	}

	a.Logger.Infof("build output: %s", output)
	return nil
}

func (a *Agent) RunLint() error {
	a.Logger.Info("starting lint...")

	cmd := exec.Command("golangci-lint", "run")
	cmd.Dir = a.Dir
	output,err := cmd.CombinedOutput()
	if err != nil{
		a.Logger.Errorf("cant do lint: %v, with output: %s",err, output)
		return err
	}

	a.Logger.Infof("lint output: %s", output)

	return nil
}