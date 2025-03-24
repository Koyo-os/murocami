package config

import (
	"fmt"
	"io"
	"os"

	"github.com/joho/godotenv"
	"gopkg.in/yaml.v3"
)

type NagentCfg struct {
	Use  bool   `yaml:"use"`
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
}

type HistoryCfg struct {
	Save bool   `yaml:"save"`
	File string `yaml:"file"`
}

type NotifyCfg struct {
	Send   bool   `yaml:"send"`
	ChatID int64  `yaml:"chat_id"`
	Token  string `yaml:"token"`
}

type Config struct {
	Port        string     `yaml:"port"`
	Host        string     `yaml:"host"`
	NotifyCfg   NotifyCfg  `yaml:"notify"`
	HistoryCfg  HistoryCfg `yaml:"history"`
	StaticDir   string     `yaml:"static_dir"`
	TempDirName string     `yaml:"temp_dir_name"`
	InputPoint  string     `yaml:"input_point"`
	OutputPoint string     `yaml:"output_point"`
	UseScpForCD bool       `yaml:"scp_for_cd"`
	UseUI       bool       `yaml:"use_ui"`
	Nagent      NagentCfg  `yaml:"nagent"`
}

func Init() (*Config, error) {
	godotenv.Load(".env")

	port := os.Getenv("PORT")
	host := os.Getenv("HOST")
	temp := os.Getenv("TEMP_DIR_NAME")
	input := os.Getenv("INPUT_POINT")
	output := os.Getenv("OUTPUT_POINT")
	token := os.Getenv("TELEGRAMM_API")

	if port != "" && host != "" && temp != "" && input != "" && output != "" {
		return &Config{
			Host:        host,
			Port:        port,
			TempDirName: temp,
			InputPoint:  input,
			OutputPoint: output,
		}, nil
	}

	file, err := os.Open("config.yaml")
	if err != nil {
		return nil, fmt.Errorf("cant open config file: %v", err)
	}

	body, err := io.ReadAll(file)
	if err != nil {
		return nil, fmt.Errorf("cant read file: %v", err)
	}

	var cfg Config

	if err = yaml.Unmarshal(body, &cfg); err != nil {
		return nil, fmt.Errorf("cant unmarshal data: %v", err)
	}

	cfg.NotifyCfg.Token = token

	return &cfg, nil
}
