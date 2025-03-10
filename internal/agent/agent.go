package agent

import (
	"os"
	"os/exec"

	"github.com/koyo-os/murocami/internal/config"
	"github.com/koyo-os/murocami/internal/utils"
	"github.com/koyo-os/murocami/pkg/logger"
)

type Agent struct{
	Dir string
	Logger *logger.Logger
	cfg *config.Config
}

func Init(cfg *config.Config) (*Agent, error) {
	logger := logger.Init()

	logger.Infof("Creating temp for %s", cfg.TempDirName)
	err := os.Mkdir(cfg.TempDirName, 0755)
	if err != nil{
		return nil,err
	}

	return &Agent{
		cfg: cfg,
		Logger: logger,
		Dir: cfg.TempDirName,
	}, nil
}

func (a *Agent) Run(url string) error {
	a.Logger.Infof("starting agent for %s", url)

	if err := utils.CloneRepo(url, a.Dir, a.Logger);err != nil{
		a.Logger.Error(err)
		return err
	}

	if err := a.RunTests();err != nil{
		a.Logger.Error(err)
		return err
	}

	if err := a.RunBuild();err != nil{
		a.Logger.Error(err)
		return err
	}

	return nil
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
