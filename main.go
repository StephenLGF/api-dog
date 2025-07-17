package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()

	r.Static("/assets", "./public/assets")
	r.StaticFile("/vite.svg", "./public/vite.svg")

	// HTML 文件服务
	r.LoadHTMLFiles("./public/index.html")

	// 其他后端接口
	r.GET("/api/hello", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello from Gin backend",
		})
	})

	r.GET("/", func(c *gin.Context) {
		c.HTML(200, "index.html", gin.H{})
	})

	r.Run(":8080")
}
