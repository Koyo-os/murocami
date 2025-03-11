package config

type PipeLine struct{
	Name string `yaml:"name"`
	Cmd string `yaml:"cmd"`
}

type PipeLineConfig struct{
	RunOn string `yaml:"run_on"`
	ServiceName string `yaml:"service_name"`
	Cmds []PipeLine `yaml:"cmds"`
}