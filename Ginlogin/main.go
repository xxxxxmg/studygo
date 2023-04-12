package main

import (
	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	c.HTML(200, "login.html", nil)
}

func DoLogin(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	c.HTML(200, "welcome.html", gin.H{
		"Username": username,
		"Password": password,
	})
}

func main() {
	e := gin.Default()
	e.LoadHTMLGlob("cshtml/*") //获取当前路径下的cshtml文件夹下的所有html文件
	e.GET("/login", Login)
	e.POST("/login", DoLogin) //POST请求交给DoLogin处理
	e.Run()
}
