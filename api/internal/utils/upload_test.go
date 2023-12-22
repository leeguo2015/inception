package utils

import (
	"inception/api/internal/global"
	"os"
	"testing"
)

func TestUploadFiileToCos(t *testing.T) {
	data, err := os.ReadFile("./upload.go")
	if err != nil {
		t.Error(err)
	}
	global.ConfigFilePath = "/Users/v_liguo/goPath/inception/configs/inceptionApi.yaml"
	global.InitConfig()
	SecretID := global.Config.Oss.SecretID
	SecretKey := global.Config.Oss.SecretKey
	BucketURL := global.Config.Oss.Url
	t.Logf("OSS: %#v\n", global.Config.Oss)
	name, err := UploadFileToCos("upload1.txt", data, BucketURL, SecretID, SecretKey)
	if err != nil {
		t.Error(err)
	}
	t.Log(name)
}
