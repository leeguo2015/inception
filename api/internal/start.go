/*
 * @Author: leeguo leeguo2015@163.com
 * @Date: 2023-10-01 17:56:56
 * @LastEditors: leeguo leeguo2015@163.com
 * @LastEditTime: 2023-12-12 00:23:01
 * @FilePath: \inception\api\internal\start.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package internal

import (
	"inception/api/internal/global"
	"inception/api/internal/middleware"

	// "time"

	"github.com/gin-gonic/gin"
)

func Start() {
	// autoMigrate()

	global.Log.Info("Starting Inception API")
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	router.Use(gin.Recovery())
	router.Use(middleware.JWTAuth())
	router.Use(middleware.Cors())
	router.GET("/", func(c *gin.Context) {
		c.String(200, "OK")
	})

	paddingRouterV1(router)
	if err := router.Run(":" + global.Config.System.Port); err != nil {
		global.Log.Error(err)
	}
}
