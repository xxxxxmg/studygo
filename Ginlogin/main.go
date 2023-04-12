package main

import (
	"login/db"
	"login/model"

	"github.com/gin-gonic/gin"
)

func init() {
	//调用db.InitMysql方法，并把数据传给host
	db.InitMysql("root:root@tcp(192.168.100.44:3306)/tongzi?charset=utf8&parseTime=True&loc=Local")
}

//登录
func Login(c *gin.Context) {
	c.HTML(200, "login.html", nil)
}

//判断邮箱密码是否正确
func DoLogin(c *gin.Context) {
	//获取页面输入的邮箱
	email := c.PostForm("email")
	//获取页面输入的密码
	password := c.PostForm("password")
	//查询数据库中这个邮箱是否存在
	var user model.User
	db.Mysql().Where("email = ?", email).Find(&user)
	//判断数据是否正确
	if user.Email == email && user.Password == password {
		//数据写入到指定的HTML网页中，并响应200
		c.HTML(200, "welcomelogin.html", gin.H{
			"Email": email,
		})
	} else {
		c.String(200, "账号或者密码错误")
	}
}

//注册
func Regist(c *gin.Context) {
	//数据写入到指定的HTML网页中，并响应200
	c.HTML(200, "regist.html", nil)
}

//判断邮箱是否存在
func DoRegist(c *gin.Context) {
	//获取邮箱
	email := c.PostForm("email")
	//获取密码
	password := c.PostForm("password")
	//获取数据库数据
	var user model.User
	db.Mysql().Debug().Where("email = ?", email).Find(&user)
	//判断邮箱是否已存在
	if email == user.Email {
		c.String(200, "邮箱已存在")
	} else {
		//不存在，写入数据库
		db.Mysql().Create(&model.User{
			Email:    email,
			Password: password,
		})
		//数据写入到指定的HTML网页中，并响应200
		c.HTML(200, "welcomeregist.html", gin.H{
			"Email":    email,
			"Password": password,
		})
	}
}

func main() {
	e := gin.Default()
	e.LoadHTMLGlob("cshtml/*") //获取当前路径下的cshtml文件夹下的所有html文件
	e.GET("/login", Login)
	e.POST("/login", DoLogin) //POST请求交给DoLogin处理
	e.GET("/regist", Regist)
	e.POST("/regist", DoRegist) //POST请求交给DoRegist处理
	e.Run()
}
