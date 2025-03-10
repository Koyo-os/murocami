package config

import "github.com/spf13/viper"

type Config struct{
	Port string `yaml:"port"`
	Host string `yaml:"host"`
	TempDirName string `yaml:"temp_dir_name"`
	InputPoint string `yaml:"input_point"`
	OutputPoint string `yaml:"output_point"`
}

func Init() (*Config, error) {
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