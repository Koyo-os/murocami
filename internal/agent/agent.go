package agent

import (
	"os"
	"os/exec"

	"github.com/koyo-os/murocami/internal/config"
	"github.com/koyo-os/murocami/pkg/logger"
)

type Agent struct{
	TempDir *os.File
	Dir string
	Logger *logger.Logger
	cfg *config.Config
}

func Init(cfg *config.Config) (*Agent, error) {
	tempDir, err := os.CreateTemp("", cfg.TempDirName)
	if err != nil{
		return nil,err
	}

	return &Agent{
		cfg: cfg,
		Logger: logger.Init(),
		Dir: tempDir.Name(),
		TempDir: tempDir,
	}, nil
}

func (a *Agent) RunTests() error {
	a.Logger.Info("starting test...")

	cmd := exec.Command("go", "test", "./...")
	cmd.Dir = a.Dir
	output, err := cmd.CombinedOutput()
	if err != nil{
		a.Logger.Error(err)
		return err
	}

	a.Logger.Infof("test output: %s", output)
	return nil
}

func (a *Agent) RunBuild() error {
	a.Logger.Info("starting build...")

	cmd := exec.Command("go", "build", "-o", a.cfg.OutputPoint, a.cfg.InputPoint)
	output,err := cmd.CombinedOutput()
	if err != nil{
		a.Logger.Error(err)
		return err
	}

	a.Logger.Infof("build output: %s", output)
	return nil
}
