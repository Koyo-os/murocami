package config

import (
	"io"
	"os"

	"gopkg.in/yaml.v3"
)

type PipeLine struct{
	Name string `yaml:"name"`
	Cmd string `yaml:"cmd"`
	MoreCmd bool `yaml:"more_cmd"`
	Commands []string `yaml:"commands"`
}

type PipeLineConfig struct{
	RunOn string `yaml:"run_on"`
	ServiceName string `yaml:"service_name"`
	Cmds []PipeLine `yaml:"cmds"`
}

func LoadPipeLineConfig() (*PipeLineConfig, error) {
	file,err := os.Open("pipeline.yaml")
	if err != nil{
		return nil,err
	}

	body,err := io.ReadAll(file)
	if err != nil{
		return nil,err
	}

	var cfg PipeLineConfig
	if err = yaml.Unmarshal(body, &cfg);err != nil{
		return nil,err
	}

	return &cfg, nil
}