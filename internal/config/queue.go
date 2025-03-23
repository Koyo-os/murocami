package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type QueueConfig struct {
	UseQueue  bool   `yaml:"use_queue"`
	NatsUrl   string `yaml:"nats_url"`
	NatsTheme string `yaml:"nats_theme"`
}

func InitQueueConfig() (*QueueConfig, error) {
	file, err := os.Open("queue_config.yaml")
	if err != nil {
		return nil, fmt.Errorf("cant get queue config: %v", err)
	}

	var cfg QueueConfig
	if err := yaml.NewDecoder(file).Decode(&cfg); err != nil {
		return nil, fmt.Errorf("cant decode config: %v", err)
	}

	return &cfg, nil
}

