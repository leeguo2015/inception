package model

import "inception/api/internal/global"

func AutoMigrate() {
	global.Log.Info("初始化用户表")
	if err := UserMigrate(); err != nil {
		global.Log.Error(err)
	}
}
