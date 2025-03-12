package agent

import (
	"os/exec"
	"strings"

	"github.com/koyo-os/murocami/internal/config"
	"github.com/koyo-os/murocami/pkg/logger"
	"gopkg.in/yaml.v3"
)

type PipeLineRunner struct{
	pipeCFG *config.PipeLineConfig
	cfg *config.Config
	logger *logger.Logger
}

const DEFAULT_PIPELINE_CFG = `
run_on: localhost
service_name: my_service 
cmds:  
  - name: Build docker image
    more_cmd : false
    cmd: docker build -t app .
  - name: Copy docker image
    more_cmd : true
    commands: |
      docker save -o app.tar app
      scp app.tar user@yourserver:/path/to
  - name: Deploy
    more_cmd : true
    commands: |
      ssh user@yourserver "docker load -i /path/to/destination/myapp.tar"
      ssh user@yourserver "docker stop myapp || true"
      ssh user@yourserver "docker rm myapp || true"
      ssh user@yourserver "docker run -d --name myapp -p 80:80 myapp:latest"`

func InitPipelineRunner(cfg *config.Config) *PipeLineRunner {
	logger := logger.Init()
	pipe,err := config.LoadPipeLineConfig()
	if err != nil{
		logger.Errorf("Cant load pipelines: %v, i use default", err)
		var def config.PipeLineConfig

		yaml.Unmarshal([]byte(DEFAULT_PIPELINE_CFG), &def)
		return &PipeLineRunner{
			pipeCFG: &def,
			cfg: cfg,
			logger: logger,
		}
	}

	return &PipeLineRunner{
		cfg: cfg,
		pipeCFG: pipe,
		logger: logger,
	}
}

func (pipe *PipeLineRunner) RunPipeline() error {
	for _, p := range pipe.pipeCFG.Cmds{
		pipe.logger.Infof("[PIPELINE] run now %s", p.Name)
		parts := strings.Split(p.Cmd, " ")
		cmd := exec.Command(parts[0], parts[1:]...)
		output, err := cmd.CombinedOutput()
		if err != nil{
			pipe.logger.Errorf("[PIPELINE] name: %s, error to run: %v, with output %s", p.Name, err, output)
			return err
		}
	}

	return nil
}