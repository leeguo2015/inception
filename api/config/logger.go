package config

type Logger struct {
	Level        string `yaml:"level"`          // level: INFO
	Prefix       string `yaml:"prefix"`         // prefix: "[inception]"
	ShowLine     bool   `yaml:"show_line"`      // show-line: true
	LogInConsole bool   `yaml:"log_in_console"` // log-in-console: true

}
