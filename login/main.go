package main

import (
	"fmt"
	"loginRegist/db"
	"loginRegist/model"
	"time"

	"github.com/gin-gonic/gin"
)

// 获取验证码，存入Redis，有效时间5分钟，
// 注册时判断验证码是否正确，     参数：邮箱，验证码，密码
// 登录 验证码登录或者密码登录

//TODO:
//修改密码功能  通过原始密码和验证码同时相等的情况才能修改
//redis 启动时就连接
//文章功能 {id,标题，作者，时间，内容}
//文章的增删改查实现，支持模糊查询

func init() {
	//调用db.InitMysql方法，并把数据传给host
	db.InitMysql("root:root@tcp(localhost:3306)/test?charset=utf8&parseTime=True&loc=Local")
}

func main() {
	e := gin.Default()
	e.LoadHTMLGlob("cshtml/*") //获取当前路径下的cshtml文件夹下的所有html文件
	e.POST("/getCode", GetCode)
	e.POST("/doregist", DoRegist)
	e.POST("/dologin", DoLogin)

	e.Run()
}

//注册
func Regist(c *gin.Context) {
	c.HTML(200, "regist.html", nil)
}

type UserReq struct {
	Email string `json:"email"`
	Value string `json:"value"`
	Code  string `json:"code"`
	Type  int64  `json:"type"`
}

type BaseResp struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

//获取验证码
func GetCode(c *gin.Context) {
	resp := BaseResp{}
	req := new(UserReq)
	if err := c.BindJSON(req); err != nil {
		resp.Msg = "解析错误"
		c.JSON(200, &resp)
		return
	}

	key := "Regist_Code_" + req.Email
	_, err := db.Redis().Set(key, 123456, 300*time.Second).Result()
	if err != nil {
		fmt.Println("Redis Err", err)
	}
	defer db.Redis().Close()
	resp.Msg = "验证码获取成功"
	c.JSON(200, resp.Msg)
}

//注册
func DoRegist(c *gin.Context) {
	resp := BaseResp{}
	req := new(UserReq)
	if err := c.BindJSON(req); err != nil {
		resp.Msg = "解析错误"
		c.JSON(200, &resp)
		return
	}

	//获取数据库数据
	user := model.User{}

	db.Mysql().Table("users").Where("email = ?", req.Email).First(&user)
	//存在
	if req.Email == user.Email {
		c.String(200, "邮箱已存在,请登录")
		return
	}

	//获取Code码
	key := "Regist_Code_" + req.Email
	code, err := db.Redis().Get(key).Result()
	if err != nil {
		c.String(200, "Redis读取错误")
		return
	}
	if req.Code == code {

		user := model.User{}
		user.Email = req.Email
		user.Password = req.Value
		db.Mysql().Table("users").Debug().Create(&user)

		resp.Msg = "注册成功"
		c.JSON(200, resp.Msg)
		db.Redis().Del(key)
	}
}

//登录
func Login(c *gin.Context) {
	c.HTML(200, "login.html", nil)
}

//登录
func DoLogin(c *gin.Context) {
	resp := BaseResp{}
	req := new(UserReq)
	if err := c.BindJSON(req); err != nil {
		resp.Msg = "解析错误"
		c.JSON(200, &resp)
		return
	}
	user := model.User{}
	db.Mysql().Table("users").Where("email = ?", req.Email).First(&user)

	if req.Type == 1 { //1.密码登录  2.验证码登录
		//密码登录
		if req.Email == user.Email && req.Value == user.Password {
			resp.Msg = "登录成功"
			c.JSON(200, resp.Msg)
			return
		}
		resp.Msg = "账户或者密码错误"
		c.JSON(200, resp.Msg)
	}
	key := "Regist_Code_" + req.Email
	code, err := db.Redis().Get(key).Result()
	defer db.Redis().Close()

	if err != nil {
		c.String(200, "Redis读取错误")
		return
	}

	//验证码登录
	if req.Email == user.Email && req.Value == code {
		resp.Msg = "登录成功"
		c.JSON(200, resp.Msg)
		db.Redis().Del(key)
		return
	}
	resp.Msg = "账户或者验证码错误"
	c.JSON(200, resp.Msg)
}
