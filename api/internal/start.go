/*
 * @Author: leeguo leeguo2015@163.com
 * @Date: 2023-10-01 17:56:56
 * @LastEditors: leeguo leeguo2015@163.com
 * @LastEditTime: 2023-12-22 23:52:02
 * @FilePath: \inception\api\internal\start.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package internal

import (
	"fmt"
	"inception/api/internal/global"
	"inception/api/internal/middleware"
	"io"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

func Start() {
	// autoMigrate()
	f, err := os.Create(time.Now().Format("2006_01_02_") + "gin.log")
	if err != nil {
		fmt.Println("无法创建日志文件:", err)
		return
	}
	defer f.Close()
	// 将 Gin 的默认日志记录器设置为写入到文件的记录器
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
	global.Log.Info("Starting Inception API")
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	router.Use(gin.Recovery())
	router.Use(middleware.JWTAuth())
	router.Use(middleware.Cors()) // 全局CORS中间件
	router.GET("/", func(c *gin.Context) {
		c.String(200, "OK")
	})

	paddingRouterV1(router)
	if err := router.Run(":" + global.Config.System.Port); err != nil {
		global.Log.Error(err)
	}
}
