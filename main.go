package main

import (
	"net/http"
	"os"
	"time"
	"zerossl_verify/config"

	"github.com/gin-gonic/gin"
)

func main() {

	config.MustLoadConfig(config.CONFIG_FILE_PATH, &config.Conf)

	r := gin.Default()

	// 定义路由处理程序
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "project: https://github.com/BacklitCat/zerossl_verify")
	})

	r.GET("/stat", func(c *gin.Context) {
		c.String(http.StatusOK, "ok")
	})

	r.GET("/stat/", func(c *gin.Context) {
		c.String(http.StatusOK, "ok")
	})

	r.GET(config.Conf.Server.UrlPath, func(c *gin.Context) {
		filename := c.Param("filename")

		//打开文件
		file, err := os.Open(filename)
		if err != nil {
			c.String(http.StatusInternalServerError, "Failed to open file")
			return
		}
		defer file.Close()

		// 设置响应头为text/plain类型
		c.Header("Content-Type", "text/plain")

		// 将文件内容传输给客户端
		http.ServeContent(c.Writer, c.Request, "", time.Now(), file)
	})

	// 运行Gin服务器
	r.Run(":" + config.Conf.Server.Port)
}
