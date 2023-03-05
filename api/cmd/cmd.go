package cmd

import (
	"inception/api/config"
	"inception/api/internal"
	"inception/api/internal/global"
	"log"
)

func Run() {
	global.Config = InitConfig()
	internal.Start()
}

const (
	COnfigFilePath = "configs/api.yaml"
)

func InitConfig() *config.Config {
	c := config.New()
	if err := c.Load(COnfigFilePath); err != nil {
		panic(err)
	}
	log.Println("config:", c)
	return c
}
