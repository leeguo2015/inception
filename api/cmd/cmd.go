package cmd

import (
	"inception/api/internal"
	"inception/api/internal/global"
)

const (
	//ConfigFilePath = "configs/api.yaml"
	ConfigFilePath = "configs/inceptionApi.yaml"
)

func Run() {
	global.InitConfig()
	internal.Start()
}
