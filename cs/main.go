package main

import (
	"cs/db"
	"cs/model"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func init() {
	//调用db.InitMysql方法，并把数据传给host
	db.InitMysql("root:root@tcp(192.168.100.44:3306)/tongzi?charset=utf8&parseTime=True&loc=Local")
}

type emailLogin struct{}

func (l *emailLogin) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	//判断请求方式是否为POST
	if req.Method != "POST" {
		resp.Write([]byte("请求方式错误"))
		return
	}

	//读取请求body
	data, err := io.ReadAll(req.Body)
	if err != nil {
		resp.Write([]byte(err.Error()))
		return
	}

	//解析请求body
	var emailLogin model.EmailLogin
	err = json.Unmarshal(data, &emailLogin)
	if err != nil {
		resp.Write([]byte(err.Error()))
		return
	}
	//查询该邮箱是否存在
	var user model.User
	db.Mysql().Where("email = ?", emailLogin.Email).Find(&user)
	if user.Email != emailLogin.Email || user.Password != emailLogin.PassWord {
		resp.Write([]byte("账号或者密码错误，请重新输入"))
		return
	}

	fmt.Println("登录成功")
	resp.Write(data)
}

type emailRegist struct{}

func (r *emailRegist) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	//1.判断请求方式
	if req.Method != "POST" {
		resp.Write([]byte("请求方式有误"))
	}
	//2.读取请求body
	data, err := io.ReadAll(req.Body)
	if err != nil {
		resp.Write([]byte(err.Error()))
		return
	}
	//3.解析请求body
	var emailRegist model.EmailRegist
	err = json.Unmarshal(data, &emailRegist)
	if err != nil {
		resp.Write([]byte(err.Error()))
		return
	}
	//4.查询数据库是否存在
	var user model.User
	db.Mysql().Where("email = ?", emailRegist.Email).Find(&user)
	if user.Email == emailRegist.Email {
		resp.Write([]byte("邮箱已存在，请直接登录"))
		return
	}

	//5.插入数据库
	db.Mysql().Create(&model.User{
		Email:    emailRegist.Email,
		Password: emailRegist.PassWord,
	})
	fmt.Println("添加成功")

	//6.响应请求
	resp.Write(data)

}

func main() {
	//登录测试
	http.Handle("/emailLogin", &emailLogin{})
	//注册测试
	http.Handle("/emailRegist", &emailRegist{})
	http.ListenAndServe(":8888", nil)
}
