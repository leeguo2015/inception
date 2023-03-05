package config

type Mysql struct {
	Host      string `yaml:"host"`
	Port      string `yaml:"port"`
	DB        string `yaml:"db"`
	User      string `yaml:"user"`
	Password  string `yaml:"password"`
	LoggLevel string `yaml:"level"`
}
