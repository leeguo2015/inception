package model

import "inception/api/internal/global"

func AutoMigrate() {
	global.Log.Info("初始化用户表")
	if err := UserMigrate(); err != nil {
		global.Log.Error(err)
	}
	global.Log.Info("初始化博客表")
	if err := BlogMigrate(); err != nil {
		global.Log.Error(err)
	}

}
