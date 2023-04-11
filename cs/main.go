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
	fmt.Println(string(data))
	resp.Write(data)
}

func main() {
	http.Handle("/emailLogin", &emailLogin{})
	http.ListenAndServe(":8888", nil)

	emailRegist := model.EmailRegist{}
	http.HandleFunc("/emailRegist", func(w http.ResponseWriter, r *http.Request) {
		//序列号结构体
		data, err := json.Marshal(emailRegist)
		if err != nil {
			w.Write([]byte(err.Error()))
			return
		}
		//响应序列化结果
		w.Write([]byte(data))
	})
}
