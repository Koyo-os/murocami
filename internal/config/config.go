package config

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

type Config struct{
	Port string `yaml:"port"`
	Host string `yaml:"host"`
	TempDirName string `yaml:"temp_dir_name"`
	InputPoint string `yaml:"input_point"`
	OutputPoint string `yaml:"output_point"`
}

func Init() (*Config, error) {
	godotenv.Load(".env")

	port := os.Getenv("PORT")
	host := os.Getenv("HOST")
	temp := os.Getenv("TEMP_DIR_NAME")
	input := os.Getenv("INPUT_POINT")
	output := os.Getenv("OUTPUT_POINT")

	if port != "" && host != "" && temp != "" && input != "" && output != "" {
		return &Config{
			Host: host,
			Port: port,
			TempDirName: temp,
			InputPoint: input,
			OutputPoint: output,
		}, nil
	}
	
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	viper.SetConfigType("yaml")

	if err := viper.ReadInConfig();err != nil{
		return nil,err
	}

	var cfg Config

	if err := viper.Unmarshal(&cfg);err != nil{
		return nil,err
	}

	return &cfg, nil
}