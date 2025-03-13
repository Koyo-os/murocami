package config

import (
	"fmt"
	"io"
	"os"

	"gopkg.in/yaml.v3"
)

type QueueConfig struct{
	UseQueue bool `yaml:"use_queue"`
	NatsUrl string `yaml:"nats_url"`
	NatsTheme string `yaml:"nats_theme"`
}

func InitQueueConfig() (*QueueConfig, error) {
	file,err := os.Open("queue_config.yaml")
	if err != nil{
		return nil,fmt.Errorf("cant get queue config: %v",err)
	}

	body,err := io.ReadAll(file)
	if err != nil{
		return nil, fmt.Errorf("error read file: %v", err)
	}

	var queueCFG QueueConfig

	if err = yaml.Unmarshal(body, queueCFG);err != nil{
		return nil, fmt.Errorf("unmarshal error: %v", err)
	}

	return &queueCFG, nil
}