package config

type Mysql struct {
	Host         string `yaml:"host"`
	Port         string `yaml:"port"`
	DB           string `yaml:"db"`
	User         string `yaml:"user"`
	Password     string `yaml:"password"`
	LoggLevel    string `yaml:"level"`
	MaxIdleConns int    `yaml:"max_idle_conns"`
	MaxOpenConns int    `yamlL:maxopen`
}

func (m *Mysql) Dsn() string {
	return m.User + ":" + m.Password + "@tcp(" + m.Host + ":" + m.Port + ")/" + m.DB + "?charset=utf8mb4&parseTime=True&loc=Local"
}
