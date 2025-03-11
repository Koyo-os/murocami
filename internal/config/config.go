package config

import (
	"fmt"
	"io"
	"os"

	"github.com/joho/godotenv"
	"gopkg.in/yaml.v3"
)

type Config struct{ 
	Port string `yaml:"port"`
	Host string `yaml:"host"`
	TelegrammApiToken string `yaml:"telegramm_api_token"`
	NotifyChatId string `yaml:"notify_chat_id"`
	TempDirName string `yaml:"temp_dir_name"`
	InputPoint string `yaml:"input_point"`
	OutputPoint string `yaml:"output_point"`
	UseScpForCD bool `yaml:"scp_for_cd"`
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
			Host: host,
			Port: port,
			TempDirName: temp,
			InputPoint: input,
			OutputPoint: output,
		}, nil
	}
	
	file, err := os.Open("config.yaml")
	if err != nil{
		return nil, fmt.Errorf("cant open config file: %v",err)
	}

	body,err := io.ReadAll(file)
	if err != nil{
		return nil, fmt.Errorf("cant read file: %v",err)
	}

	var cfg Config

	if err = yaml.Unmarshal(body, &cfg);err != nil{
		return nil, fmt.Errorf("cant unmarshal data: %v",err)
	}

	cfg.TelegrammApiToken = token

	return &cfg, nil
}