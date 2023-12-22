package config

type System struct {
	Host string `yaml:"host"` // : local
	Port string `yaml:"port"` // : 3306
}

type Oss struct {
	SecretID  string `yaml:"secret_id"`  // :
	SecretKey string `yaml:"secret_secret"` // :
	Url       string `yaml:"url"`        // :
}
